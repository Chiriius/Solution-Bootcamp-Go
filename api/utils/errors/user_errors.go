package errorss

import (
	"fmt"
)

type Error struct {
	Code    int
	Message string
}

func (e *Error) Error() string {
	return fmt.Sprintf("code %d: %s", e.Code, e.Message)
}

var ErrorInterfaceDifType = &Error{
	Code:    2002,
	Message: "The interface type of the request is not the expected one",
}

var ErrorUserNotFound = &Error{
	Code:    404,
	Message: "User not found",
}
