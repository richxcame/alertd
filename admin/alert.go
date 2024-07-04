package admin

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"time"
)

func (admin *Admin) SendSMS(message string) error {
	url := os.Getenv("GOSMS_BASE_URL")
	apiKey := os.Getenv("GOSMS_API_KEY")
	payload := map[string]interface{}{
		"to":      admin.Phone,
		"text":    message,
		"api_key": apiKey,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		return errors.New("failed to send sms")
	}

	return nil
}
