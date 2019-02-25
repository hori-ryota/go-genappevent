// Code generated ; DO NOT EDIT

package domain

import (
	"fmt"
	"strings"

	"github.com/hori-ryota/zaperr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ErrorCode string

const (
	errorUnknown ErrorCode = "Unknown"
)

func (c ErrorCode) String() string {
	return string(c)
}

type Error interface {
	Error() string
	Details() []ErrorDetail

	IsUnknown() bool
}

func newError(source error, code ErrorCode, details ...ErrorDetail) Error {
	return errorImpl{
		source:  source,
		code:    code,
		details: details,
	}
}

func ErrorUnknown(source error, details ...ErrorDetail) Error {
	return newError(source, errorUnknown, details...)
}

type errorImpl struct {
	source  error
	code    ErrorCode
	details []ErrorDetail
}

func (e errorImpl) Error() string {
	return fmt.Sprintf("%s:%s:%s", e.code, e.details, e.source)
}
func (e errorImpl) Details() []ErrorDetail {
	return e.details
}

func (e errorImpl) IsUnknown() bool {
	return e.code == errorUnknown
}

type ErrorDetail struct {
	code ErrorDetailCode
	args []string
}

func newErrorDetail(code ErrorDetailCode, args ...string) ErrorDetail {
	return ErrorDetail{
		code: code,
		args: args,
	}
}

func (e ErrorDetail) String() string {
	return strings.Join(append([]string{e.code.String()}, e.args...), ",")
}

type ErrorDetailCode string

func (c ErrorDetailCode) String() string {
	return string(c)
}

func (e errorImpl) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	zaperr.ToNamedField("sourceError", e.source).AddTo(enc)
	zap.String("code", string(e.code)).AddTo(enc)
	zap.Any("details", e.details)
	return nil
}
