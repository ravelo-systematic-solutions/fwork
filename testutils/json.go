package testutils

import (
	"encoding/json"
	"github.com/ravelo-systematic-solutions/fwork/exceptions"
	"io"
)

//JsonToVar contains the Scope's body in JSON format
func JsonToVar(reader io.Reader, body interface{}) error {
	err := json.NewDecoder(reader).Decode(body)
	if err != nil {
		e := exceptions.NewBuilder()
		e.SetCode(exceptions.ResourceInvalidCode)
		e.SetMessage(exceptions.ResourceInvalidMessage)

		return e.Build()
	}

	return nil
}
