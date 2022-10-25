package exceptions

//builder eases the generation of
//exceptions by using the Builder
//design pattern.
type builder struct {
	exception Exception
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

//Exception Retrieves Exception instance
func (b *builder) Exception() *Exception {
	return &b.exception
}

//Include adds a new piece of atomic
//exception Data to the struct
func (b *builder) Include(data Data) {
	b.exception.Data = append(b.exception.Data, data)
}

//SetMessage overwrites its value
func (b *builder) SetMessage(message Message) {
	b.exception.Message = message
}

//SetCode overwrites its value
func (b *builder) SetCode(code Code) {
	b.exception.Code = code
}
