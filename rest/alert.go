package rest

import (
	"encoding/json"
	"errors"
)

// Alert structure for alerts
type Alert struct {
	Acknowledged bool   `json:"acknowledged"`
	Class        string `json:"class"`
	Count        int    `json:"count"`
	ID           string `json:"id"`
	Level        string `json:"level"`
	Message      string `json:"message"`
	Ref          struct {
		Cluster  string `json:"cluster,omitempty"`
		Host     string `json:"host,omitempty"`
		Pool     string `json:"pool,omitempty"`
		Template string `json:"template,omitempty"`
		Guest    string `json:"guest,omitempty"`
		Username string `json:"username,omitempty"`
		Storage  string `json:"storage,omitempty"`
		File     string `json:"file,omitempty"`
		Device   string `json:"device,omitempty"`
	} `json:"ref"`
	Tags      []string    `json:"tags"`
	Timestamp interface{} `json:"timestamp"`
	Type      string      `json:"type"`
}

func (alert Alert) String() string {
	json, _ := json.MarshalIndent(alert, "", "  ")
	return string(json)
}

// ListAlerts lists all alerts with an optional filter string
func (client *Client) ListAlerts(query string) ([]Alert, error) {
	var alerts []Alert
	path := "alerts"
	if query != "" {
		path += "?" + query
	}
	body, err := client.request("GET", path, nil)
	if err != nil {
		return alerts, err
	}
	err = json.Unmarshal(body, &alerts)
	return alerts, err
}

// GetAlert requests a single Alert by id
func (client *Client) GetAlert(id string) (Alert, error) {
	var alert Alert
	if id == "" {
		return alert, errors.New("name cannot be empty")
	}
	body, err := client.request("GET", "alert/"+id, nil)
	if err != nil {
		return alert, err
	}
	err = json.Unmarshal(body, &alert)
	return alert, err
}

// Acknowledge marks an alert as acknowledged
func (alert *Alert) Acknowledge(client *Client) error {
	if alert.ID == "" {
		return errors.New("id cannot be empty")
	}
	_, err := client.request("POST", "alert/"+alert.ID+"/acknowledge", nil)
	return err
}
