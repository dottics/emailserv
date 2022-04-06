package emailserv

import (
	"fmt"
	"github.com/dottics/dutil"
	"github.com/johannesscr/micro/microtest"
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

func TestService_SendMail(t *testing.T) {
	tests := []struct {
		name     string
		msg      *Message
		exchange *microtest.Exchange
		e        dutil.Error
	}{
		{
			name: "fail message validation",
			e:    nil,
		},
		{
			name: "fail message marshal",
			e:    nil,
		},
		{
			name: "fail message send",
			e:    nil,
		},
		{
			name: "fail response decoding",
			e:    nil,
		},
		{
			name: "403 Forbidden",
			e:    nil,
		},
		{
			name: "500 Internal Server Error",
			e:    nil,
		},
		{
			name: "200 Success",
			e:    nil,
		},
	}

	s := NewService("")
	ms := microtest.MockServer(s)
	defer ms.Server.Close()

	for i, tc := range tests {
		name := fmt.Sprintf("%d %s", i, tc.name)
		t.Run(name, func(t *testing.T) {
			e := s.SendMail(tc.msg)
			if e != nil {
				if e.Error() != tc.e.Error() {
					t.Errorf("expected error %v got %v", tc.e, e)
				}
			} else if tc.e != nil {
				t.Errorf("expected error %v got nil", tc.e)
			}
		})
	}
}
