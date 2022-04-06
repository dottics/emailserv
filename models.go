package emailserv

import (
	"net/mail"
	"strings"
)

type Headers mail.Header

// Add appends a slice of values to the beginning of the slice of header
// values.
func (h *Headers) Add(key string, values []string) {
	(*h)[key] = append(values, (*h)[key]...)
}

// Get returns the first value for a given key.
func (h *Headers) Get(key string) string {
	if len((*h)[key]) == 0 {
		return ""
	}
	return (*h)[key][0]
}

// MarshalKey parses all the values for a key into the format needed for
// the header to be included into an email.
func (h *Headers) MarshalKey(key string) string {
	if len((*h)[key]) == 0 {
		return ""
	}
	return strings.Join((*h)[key], "; ")
}

// Message is the structure that is to be used to create and send an email
// the Body specifically should be HTML including a meta tag such as
//
// <html lang="en">
// <head>
// 	<meta charset="UTF-8">
// 	<style>...add your styling here...</style>
// </head>
// <body>
// 	...add your content here...
//  ...all images should be included as embedded image or SVG...
// </body>
// </html>
//
// The API you are building should essentially take the data and use the
// HTML Template Execute functionality built into go to populate the
// Message Body before sending.
type Message struct {
	Headers Headers        `json:"headers"`
	From    mail.Address   `json:"from"`
	To      []mail.Address `json:"to"`
	CC      []mail.Address `json:"cc"`
	BCC     []mail.Address `json:"bcc"`
	ReplyTo mail.Address   `json:"replyTo"`
	Subject string         `json:"subject"`
	Body    string         `json:"body"`
}

func (m *Message) Validate() map[string][]string {
	var errors = make(map[string][]string, 1)
	if m.From.Address == "" {
		errors["from"] = []string{"address required"}
	}
	if len(m.To) == 0 {
		errors["to"] = []string{"minimum 1 address"}
	} else {
		if m.To[0].Address == "" {
			errors["to"] = []string{"address required"}
		}
	}
	if m.ReplyTo.Address == "" {
		errors["replyTo"] = []string{"address required"}
	}
	if m.Subject == "" {
		errors["subject"] = []string{"required"}
	}
	if m.Body == "" {
		errors["body"] = []string{"required"}
	}
	return errors
}
