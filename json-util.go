package util

import (
	"encoding/json"
	"io"
)

// DecodeJSON Decodes json into a reciving variable
func DecodeJSON(r io.Reader, v interface{}) error {
	err := json.NewDecoder(r).Decode(v)
	if err != nil {
		return err
	}
	return nil
}
