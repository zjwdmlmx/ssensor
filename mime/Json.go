package mime

import (
	"bytes"
	"encoding/json"
	"io"
)

func JSON(reader io.Reader, v interface{}) (err error) {
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(v)
	return
}

func JSONString(v interface{}) (str []byte, err error) {
	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	err = encoder.Encode(v)
	if err != nil {
		return
	}

	str = buffer.Bytes()
	return
}
