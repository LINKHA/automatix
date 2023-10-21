package common

import "fmt"

type Error struct {
	ErrCode int
	ErrMsg  string
}

func (e *Error) Error() string {
	return fmt.Sprintf("Error code: %d, Error message: %s", e.ErrCode, e.ErrMsg)
}

var (
	Asddff = &Error{ErrCode: 1001, ErrMsg: "Dss"}
)

func InitError() {

}
