package json

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)


func ReadJson(r *http.Request, target interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &target)
	if err != nil {
		return err
	}

	return err
}

// ReadJsonLimited https://stackoverflow.com/questions/63378861/gos-http-maxbytesreader-why-pass-in-writer
func ReadJsonLimited(w http.ResponseWriter, r *http.Request, target interface{}, maxSize ...int64) error {

	var maxBytesSize int64 = 1048576 // 1 Mb by default
	if len(maxSize) > 0 {
		maxBytesSize = maxSize[0]
	}

	r.Body = http.MaxBytesReader(w, r.Body, maxBytesSize)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &target)
	if err != nil {
		return err
	}

	return err
}

func ReadResponseJson(r *http.Response, target interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &target)
	if err != nil {
		return err
	}

	return err
}

