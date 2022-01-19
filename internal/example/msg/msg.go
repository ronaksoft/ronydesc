package msg

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

type SumRequest struct {
	desc.Message

	Val1 int32 `json:"val1"`
	Val2 int32 `json:"val2"`
}

type SumResponse struct {
	desc.Message

	Sum int32 `json:"sum"`
}
