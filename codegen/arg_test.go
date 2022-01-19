package codegen_test

import (
	"os"
	"testing"

	"github.com/ronaksoft/ronydesc/codegen"
	"github.com/ronaksoft/ronydesc/internal/example/ex1"
	"github.com/ronaksoft/ronydesc/internal/example/ex2"
)

func TestGen(t *testing.T) {
	cfg := codegen.Config{
		DstFileName:   "sample.gen.go",
		DstFolderPath: "./_hdd",
		DstPkgName:    "sample",
	}
	_ = os.MkdirAll("./_hdd", os.ModePerm)

	if err := codegen.Generate(cfg, ex1.ServiceA{}, ex2.ServiceB{}); err != nil {
		t.Fatal(err)
	}
}
