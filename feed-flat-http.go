package getstream

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// get request helper
func (f *FlatFeed) get(path string, signature string) ([]byte, error) {
	res, err := f.request("GET", path, signature, nil)
	return res, err
}

// post request helper
func (f *FlatFeed) post(path string, signature string, payload []byte) ([]byte, error) {
	res, err := f.request("POST", path, signature, payload)
	return res, err
}

// delete request helper
func (f *FlatFeed) del(path string, signature string, payload []byte) error {
	_, err := f.request("DELETE", path, signature, payload)
	return err
}

// request helper
func (f *FlatFeed) request(method, path string, signature string, payload []byte) ([]byte, error) {

	// create url.URL instance with query params
	absURL, err := f.Client.absoluteURL(path)
	if err != nil {
		return nil, err
	}

	// create a new http request
	req, err := http.NewRequest(method, absURL.String(), bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	// set the Auth headers for the http request
	req.Header.Set("Content-Type", "application/json")
	if f.Token() != "" {
		req.Header.Set("Authorization", signature)
	}

	// perform the http request
	resp, err := f.Client.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// read the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// handle the response
	switch {
	case resp.StatusCode/100 == 2: // SUCCESS
		if body != nil {
			return body, nil
		}
		return nil, nil
	default:
		var respErr []byte
		err = json.Unmarshal(respErr, err)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(respErr))
	}
}
