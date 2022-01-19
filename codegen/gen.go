package codegen

import (
	"embed"
	"os"
	"path/filepath"
	"text/template"

	"github.com/ronaksoft/ronydesc/internal/gen"
)

//go:embed templates/*
var templates embed.FS

func Generate(cfg Config) error {
	arg := generateTemplateArg()
	arg.PkgName = cfg.DstPkgName

	out, err := os.Create(filepath.Join(cfg.DstFolderPath, cfg.DstFileName))
	if err != nil {
		return err
	}

	t, err := template.ParseFS(templates, "templates/*.gotmpl")
	if err != nil {
		return err
	}

	err = t.Funcs(TemplateFunctions).Execute(out, arg)
	if err != nil {
		return err
	}

	return out.Close()
}

func generateTemplateArg() *Arg {
	arg := newArg()
	gen.ForEachService(
		func(s gen.Service) {
			arg.extractService(s)
		},
	)

	return arg
}
