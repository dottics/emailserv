package emailserv

import (
	"github.com/dottics/dutil"
	"github.com/johannesscr/micro/msp"
	"net/http"
	"net/url"
)

// Service create a new Service type to add methods onto the type
type Service struct {
	msp.Service
}

type Config struct {
	UserToken string
	APIKey    string
	Header    http.Header
	Values    url.Values
}

// NewService creates a microservice-package instance. The
// microservice-package is an instance that loads the environmental
// variables to be able to connect to the specific microservice. The
// microservice-package contains all the implementations to correctly
// exchange with the microservice.
func NewService(config Config) *Service {
	s := &Service{
		Service: *msp.NewService(msp.Config{
			Name:      "email",
			UserToken: config.UserToken,
			APIKey:    config.APIKey,
			Header:    config.Header,
			Values:    config.Values,
		}),
	}
	return s
}

// SetURL sets the scheme and host of the service. Also makes the service
// a mock-able service with `microtest`.
func (s *Service) SetURL(scheme, host string) {
	s.URL.Scheme = scheme
	s.URL.Host = host
}

// SendMail validates the message then sends the message to the
// email-service with a post request. Returns the errors if there are any
// otherwise returns nil.
func (s *Service) SendMail(msg *Message) dutil.Error {
	// convert from the local service to msp.Service
	// set the path of the request
	s.URL.Path = "/send"
	errors := msg.Validate()

	if len(errors) != 0 {
		e := &dutil.Err{
			Status: 400,
			Errors: errors,
		}
		return e
	}
	payload, e := dutil.MarshalReader(msg)
	if e != nil {
		return e
	}
	res, e := s.DoRequest("POST", s.URL, nil, nil, payload)
	if e != nil {
		return e
	}
	resp := struct {
		Message string              `json:"message"`
		Errors  map[string][]string `json:"errors"`
	}{}
	_, e = msp.Decode(res, &resp)
	if e != nil {
		return e
	}

	if res.StatusCode != 200 {
		e := &dutil.Err{
			Status: res.StatusCode,
			Errors: resp.Errors,
		}
		return e
	}
	return nil
}
