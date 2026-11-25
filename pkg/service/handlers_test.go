// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package service

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

func TestWriteSetsContentTypeBeforeWriteHeader(t *testing.T) {
	rec := httptest.NewRecorder()
	write(rec, Response{Code: http.StatusOK, Message: "valid file"})
	require.Equal(t, "application/json; charset=utf-8", rec.Result().Header.Get("Content-Type"))
	require.Equal(t, http.StatusOK, rec.Code)
}

func TestReformatSetsXMLContentType(t *testing.T) {
	sample := filepath.Join("..", "..", "data", "samples", "ctr_batch.xml")
	raw, err := os.ReadFile(sample)
	require.NoError(t, err)

	var body bytes.Buffer
	mw := multipart.NewWriter(&body)
	fw, err := mw.CreateFormFile("input", "ctr_batch.xml")
	require.NoError(t, err)
	_, err = fw.Write(raw)
	require.NoError(t, err)
	require.NoError(t, mw.Close())

	router := mux.NewRouter()
	require.NoError(t, ConfigureHandlers(router))

	req := httptest.NewRequest(http.MethodPost, "/reformat", &body)
	req.Header.Set("Content-Type", mw.FormDataContentType())
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)
	require.Equal(t, "application/xml; charset=utf-8", rec.Result().Header.Get("Content-Type"))
	require.Contains(t, rec.Body.String(), "EFilingBatchXML")
}

func postMultipart(t *testing.T, path string, fields map[string]string, input []byte) *httptest.ResponseRecorder {
	t.Helper()

	var body bytes.Buffer
	mw := multipart.NewWriter(&body)
	if input != nil {
		fw, err := mw.CreateFormFile("input", "input.xml")
		require.NoError(t, err)
		_, err = fw.Write(input)
		require.NoError(t, err)
	}
	for k, v := range fields {
		require.NoError(t, mw.WriteField(k, v))
	}
	require.NoError(t, mw.Close())

	router := mux.NewRouter()
	require.NoError(t, ConfigureHandlers(router))

	req := httptest.NewRequest(http.MethodPost, path, &body)
	req.Header.Set("Content-Type", mw.FormDataContentType())
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	return rec
}

func readSample(t *testing.T, name string) []byte {
	t.Helper()

	raw, err := os.ReadFile(filepath.Join("..", "..", "data", "samples", name))
	require.NoError(t, err)
	return raw
}

func TestHandlersStatusCodes(t *testing.T) {
	valid := readSample(t, "ctr_batch.xml")
	invalid := bytes.Replace(valid, []byte(`TotalAmount="47000"`), []byte(`TotalAmount="1"`), 1)
	require.NotEqual(t, valid, invalid)
	unparseable := []byte("this is not a batch report")

	cases := []struct {
		name   string
		path   string
		fields map[string]string
		input  []byte
		code   int
		body   string
	}{
		{"validator valid", "/validator", nil, valid, http.StatusOK, `"status":"valid file"`},
		{"validator missing input", "/validator", nil, nil, http.StatusBadRequest, `"error"`},
		{"validator unparseable", "/validator", nil, unparseable, http.StatusBadRequest, "invalid input data"},
		{"validator validation failure", "/validator", nil, invalid, http.StatusUnprocessableEntity, "TotalAmount"},
		{"validator disable_attrs", "/validator", map[string]string{DisableValidateAttrs: "true"}, invalid, http.StatusOK, `"status":"valid file"`},
		{"reformat valid", "/reformat", nil, valid, http.StatusOK, "EFilingBatchXML"},
		{"reformat missing input", "/reformat", nil, nil, http.StatusBadRequest, `"error"`},
		{"reformat unparseable", "/reformat", nil, unparseable, http.StatusBadRequest, "invalid input data"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			rec := postMultipart(t, tc.path, tc.fields, tc.input)
			require.Equal(t, tc.code, rec.Code, rec.Body.String())
			require.Contains(t, rec.Body.String(), tc.body)
			if tc.code != http.StatusOK {
				require.Equal(t, "application/json; charset=utf-8", rec.Result().Header.Get("Content-Type"))
			}
		})
	}
}

func TestValidatorRejectsNonMultipart(t *testing.T) {
	router := mux.NewRouter()
	require.NoError(t, ConfigureHandlers(router))

	req := httptest.NewRequest(http.MethodPost, "/validator", bytes.NewReader([]byte("plain body")))
	req.Header.Set("Content-Type", "text/plain")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusBadRequest, rec.Code)
}
