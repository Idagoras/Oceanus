package errno

import "encoding/json"

type Error interface {
	i()
	WithData(data interface{}) Error
	WithID(id string) Error
	ToString() string
	Code() int
	Message() string
	Data() interface{}
}

var _ Error = (*err)(nil)

type err struct {
	code int         `json:"code"`
	msg  string      `json:"msg"`
	data interface{} `json:"data""`
	id   string      `json:"id,omitempty"`
}

func NewError(code int, msg string) Error {
	return &err{
		code: code,
		msg:  msg,
		data: nil,
	}
}

func (e *err) i() {

}

func (e *err) WithData(data interface{}) Error {
	e.data = data
	return e
}

func (e *err) WithID(id string) Error {
	e.id = id
	return e
}

func (e *err) ToString() string {
	raw, _ := json.Marshal(e)
	return string(raw)
}

func (e *err) Code() int {
	return e.code
}

func (e *err) Message() string {
	return e.msg
}

func (e *err) Data() interface{} {
	return e.data
}
