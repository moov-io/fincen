// generated-from:1ec69ce9bb402d9418f7d53d60660184fa7a9b7944fda9751fd8cc32c25791fe DO NOT REMOVE, DO UPDATE

package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
)

func (s TestEnvironment) MakeRequest(method string, target string, body interface{}) *http.Request {
	jsonBody := bytes.Buffer{}
	if body != nil {
		json.NewEncoder(&jsonBody).Encode(body)
	}

	target = strings.ReplaceAll(target, "{accountID}", s.AccountID)
	return httptest.NewRequest(method, target, &jsonBody)
}

func (s TestEnvironment) MakeCall(req *http.Request, body interface{}) *http.Response {
	rec := httptest.NewRecorder()
	s.PublicRouter.ServeHTTP(rec, req)
	res := rec.Result()
	defer res.Body.Close()

	if body != nil {
		json.NewDecoder(res.Body).Decode(&body)
	}

	return res
}
