package params

import (
	"net/http"
	"net/url"
	"testing"
)

type Q struct {
	Labels     []string `http:"l"`
	MaxResults int      `http:"max"`
	Exact      bool     `http:"x"`
	Email      string   `kind:"email"`
}

func TestUnpackSuccess(t *testing.T) {
	form, _ := url.ParseQuery("l=hello&email=hoge@hoge.com")
	r := &http.Request{Form: form}
	var q Q
	err := Unpack(r, &q)
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

func TestUnpackValidationFailure(t *testing.T) {
	form, _ := url.ParseQuery("l=hello&email=hogehoge.com")
	r := &http.Request{Form: form}
	var q Q
	err := Unpack(r, &q)
	if err == nil {
		t.Errorf("validation error expected!")
	}
}
