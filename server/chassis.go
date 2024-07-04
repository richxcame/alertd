package server

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func (server *ServerILO) GetChassisStatus() (*ChassisStatus, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr, Timeout: 30 * time.Second}

	req, err := http.NewRequest("GET", server.Host+"/redfish/v1/Chassis/1", nil)
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

	bodyIndex := []byte(body)
	json.Unmarshal(bodyIndex, &bodyIndex)
	fmt.Println(string(bodyIndex))
	var chassisStatus ChassisStatus
	if err := json.Unmarshal(body, &chassisStatus); err != nil {
		return nil, err
	}

	return &chassisStatus, nil
}
