package codegen_test

import (
	"testing"

	"github.com/ronaksoft/ronydesc/codegen"
	_ "github.com/ronaksoft/ronydesc/internal/example/ex1"
)

func TestGen(t *testing.T) {
	codegen.Generate()
}
