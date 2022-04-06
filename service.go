package emailserv

import (
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
