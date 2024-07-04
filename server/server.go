package server

import (
	"encoding/json"
	"os"
)

type ServerILO struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Servers []ServerILO

func GetAll(filepath string) (Servers, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var servers Servers
	err = json.Unmarshal(data, &servers)
	if err != nil {
		return nil, err
	}

	return servers, nil
}
