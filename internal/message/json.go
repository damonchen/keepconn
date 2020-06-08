package message

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
)

// JSONMessage json message format
type JSONMessage struct{}

// Marshal marshal
func (j *JSONMessage) Marshal(v interface{}) (io.Reader, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(data), nil
}

// Unmarshal unmarshal
func (j *JSONMessage) Unmarshal(reader io.Reader, v interface{}) error {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, v)
	return err
}
