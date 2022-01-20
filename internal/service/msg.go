package service

import (
	"github.com/goccy/go-json"
	"github.com/ronaksoft/ronykit"
)

/*
   This Code is Auto-Generated by RonyDESC. Please DO NOT EDIT.
   Generated at:
*/

var _ ronykit.Message = (*EchoRequest)(nil)

type EchoRequest struct {
	RandomID  string `json:"randomId" `
	RandomInt int32  `json:"randomInt" `
}

func (x *EchoRequest) Marshal() ([]byte, error) {
	return json.Marshal(x)
}

var _ ronykit.Message = (*EchoResponse)(nil)

type EchoResponse struct {
	RandomID  string `json:"randomId" `
	RandomInt int32  `json:"randomInt" `
}

func (x *EchoResponse) Marshal() ([]byte, error) {
	return json.Marshal(x)
}
