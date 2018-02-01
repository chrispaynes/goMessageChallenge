package parser

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/mail"
	"strings"

	"github.com/jhillyerd/go.enmime"
)

// ReadMessageFile reads an email message file from disk
func ReadMessageFile(file string) ([]byte, error) {
	return ioutil.ReadFile(file)
}

// ParseMail parses a mail message
func ParseMail(msg string) *mail.Message {
	r := strings.NewReader(msg)
	m, err := mail.ReadMessage(r)

	if err != nil {
		log.Fatal(err)
	}

	return m
}

// GetMessage gets an email message with specific properties
func GetMessage(m *mail.Message, keys ...string) (map[string]string, error) {
	mh := m.Header

	msg := map[string]string{}

	body, err := enmime.ParseMIMEBody(m)

	if err != nil {
		log.Fatal(err)
	}

	// dynamically assign header map keys
	for _, key := range keys {
		switch key {
		case "Body":
			if body.IsTextFromHTML {
				msg[key] = fmt.Sprintf(body.HTML)
			} else {
				msg[key] = body.Text
			}
			break
		default:
			msg[key] = mh.Get(key)
		}
	}

	if err := checkForEmptyProps(msg); err != nil {
		return nil, err
	}

	return msg, nil
}

// checkForEmptyProps checks if an email message is missing requested properties
func checkForEmptyProps(msg map[string]string) error {
	missing := []string{}

	for h := range msg {
		if msg[h] == "" {
			missing = append(missing, h)
		}
	}

	if len(missing) > 0 {
		return fmt.Errorf("missing or malformed email header(s): %s", strings.Join(missing, ", "))
	}

	return nil
}
