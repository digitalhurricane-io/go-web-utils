package json

import (
	"encoding/json"
	"net/http"
	"fmt"
)

// JSender Easily send common json responses
type JSender struct {
	responseWriter 	http.ResponseWriter

	statusCode 		int

	// error message to send
	errorMessage 	string

	// data to be json encoded
	json 			interface{}
}

// Status Set the status code for response
func (rw JSender) Status(statusCode int) JSender {
	rw.statusCode = statusCode
	return rw
}

// Send Must be called in order to send response
func (rw JSender) Send() {

	if rw.statusCode == 0 && rw.errorMessage != "" {
		rw.statusCode = 400

	} else if rw.statusCode == 0 {
		rw.statusCode = 200
	}

	var data = rw.json

	if rw.errorMessage != "" {
		data = map[string]string{ "error": rw.errorMessage }
	}

	rw.responseWriter.WriteHeader(rw.statusCode)

	if data == nil {
		// No response body
		return
	}

	json.NewEncoder(rw.responseWriter).Encode(data)
}

// Json The data to be encoded as json in the response
func (rw JSender) Json(data interface{}) JSender {
	rw.json = data
	return rw
}

// MissingParams 400 status {"error": "missing or invalid params"}
func (rw JSender) MissingParams() JSender {
	rw.statusCode = 400
	rw.json = map[string]string{"error": "Missing or invalid params"}
	return rw
}

// NotFound 404 status {"error": "Resource not found"}
func (rw JSender) NotFound() JSender {
	rw.statusCode = 404
	rw.json = map[string]string{"error": "Resource not found"}
	return rw
}

// Success 200 status {"success": true}
func (rw JSender) Success() JSender {
	if rw.statusCode == 0 {
		rw.statusCode = 200
	}

	rw.json = map[string]bool{"success": true}

	return rw
}

// Error Set custom error message {"error": message}
// If no status code is set upon calling Send(), status code
// will automatically be set to 500
func (rw JSender) Error(message string) JSender {
	rw.errorMessage = message
	return rw
}

// Errorf Set custom formatted error message {"error": message}
// If no status code is set upon calling Send(), status code
// will automatically be set to 500
func (rw JSender) Errorf(format string, a ...interface{}) JSender {
	rw.errorMessage = fmt.Sprintf(format, a...)
	return rw
}

// InternalError Set generic error message {"error": "An error occurred"}
func (rw JSender) InternalError() JSender {
	rw.statusCode = 500
	rw.errorMessage = "An error occurred"
	return rw
}

// DBError 500 status. {"error": "DB Error"}
func (rw JSender) DBError() JSender {
	rw.statusCode = 500
	rw.errorMessage = "DB Error"
	return rw
}

// JsonParseError 400 status {"error": "JSON badly formatted"}
func (rw JSender) JsonParseError() JSender {
	rw.errorMessage = "JSON badly formatted"
	return rw
}

func Response(w http.ResponseWriter) JSender {
	w.Header().Set("Content-Type", "application/json")
	return JSender{responseWriter: w}
}