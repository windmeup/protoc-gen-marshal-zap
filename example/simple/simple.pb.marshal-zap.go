// Code generated by protoc-gen-marshal-zap. DO NOT EDIT.
//
// source: example/simple/simple.proto

package simple

import (
	zapcore "go.uber.org/zap/zapcore"
)

func (x *SimpleMessage) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	if x == nil {
		return nil
	}

	enc.AddString("message", x.Message)

	enc.AddString("secret_message", "[MASKED]")

	return nil
}
