package admin

import (
	"encoding/json"
	"os"
)

type Admin struct {
	Fullname string `json:"fullname"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

type Admins []Admin

func GetAll(filepath string) (Admins, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var admins Admins
	err = json.Unmarshal(data, &admins)
	if err != nil {
		return nil, err
	}

	return admins, nil
}
