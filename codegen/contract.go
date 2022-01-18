package codegen

type Contract struct {
	name   string
	input  Message
	output Message

	restMethod string
	restPath   string
}
