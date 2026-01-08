package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/smtp"

	"main/internal/config"
)

// GenerateToken creates a random hex token of n bytes
func GenerateToken(n int) (string, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

// SendEmail sends a generic email
func SendEmail(toEmail, subject, body string) error {
	host := config.GetEnv("SMTP_HOST", "")
	port := config.GetEnv("SMTP_PORT", "587")
	user := config.GetEnv("SMTP_USER", "")
	pass := config.GetEnv("SMTP_PASS", "")

	if host == "" || user == "" || pass == "" {
		fmt.Printf("Mock Email to %s:\nSubject: %s\n%s\n", toEmail, subject, body)
		return nil
	}

	smtpAddr := fmt.Sprintf("%s:%s", host, port)
	auth := smtp.PlainAuth("", user, pass, host)

	msg := []byte("To: " + toEmail + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/plain; charset=UTF-8\r\n\r\n" +
		body + "\r\n")

	return smtp.SendMail(smtpAddr, auth, user, []string{toEmail}, msg)
}

// SendVerificationEmail sends a simple verification email with a link containing the token
func SendVerificationEmail(toEmail, token string) error {
	baseURL := config.GetEnv("APP_BASE_URL", "http://localhost:8080")
	verifyLink := fmt.Sprintf("%s/api/v1/verify-email?token=%s", baseURL, token)
	body := "Please verify your email by clicking the link: " + verifyLink
	return SendEmail(toEmail, "Verify your email", body)
}

// SendOrderStatusEmail sends an email to the user when order status changes
func SendOrderStatusEmail(toEmail, orderNumber, status string) error {
	subject := fmt.Sprintf("Order Status Update - %s", orderNumber)
	body := fmt.Sprintf("Your order %s status has been updated to: %s\n\nThank you for shopping with us!", orderNumber, status)
	return SendEmail(toEmail, subject, body)
}
