package emailserv

import (
	"github.com/dottics/dutil"
	"github.com/johannesscr/micro/msp"
)

// Service create a new Service type to add methods onto the type
type Service msp.Service

// NewService creates a microservice-package instance. The
// microservice-package is an instance that loads the environmental
// variables to be able to connect to the specific microservice. The
// microservice-package contains all the implementations to correctly
// exchange with the microservice.
func NewService(token string) *Service {
	return (*Service)(msp.NewService(token, "email"))
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
	ms := (*msp.Service)(s)
	// set the path of the request
	ms.URL.Path = "/send"
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
	res, e := ms.NewRequest("POST", ms.URL.String(), nil, payload)
	if e != nil {
		return e
	}
	resp := struct {
		Message string              `json:"message"`
		Errors  map[string][]string `json:"errors"`
	}{}
	_, e = ms.Decode(res, &resp)
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
