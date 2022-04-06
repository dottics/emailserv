package emailserv

import (
	"os"
	"testing"
)

func TestNewService(t *testing.T) {
	// set the env vars NewService should automatically get them
	schemeEnv := "EMAIL_SERVICE_SCHEME"
	scheme := "https"
	err := os.Setenv(schemeEnv, scheme)
	if err != nil {
		t.Errorf("unexpected error setting %s: %v", schemeEnv, err)
	}
	hostEnv := "EMAIL_SERVICE_HOST"
	host := "mail.dottics.com"
	err = os.Setenv(hostEnv, host)
	if err != nil {
		t.Errorf("unexpected error setting %s: %v", hostEnv, err)
	}
	token := "my-test-token"
	ms := NewService(token)
	if ms.URL.Scheme != scheme {
		t.Errorf("expected Email Service to have Scheme %s got %s",
			scheme, ms.URL.Scheme,
		)
	}
	if ms.URL.Host != host {
		t.Errorf("expected Email Service to have Host %s got %s",
			host, ms.URL.Host,
		)
	}
}
