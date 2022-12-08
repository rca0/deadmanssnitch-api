package deadmanssnitch

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

func (s *SnitchService) NewSnitch(snitch *Snitch) (*Snitch, *Response, error) {
	u := "/v1/snitches"
	v := new(Snitch)

	payload := &Snitch{
		Name:        snitch.Name,
		Interval:    snitch.Interval,
		Alert_type:  snitch.Alert_type,
		Alert_email: snitch.Alert_email,
		Notes:       snitch.Notes,
		Tags:        snitch.Tags,
	}

	response, err := s.client.newRequestDo("POST", u, nil, payload, &v)
	// to-do
	// compare if snitch already exists

	if err != nil {
		return nil, nil, err
	}

	return v, response, nil
}
