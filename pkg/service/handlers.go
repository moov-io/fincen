package service

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	batch "github.com/moov-io/fincen/pkg/batch"
)

const (
	DisableValidateAttrs = "disable_attrs"
	GenerateAttrs        = "generate_attrs"
	Format               = "format"
	JsonFormat           = "json"
	XmlFormat            = "xml"
)

type Response struct {
	Code    int
	Message string
	Format  string
	Report  *batch.EFilingBatchXML
}

// configure handlers
func ConfigureHandlers(r *mux.Router) error {
	r.HandleFunc("/ping", ping).Methods("GET")
	r.HandleFunc("/validator", validator).Methods("POST")
	r.HandleFunc("/reformat", reformat).Methods("POST")
	return nil
}

// health - health check
func ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("PONG"))
}

func validator(w http.ResponseWriter, r *http.Request) {
	buf, err := getInputFileFromRequest(r)
	if err != nil {
		write(w, Response{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	reportXml, err := batch.CreateReportWithBuffer(buf)
	if err != nil {
		write(w, Response{Code: http.StatusNotImplemented, Message: err.Error()})
		return
	}

	var args []string
	if getFormBoolValue(r, DisableValidateAttrs) {
		args = append(args, "true")
	}

	err = reportXml.Validate(args...)
	if err != nil {
		write(w, Response{Code: http.StatusNotImplemented, Message: err.Error()})
		return
	}

	write(w, Response{Code: http.StatusOK, Message: "valid file"})
}

func reformat(w http.ResponseWriter, r *http.Request) {
	buf, err := getInputFileFromRequest(r)
	if err != nil {
		write(w, Response{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	reportXml, err := batch.CreateReportWithBuffer(buf)
	if err != nil {
		write(w, Response{Code: http.StatusNotImplemented, Message: err.Error()})
		return
	}

	if getFormBoolValue(r, GenerateAttrs) {
		err = reportXml.GenerateAttrs()
	}

	if err != nil {
		write(w, Response{Code: http.StatusNotImplemented, Message: err.Error()})
		return
	}

	format := XmlFormat
	if getFromStringValue(r, Format) == JsonFormat {
		format = JsonFormat
	}

	write(w, Response{Code: http.StatusOK, Format: format, Report: reportXml})
}

func write(w http.ResponseWriter, response Response) {

	w.WriteHeader(response.Code)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if response.Report != nil {
		if response.Format == XmlFormat {
			w.Header().Set("Content-Type", "application/xml; charset=utf-8")
			xml.NewEncoder(w).Encode(response.Report)
		} else {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			json.NewEncoder(w).Encode(response.Report)
		}
		return
	}

	if response.Code == http.StatusOK {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": response.Message,
		})
	} else {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": response.Message,
		})
	}
}

func getInputFileFromRequest(r *http.Request) ([]byte, error) {
	inputFile, _, err := r.FormFile("input")
	if err != nil {
		return nil, err
	}
	defer inputFile.Close()

	var input bytes.Buffer
	if _, err = io.Copy(&input, inputFile); err != nil {
		return nil, err
	}

	return input.Bytes(), nil
}

func getFormBoolValue(r *http.Request, name string) bool {
	value := r.FormValue(name)

	ret, err := strconv.ParseBool(value)
	if err != nil {
		return false
	}

	return ret
}

func getFromStringValue(r *http.Request, name string) string {
	return r.FormValue(name)
}
