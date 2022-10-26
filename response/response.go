package response

import "fwork/exceptions"

type Void struct{}

//Success is
type Success struct {
	Payload       any    `json:"payload"`
	TransactionId string `json:"transaction_id"`
}

//Exception is a response which contains
//api errors
type Exception struct {
	Payload       exceptions.Exception `json:"payload"`
	TransactionId string               `json:"transaction_id"`
}
