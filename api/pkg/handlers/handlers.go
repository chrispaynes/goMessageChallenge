package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"goMessageChallenge/api/pkg/parser"
	"net/http"
)

// EmailMessage represents a email message with selected headers and a body
type EmailMessage struct {
	Date        string `json:"Date"`
	From        string `json:"From"`
	To          string `json:"To"`
	Subject     string `json:"Subject"`
	MessageID   string `json:"MessageId"`
	ContentType string `json:"ContentType"`
	Body        string `json:"Body"`
}

// // GetEmails gets list of email messages
// func GetEmails(w http.ResponseWriter, req *http.Request) {
// 	fmt.Fprintf(w, "GetEmails")
// }

// // GetEmail get a single email based on Id
// func GetEmail(w http.ResponseWriter, req *http.Request) {
// 	fmt.Fprintf(w, "GetEmail")
// }

// PostEmail receives a POSTed email, parses it and returns specific headers and the body
func PostEmail(w http.ResponseWriter, req *http.Request) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(req.Body)
	newStr := buf.String()

	parsedMail := parser.ParseMail(newStr)
	resp, _ := parser.GetMessage(parsedMail, "Date", "From", "To", "Subject", "Message-ID", "Content-Type", "Body")

	jsonBytes, err := convertMapToJSON(resp)

	_, err = writeJSON(w, jsonBytes, err)

	if err != nil {
		fmt.Printf("server failed to respond to email parse request %s", err)
	}
}

// convertMapToJSON converts the message map into a marshallable struct with field tags
func convertMapToJSON(resp map[string]string) ([]byte, error) {
	msg := EmailMessage{
		Date:        resp["Date"],
		From:        resp["From"],
		To:          resp["To"],
		Subject:     resp["Subject"],
		MessageID:   resp["Message-ID"],
		ContentType: resp["Content-Type"],
		Body:        resp["Body"],
	}

	return json.Marshal(msg)
}

// writeJSON writes an email message JSON server response or an empty email message upon error
func writeJSON(w http.ResponseWriter, b []byte, err error) (int, error) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://ui-gmc.localhost")

	if err != nil {
		emptyObj, _ := json.Marshal(EmailMessage{})
		w.Write(emptyObj)
	}

	return w.Write(b)
}

// GetHealth serves as a simple server health check
func GetHealth(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://ui-gmc.localhost")
	w.WriteHeader(200)
	w.Write([]byte("OK"))
}
