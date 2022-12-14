package api

import (
	"encoding/json"
	"fmt"
)

const urlPath string = "/v1/snitches"

type SnitchService service

type Snitch struct {
	Name        string   `json:"name,omitempty"`
	Interval    string   `json:"interval,omitempty"`
	Alert_type  string   `json:"alert_type,omitempty"`
	Alert_email []string `json:"alert_email,omitempty"`
	Notes       string   `json:"notes,omitempty"`
	Tags        []string `json:"tags,omitempty"`
}

type SnitchResponse struct {
	Token       string      `json:"token,omitempty"`
	Href        string      `json:"href,omitempty"`
	Name        string      `json:"name,omitempty"`
	Status      string      `json:"status,omitempty"`
	CreatedAt   string      `json:"created_at,omitempty"`
	CheckInUrl  string      `json:"check_in_url,omitempty"`
	CheckedInAt string      `json:"checked_in_at,omitempty"`
	Type        interface{} `json:"type,omitempty"`
	Interval    string      `json:"interval,omitempty"`
	Alert_type  string      `json:"alert_type,omitempty"`
	Alert_email []string    `json:"alert_email,omitempty"`
	Notes       string      `json:"notes,omitempty"`
	Tags        []string    `json:"tags,omitempty"`
}

func (s *SnitchService) NewSnitch(snitch *Snitch) (*SnitchResponse, error) {
	v := new(Snitch)
	newSnitchResponse := SnitchResponse{}

	payload := &Snitch{
		Name:        snitch.Name,
		Interval:    snitch.Interval,
		Alert_type:  snitch.Alert_type,
		Alert_email: snitch.Alert_email,
		Notes:       snitch.Notes,
		Tags:        snitch.Tags,
	}

	body, err := s.client.newRequestDo("POST", urlPath, nil, payload, &v)
	// to-do
	// compare if snitch already exists

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &newSnitchResponse)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(body))

	return &newSnitchResponse, nil
}

func (s *SnitchService) GetSnitches() (*[]SnitchResponse, error) {
	listSnitch := []SnitchResponse{}

	body, err := s.client.newRequestDo("GET", urlPath, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &listSnitch)
	if err != nil {
		return nil, err
	}

	return &listSnitch, err
}

func (s *SnitchService) GetSnitch(token string) (*SnitchResponse, error) {
	newSnitchResponse := SnitchResponse{}

	body, err := s.client.newRequestDo("GET", urlPath, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &newSnitchResponse)
	if err != nil {
		return nil, err
	}

	return &newSnitchResponse, nil
}
