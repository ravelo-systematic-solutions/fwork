package exceptions

type Code string

const (
	DuplicatedKeyCode Code = "fwork_dk"
	KeyNotFoundCode        = "fwork_knf"
	InvalidJsonCode        = "fwork_ij"
)

type Message string

const (
	DuplicatedKeyMessage Message = "duplicated key"
	KeyNotFoundMessage           = "key not found"
	InvalidJsonMessage           = "invalid json"
)
