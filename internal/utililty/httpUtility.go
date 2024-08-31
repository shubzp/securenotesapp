package utililty

import (
	"encoding/json"
	"log"
	"net/http"
)

type JSONResponseWriter struct {
	ResponseWriter http.ResponseWriter
	ResponseBody   map[string]interface{}
}

type IJSONResponseWriter interface {
	AddKeyValue(key string, value interface{})
	WriteResponse() error
}

func MakeResponseBody(w http.ResponseWriter) IJSONResponseWriter {
	return &JSONResponseWriter{
		ResponseWriter: w,
		ResponseBody:   make(map[string]interface{}),
	}
}

func (jsonResponseWriter *JSONResponseWriter) AddKeyValue(key string, value interface{}) {
	jsonResponseWriter.ResponseBody[key] = value
}

func (jsonResponseWriter *JSONResponseWriter) WriteResponse() error {
	respBody, err := json.Marshal(jsonResponseWriter.ResponseBody)
	if err != nil {
		log.Printf("error while marshalling : %v", err)
		return err
	}

	jsonResponseWriter.ResponseWriter.Write(respBody)
	return nil
}
