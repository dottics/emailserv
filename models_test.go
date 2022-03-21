package emailserv

import (
	"testing"
)

func TestHeaders_Add(t *testing.T) {
	k := "test-key"
	h := Headers{}
	if len(h[k]) != 0 {
		t.Errorf("expected h[test_key] to have length 0 got %v", len(h[k]))
	}

	h.Add(k, []string{"one", "two"})
	if len(h[k]) == 0 {
		t.Errorf("expected h['test_key'] to not have length 0 got '%v'", len(h[k]))
	}
	if h[k][0] != "one" {
		t.Errorf("expected h['test_key'][0] to be 'one' got '%v'", h[k][0])
	}
	if h[k][1] != "two" {
		t.Errorf("expected h['test_key'][0] to be 'two' got '%v'", h[k][0])
	}

	h.Add(k, []string{"three", "four"})
	if h[k][0] != "three" {
		t.Errorf("expected h['test_key'][0] to be 'three' got '%v'", h[k][0])
	}
	if h[k][1] != "four" {
		t.Errorf("expected h['test_key'][0] to be 'four' got '%v'", h[k][0])
	}

	if h[k][2] != "one" {
		t.Errorf("expected h['test_key'][0] to be 'one' got '%v'", h[k][0])
	}
	if h[k][3] != "two" {
		t.Errorf("expected h['test_key'][0] to be 'two' got '%v'", h[k][0])
	}
}

func TestHeaders_Get(t *testing.T) {
	tests := []struct {
		h Headers
		o string
	}{
		{
			h: Headers{},
			o: "",
		},
		{
			h: Headers{
				"Content-Type": []string{"text/html", "charset=UTF-8"},
			},
			o: "text/html",
		},
		{
			h: Headers{
				"Other": []string{"text/html", "charset=UTF-8"},
			},
			o: "",
		},
	}
	for i, tc := range tests {
		t.Run(string(rune(i)), func(t *testing.T) {
			o := tc.h.Get("Content-Type")
			if o != tc.o {
				t.Errorf("expected '%s' got '%s'", tc.o, o)
			}
		})
	}
}

func TestHeaders_MarshalKey(t *testing.T) {
	tests := []struct {
		h Headers
		o string
	}{
		{
			h: Headers{},
			o: "",
		},
		{
			h: Headers{
				"other": []string{"text/html", "charset=UTF-8"},
			},
			o: "",
		},
		{
			h: Headers{
				"Content-Type": []string{"text/html"},
			},
			o: "text/html",
		},
		{
			h: Headers{
				"Content-Type": []string{"text/html", "charset=UTF-8"},
			},
			o: "text/html; charset=UTF-8",
		},
	}

	for i, tc := range tests {
		t.Run(string(rune(i)), func(t *testing.T) {
			o := tc.h.MarshalKey("Content-Type")
			if o != tc.o {
				t.Errorf("expected '%s' got '%s'", tc.o, o)
			}
		})
	}
}
