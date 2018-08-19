package auth

import (
	"net/http"
	"testing"

	"github.com/dscottboggs/attest"
)

func basicAuthTestHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func TestBasicAuth(t *testing.T) {
	test := attest.Test{t}
	username := User("test basicauth user's name")
	testpass := "test basicauth user's password. such strong. much protect."
	test.Handle(CreateNewUser(string(username), testpass))
	rec, req := test.NewRecorder("/")
	req.SetBasicAuth(string(username), testpass)
	basicAuthTestHandler(BasicAuth(rec, req))
	res := rec.Result()
	test.Equals(http.StatusOK, res.StatusCode)
}

func TestBasicAuthFails(t *testing.T) {
	test := attest.Test{t}
	// test that user which is not present is denied
	rec, req := test.NewRecorder("/")
	req.SetBasicAuth("non-user", "irrellevant value")
	basicAuthTestHandler(BasicAuth(rec, req))
	res := rec.Result()
	test.Equals(http.StatusForbidden, res.StatusCode)
	// start over, this time with correct username but incorrect password
	username := User("test basicauth denies user's name")
	testpass := "test basicauth denies user's password. such strong. much protect."
	test.Handle(CreateNewUser(string(username), testpass))
	rec, req = test.NewRecorder("/")
	req.SetBasicAuth(string(username), "not the correct password")
	basicAuthTestHandler(BasicAuth(rec, req))
	res = rec.Result()
	test.Equals(http.StatusForbidden, res.StatusCode)
}
