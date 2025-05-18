package main

import (
  "bytes"
  "encoding/json"
  "fmt"
  "io"
  "log"
  "net/http"
  "net/url"
  "strings"
)

// AuthRequest represents login request payload
type AuthRequest struct {
  Email    string `json:"email"`
  Password string `json:"password"`
}

// AuthResponse represents login response from Eskiz
type AuthResponse struct {
  Message string `json:"message"`
  Data    struct {
    Token string `json:"token"`
  } `json:"data"`
  TokenType string `json:"token_type"`
  Status    bool   `json:"status"`
}

// SMSResponse represents response from SMS send endpoint
type SMSResponse struct {
  Status  bool   `json:"status"`
  Message string `json:"message"`
}

// Get JWT token from Eskiz
func getJWTToken(email, password string) (string, error) {
  url := "https://notify.eskiz.uz/api/auth/login"

  reqBody := AuthRequest{
    Email:    email,
    Password: password,
  }

  jsonBody, err := json.Marshal(reqBody)
  if err != nil {
    return "", fmt.Errorf("failed to marshal auth request: %w", err)
  }

  req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
  if err != nil {
    return "", fmt.Errorf("failed to create auth request: %w", err)
  }

  req.Header.Set("Content-Type", "application/json")

  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    return "", fmt.Errorf("auth request failed: %w", err)
  }
  defer resp.Body.Close()

  var authResp AuthResponse
  if err := json.NewDecoder(resp.Body).Decode(&authResp); err != nil {
    return "", fmt.Errorf("failed to decode auth response: %w", err)
  }

  if authResp.Message != "token_generated" {
    return "", fmt.Errorf("authentication failed: %s", authResp.Message)
  }

  return authResp.Data.Token, nil
}

// Send SMS using Eskiz API
func sendSMSEskiz(phone, message, token, senderID string) error {
  apiURL := "https://notify.eskiz.uz/api/message/sms/send"

  data := url.Values{}
  data.Set("mobile_phone", phone)
  data.Set("message", message)
  data.Set("from", senderID)

  req, err := http.NewRequest("POST", apiURL, strings.NewReader(data.Encode()))
  if err != nil {
    return fmt.Errorf("failed to create SMS request: %w", err)
  }

  req.Header.Set("Authorization", "Bearer "+token)
  req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    return fmt.Errorf("failed to send SMS request: %w", err)
  }
  defer resp.Body.Close()

  bodyBytes, err := io.ReadAll(resp.Body)
  if err != nil {
    return fmt.Errorf("failed to read SMS response body: %w", err)
  }
  body := string(bodyBytes)

  log.Println("Raw SMS API response:", body)

  // If the response is a simple string, just check it:
  if resp.StatusCode != http.StatusOK {
    return fmt.Errorf("failed to send SMS, status: %d, response: %s", resp.StatusCode, body)
  }
  return nil
}

func main() {
  email := "dostonxoshimov2005@gmail.com"
  password := "0LplYwjGjAlf8QE8dTjpjOzkz7NU6xACaylKg72O" // Use your Eskiz password/API key
  phone := "998938860094"                                // recipient phone without +
  message := "Bu Eskiz dan test"
  senderID := "4546" // Your Eskiz sender ID

  token, err := getJWTToken(email, password)
  if err != nil {
    log.Fatal("Failed to get JWT token:", err)
  }

  err = sendSMSEskiz(phone, message, token, senderID)
  if err != nil {
    log.Fatal("Error sending SMS:", err)
  }
}
