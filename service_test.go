package emailserv

import (
	"fmt"
	"github.com/dottics/dutil"
	"github.com/johannesscr/micro/microtest"
	"net/mail"
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
	ms := NewService(Config{UserToken: token})
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
		msg      Message
		exchange *microtest.Exchange
		e        dutil.Error
	}{
		{
			name: "fail message validation",
			msg:  Message{},
			e: &dutil.Err{
				Status: 400,
				Errors: map[string][]string{
					"from":    {"address required"},
					"to":      {"minimum 1 address"},
					"replyTo": {"address required"},
					"subject": {"required"},
					"body":    {"required"},
				},
			},
		},
		//{
		//	name: "fail message marshal",
		//	msg: Message{
		//		From:    mail.Address{Address: "from@mail.service.com"},
		//		To:      []mail.Address{{Address: "to@mail.service.com"}},
		//		ReplyTo: mail.Address{Address: "replyTo@mail.service.com"},
		//		Subject: "test mail subject",
		//		Body:    `<html><head></head><body><div class="forgot-quotes-here></div></body></html>`,
		//	},
		//	e: &dutil.Err{
		//		Errors: map[string][]string{
		//			"marshal": {"cannot marshal"},
		//		},
		//	},
		//},
		//{
		//	name: "fail message send",
		//	e:    nil,
		//},
		{
			name: "fail response decoding",
			msg: Message{
				From:    mail.Address{Address: "from@mail.service.com"},
				To:      []mail.Address{{Address: "to@mail.service.com"}},
				ReplyTo: mail.Address{Address: "replyTo@mail.service.com"},
				Subject: "test mail subject",
				Body:    `<html><head></head><body><div class="proper class"">test</div></body></html>`,
			},
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 403,
					Body:   `{"message":"","data":null,"errors":{"permission":["Please ensure you have permission]}}`,
				},
			},
			e: &dutil.Err{
				Status: 500,
				Errors: map[string][]string{
					"unmarshal": {"unexpected end of JSON input"},
				},
			},
		},
		{
			name: "403 Forbidden",
			msg: Message{
				From:    mail.Address{Address: "from@mail.service.com"},
				To:      []mail.Address{{Address: "to@mail.service.com"}},
				ReplyTo: mail.Address{Address: "replyTo@mail.service.com"},
				Subject: "test mail subject",
				Body:    `<html><head></head><body><div class="proper class"">test</div></body></html>`,
			},
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 403,
					Body:   `{"message":"","data":null,"errors":{"permission":["Please ensure you have permission"]}}`,
				},
			},
			e: &dutil.Err{
				Status: 403,
				Errors: map[string][]string{
					"permission": {"Please ensure you have permission"},
				},
			},
		},
		{
			name: "500 Internal Server Error",
			msg: Message{
				From:    mail.Address{Address: "from@mail.service.com"},
				To:      []mail.Address{{Address: "to@mail.service.com"}},
				ReplyTo: mail.Address{Address: "replyTo@mail.service.com"},
				Subject: "test mail subject",
				Body:    `<html><head></head><body><div class="proper class"">test</div></body></html>`,
			},
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 500,
					Body:   `{"message":"","data":null,"errors":{"internal_server_error":["some unexpected error"]}}`,
				},
			},
			e: &dutil.Err{
				Status: 500,
				Errors: map[string][]string{
					"internal_server_error": {"some unexpected error"},
				},
			},
		},
		{
			name: "200 Success",
			msg: Message{
				From:    mail.Address{Address: "from@mail.service.com"},
				To:      []mail.Address{{Address: "to@mail.service.com"}},
				ReplyTo: mail.Address{Address: "replyTo@mail.service.com"},
				Subject: "test mail subject",
				Body:    `<html><head></head><body><div class="proper class"">test</div></body></html>`,
			},
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 200,
					Body:   `{"message":"email sent successfully","data":null,"errors":null}`,
				},
			},
			e: nil,
		},
	}

	s := NewService(Config{})
	ms := microtest.MockServer(s)
	defer ms.Server.Close()

	for i, tc := range tests {
		name := fmt.Sprintf("%d %s", i, tc.name)
		t.Run(name, func(t *testing.T) {
			// add mock exchange from the service
			if tc.exchange != nil {
				ms.Append(tc.exchange)
			}

			// test and send mail
			//log.Print("MESSAGE:", &tc.msg, tc.msg)
			e := s.SendMail(&tc.msg)
			if e != nil {
				if tc.e == nil {
					t.Errorf("expected error")
				} else if e.Error() != tc.e.Error() {
					t.Errorf("expected error %v got %v", tc.e, e)
				}
			} else if tc.e != nil {
				t.Errorf("expected error %v got nil", tc.e)
			}
		})
	}
}
