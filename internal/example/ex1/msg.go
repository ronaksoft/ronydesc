package ex1

import "github.com/ronaksoft/ronydesc/desc"

type EchoRequest struct {
	desc.Message

	RandomID  string `json:"randomId"`
	RandomInt int32  `json:"randomInt"`
}

type EchoResponse struct {
	desc.Message

	RandomID  string `json:"randomId"`
	RandomInt int32  `json:"randomInt"`
}
