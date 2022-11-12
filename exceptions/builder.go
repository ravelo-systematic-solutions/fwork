package exceptions

//builder eases the generation of
//exceptions by using the Builder
//design pattern.
type builder struct {
	Exception
}

//NewBuilder retrieves an empty
//Exception builder which contains
//no message, code and an empty
//Data list for manipulation
func NewBuilder() *builder {
	return &builder{Exception{
		Data: make([]Data, 0),
	}}
}

//Build Retrieves Exception instance
func (b *builder) Build() *Exception {
	return &b.Exception
}

//Include adds a new piece of atomic
//exception Data to the struct
func (b *builder) Include(data Data) {
	b.Data = append(b.Data, data)
}

//SetMessage overwrites its value
func (b *builder) SetMessage(message Message) {
	b.Message = message
}

//SetCode overwrites its value
func (b *builder) SetCode(code Code) {
	b.Code = code
}

//IsEmpty retrieves if the exception
//contains metadata
func (b *builder) IsEmpty() bool {
	return len(b.Data) == 0
}
