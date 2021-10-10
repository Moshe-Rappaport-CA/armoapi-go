package clientapi

import (
	"encoding/json"
	"fmt"
)

// Result contains the result of calling Request.Do().
type Result struct {
	body       []byte
	err        error
	statusCode int
}

func (r Result) Decode(raw []byte, obj interface{}) error {
	return json.Unmarshal(raw, obj)
}

// Raw returns the raw result.
func (r Result) Raw() ([]byte, error) {
	return r.body, r.err
}

// StatusCode returns the HTTP status code of the request. (Only valid if no
// error was returned.)
func (r Result) StatusCode(statusCode *int) Result {
	*statusCode = r.statusCode
	return r
}

func (r Result) Into(obj interface{}) error {
	if r.err != nil {
		return r.err
	}

	if len(r.body) == 0 {
		return fmt.Errorf("0-length response with status code: %d", r.statusCode)
	}

	err := r.Decode(r.body, obj)
	if err != nil {
		return err
	}

	return nil
}
