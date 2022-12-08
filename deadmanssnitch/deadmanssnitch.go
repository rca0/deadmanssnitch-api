package deadmanssnitch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/google/go-querystring/query"
)

const (
	defaultBaseURL = "https://api.deadmanssnitch.com"
	contentType    = "application/json"
	useragent      = "deadmanssnitch-api"
)

type Config struct {
	BaseURL    string
	ApiKey     string
	UserAgent  string
	HTTPClient *http.Client
	debug      bool
}

type Client struct {
	baseURL *url.URL
	client  *http.Client
	Config  *Config
	Snitch  *SnitchService
}

type Response struct {
	Response  *http.Response
	BodyBytes []byte
}

type RequestOptions struct {
	Type  string
	Label string
	Value string
}

type service struct {
	client *Client
}

func NewClient(config *Config) (*Client, error) {
	if config.HTTPClient == nil {
		config.HTTPClient = http.DefaultClient
	}

	if config.BaseURL == "" {
		config.BaseURL = defaultBaseURL
	}

	config.UserAgent = useragent

	baseURL, err := url.Parse(config.BaseURL)
	if err != nil {
		return nil, err
	}

	c := &Client{
		baseURL: baseURL,
		client:  config.HTTPClient,
		Config:  config,
	}

	c.Snitch = &SnitchService{c}

	// InitCache(c)
	// PopulateCache(c)

	return c, nil
}

func (c *Client) newRequest(method, url string, body interface{}, opts ...RequestOptions) (*http.Request, error) {
	var buf io.ReadWriter

	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	if c.Config.debug {
		log.Printf("[DEBUG] Deadmanssnitch - Preparing %s requst to %s with body: %s", method, url, buf)
	}

	u := c.baseURL.String() + url

	req, err := http.NewRequest(method, u, buf)
	if err != nil {
		return nil, err
	}

	if len(opts) > 0 {
		for _, o := range opts {
			if o.Type == "header" {
				req.Header.Add(o.Label, o.Value)
			}
		}
	}

	req.Header.Add("Content-Type", contentType)
	if c.Config.UserAgent != "" {
		req.Header.Add("User-Agent", c.Config.UserAgent)
	}
	req.SetBasicAuth(c.Config.ApiKey, os.Getenv("DEADMANSSNITCH_APIKEY"))

	return req, nil
}

func (c *Client) newRequestDo(method, url string, qryOptions, body, v interface{}) ([]byte, error) {
	if qryOptions != nil {
		values, err := query.Values(qryOptions)
		if err != nil {
			return nil, err
		}

		if v := values.Encode(); v != "" {
			url = fmt.Sprintf("%s?%s", url, v)
		}
	}

	req, err := c.newRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	return c.do(req, v)
}

func (c *Client) do(req *http.Request, v interface{}) ([]byte, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := c.checkResponse(resp, bodyBytes); err != nil {
		return nil, err
	}

	if v != nil {
		if err := c.DecodeJSON(bodyBytes, v); err != nil {
			return nil, nil
		}
	}

	return bodyBytes, nil
}

func (c *Client) checkResponse(r *http.Response, b []byte) error {
	if r.StatusCode >= 200 && r.StatusCode <= 299 {
		return nil
	}
	return c.decodeErrorResponse(r, b)
}

func (c *Client) DecodeJSON(body []byte, v interface{}) error {
	return json.Unmarshal(body, v)
}

func (c *Client) decodeErrorResponse(r *http.Response, b []byte) error {
	v := &errorResponse{Error: &Error{ErrorResponse: r}}

	if err := c.DecodeJSON(b, v); err != nil {
		return fmt.Errorf("%s API call to %s failed: %v", r.Request.Method, r.Request.URL.String(), r.Status)
	}

	return v.Error
}
