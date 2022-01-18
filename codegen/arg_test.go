package codegen_test

import (
	"fmt"
	"testing"

	"github.com/ronaksoft/ronydesc/codegen"
	_ "github.com/ronaksoft/ronydesc/internal/example/ex1"
)

func TestGen(t *testing.T) {
	arg := codegen.Generate()
	for _, m := range arg.Messages() {
		fmt.Println(m.String())
	}
}
