package server

import (
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

func (server *ServerILO) GetSystemsStatus() (*SystemsStatus, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr, Timeout: 30 * time.Second}

	req, err := http.NewRequest("GET", server.Host+"/redfish/v1/Systems/1", nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(server.Username, server.Password)
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var systemsStatus SystemsStatus
	if err := json.Unmarshal(body, &systemsStatus); err != nil {
		return nil, err
	}

	return &systemsStatus, nil
}
