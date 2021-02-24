package captcha

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Captcha struct {
	secret string
}

func NewConnector(secret string) *Captcha {
	return &Captcha{
		secret: secret,
	}
}

func (c *Captcha) Verify(token string) (bool, error) {
	var response struct {
		Success bool `json:"success"`
	}
	resp, err := http.Post("https://www.recaptcha.net/recaptcha/api/siteverify",
		"application/x-www-form-urlencoded",
		strings.NewReader(fmt.Sprintf("secret=%s&response=%s", c.secret, token)))
	if err != nil {
		return false, err
	}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return false, err
	}
	err = resp.Body.Close()
	if err != nil {
		return false, err
	}
	return response.Success, nil
}
