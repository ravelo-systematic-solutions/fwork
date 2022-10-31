package response

type Void struct{}

//Success is
type Success struct {
	Payload any `json:"payload"`
}
