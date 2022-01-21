package codegen_test

import (
	"os"
	"testing"

	"github.com/ronaksoft/ronydesc/codegen"
	"github.com/ronaksoft/ronydesc/desc"
	"github.com/ronaksoft/ronydesc/internal/example/schema/ex1"
)

func TestGen(t *testing.T) {
	input := codegen.Config{
		DstFolderPath: "./_hdd/ex1",
		DstPkgName:    "ex1",
		Messages: []desc.IMessage{
			ex1.EchoRequest{}, ex1.EchoResponse{},
		},
		Services: []desc.IService{
			ex1.ServiceA{},
		},
	}
	_ = os.MkdirAll(input.DstFolderPath, os.ModePerm)

	if err := codegen.Generate(input); err != nil {
		t.Fatal(err)
	}
}
