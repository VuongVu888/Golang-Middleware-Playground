package error

import (
	"errors"
	"net/http"

	"google.golang.org/grpc/codes"
)

type StatusCode string

const (
	Internal    StatusCode = "Internal"
	Unavailable StatusCode = "Unavailable"
)

var mapGrpcStatusCode = map[StatusCode]codes.Code{
	Internal:    codes.Internal,
	Unavailable: codes.Unavailable,
}

var mapHTTPStatusCode = map[StatusCode]uint32{
	Internal:    http.StatusInternalServerError,
	Unavailable: http.StatusServiceUnavailable,
}

func (s StatusCode) GrpcStatusCode() codes.Code {
	return mapGrpcStatusCode[s]
}

func (s StatusCode) HTTPStatusCode() uint32 {
	return mapHTTPStatusCode[s]
}

type Error struct {
	error
	statusCode StatusCode
}

func (e *Error) StatusCode() StatusCode {
	return e.statusCode
}

func As(err error) (*Error, bool) {
	cerr := &Error{}
	if ok := errors.As(err, &cerr); !ok {
		return nil, false
	}

	return cerr, true
}
