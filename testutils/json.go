package testutils

import (
	"encoding/json"
	"io"
	"log"
)

// JsonBody contains the Scope's body in JSON format
func JsonToVar(reader io.Reader, body interface{}) {
	err := json.NewDecoder(reader).Decode(body)
	if err != nil {
		log.Fatalln(err)
	}
}
