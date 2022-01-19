package codegen_test

import (
	"os"
	"testing"

	"github.com/ronaksoft/ronydesc/codegen"
	_ "github.com/ronaksoft/ronydesc/internal/example/ex1"
)

func TestGen(t *testing.T) {
	cfg := codegen.Config{
		DstFileName:   "sample.gen.go",
		DstFolderPath: "./_hdd",
		DstPkgName:    "sample",
	}
	_ = os.MkdirAll("./_hdd", os.ModePerm)

	if err := codegen.Generate(cfg); err != nil {
		t.Fatal(err)
	}
}
