package json

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

// JWriter Easily send common json responses
type JWriter struct {
	responseWriter 	http.ResponseWriter

	statusCode 		int

	// error message to send
	errorMessage 	string

	// data to be json encoded
	json 			interface{}
}

// Status Set the status code for response
func (rw JWriter) Status(statusCode int) JWriter {
	rw.statusCode = statusCode
	return rw
}

// Send Must be called in order to send response
func (rw JWriter) Send() error {
	if rw.statusCode == 0 {
		err := errors.New("status code must be set before calling Send() by calling Status()")
		log.Println(err)
		return err
	}

	var data = rw.json

	if rw.errorMessage != "" {
		data = map[string]string{ "error": rw.errorMessage }
	}

	rw.responseWriter.WriteHeader(rw.statusCode)

	if data == nil {
		// No response body
		return nil
	}

	err := json.NewEncoder(rw.responseWriter).Encode(data)
	if err != nil {
		log.Println(err)
	}
	return err
}

// Json The data to be encoded as json in the response
func (rw JWriter) Json(data interface{}) JWriter {
	rw.json = data
	return rw
}

// MissingParams 400 status {"error": "missing or invalid params"}
func (rw JWriter) MissingParams() JWriter {
	rw.statusCode = 400
	rw.json = map[string]string{"error": "Missing or invalid params"}
	return rw
}

// NotFound 404 status {"error": "Resource not found"}
func (rw JWriter) NotFound() JWriter {
	rw.statusCode = 404
	rw.json = map[string]string{"error": "Resource not found"}
	return rw
}

// Success 200 status {"success": true}
func (rw JWriter) Success() JWriter {
	if rw.statusCode == 0 {
		rw.statusCode = 200
	}

	rw.json = map[string]bool{"success": true}

	return rw
}

// Error Set custom error message {"error": message}
func (rw JWriter) Error(message string) JWriter {
	rw.errorMessage = message
	return rw
}

// GenericError Set generic error message {"error": "An error occurred"}
func (rw JWriter) GenericError() JWriter {
	rw.errorMessage = "An error occurred"
	return rw
}

// JsonParseError 400 status {"error": "JSON badly formatted"}
func (rw JWriter) JsonParseError() JWriter {
	rw.errorMessage = "JSON badly formatted"
	return rw
}

func Response(w http.ResponseWriter) JWriter {
	w.Header().Set("Content-Type", "application/json")
	return JWriter{responseWriter: w}
}