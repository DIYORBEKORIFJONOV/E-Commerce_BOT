package until

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type PhoneSetting struct {
	authUrl  string
	smsUrl   string
	email    string
	senderID string
	password string
}

func NewPhoneSetting(authUrl, smsUrl, email, senderId, password string) *PhoneSetting {
	return &PhoneSetting{
		authUrl:  authUrl,
		smsUrl:   smsUrl,
		email:    email,
		senderID: senderId,
		password: password,
	}
}

func (p *PhoneSetting) SendSMS(phone, message string) error {
	token, err := p.getToken(p.email, p.password)
	if err != nil {
		return fmt.Errorf("ошибка получения токена: %w", err)
	}
	if err = p.sendSMSRequest(phone, message, token); err != nil {
		return fmt.Errorf("ошибка отправки SMS: %w", err)
	}

	return nil
}

func (p *PhoneSetting) Generate4DigitCode() string {
	code := rand.Intn(9000) + 1000
	return strconv.Itoa(code)
}

func (p *PhoneSetting) getToken(email, password string) (string, error) {
	payload, _ := json.Marshal(map[string]string{
		"email":    email,
		"password": password,
	})

	req, _ := http.NewRequest("POST", p.authUrl, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		Message string `json:"message"`
		Data    struct {
			Token string `json:"token"`
		} `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}
	if result.Data.Token == "" {
		return "", errors.New("токен не получен: " + result.Message)
	}
	return result.Data.Token, nil
}

func (p *PhoneSetting) sendSMSRequest(phone, message, token string) error {
	form := url.Values{}
	form.Set("mobile_phone", phone)
	form.Set("message", message)
	form.Set("from", p.senderID)

	req, _ := http.NewRequest("POST", p.smsUrl, strings.NewReader(form.Encode()))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var body bytes.Buffer
		body.ReadFrom(resp.Body)
		return fmt.Errorf("код %d, ответ: %s", resp.StatusCode, body.String())
	}
	return nil
}
