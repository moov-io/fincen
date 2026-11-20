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
