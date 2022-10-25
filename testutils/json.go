package testutils

import (
	"encoding/json"
	"fwork/exceptions"
	"io"
)

// JsonBody contains the Scope's body in JSON format
func JsonToVar(reader io.Reader, body interface{}) error {
	err := json.NewDecoder(reader).Decode(body)
	if err != nil {
		e := exceptions.NewBuilder()
		e.SetCode(exceptions.InvalidJsonCode)
		e.SetMessage(exceptions.InvalidJsonMessage)

		return e.Exception()
	}

	return nil
}
