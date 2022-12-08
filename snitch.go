package main

type Snitch struct {
	Name        string
	interval    string
	alert_type  string
	alert_email []string
	notes       string
	tags        []string
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

func NewSnitch(client *Client) *Snitch {
	// To be implemented
	return &Snitch{}
}
