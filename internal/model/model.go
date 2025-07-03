package model

import (
	"go.uber.org/zap/zapcore"

	"go-microservice/internal/constants"
)

type Response struct {
	Message string `json:"message"`
}

// MarshalLogObject implements zap.Logger's object to field encoder
func (r Response) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString(constants.Message, r.Message)
	return nil
}
