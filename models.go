package emailserv

import "strings"

type Headers map[string][]string

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

type Message struct {
	Headers Headers  `json:"headers"`
	From    string   `json:"from"`
	To      []string `json:"to"`
	CC      []string `json:"cc"`  // TODO: to be added
	BCC     []string `json:"bcc"` // TODO: to be added
	ReplyTo string   `json:"replyTo"`
	Subject string   `json:"subject"`
	Body    string   `json:"body"`
}
