package codegen

import (
	"embed"
	"os"
	"path/filepath"
	"text/template"

	"github.com/ronaksoft/ronydesc/desc"
)

//go:embed templates/*
var stdTemplatesFS embed.FS

func Generate(cfg Config, services ...desc.Service) error {
	arg := generateTemplateArg(services...)
	arg.PkgName = cfg.DstPkgName

	//dirEntries, err := stdTemplatesFS.ReadDir("templates")
	//if err != nil {
	//	return err
	//}
	//
	//for _, de := range dirEntries {
	//	if de.IsDir() {
	//		continue
	//	}
	//	stdTemplatesFS.Open(de.Name())
	//}

	out, err := os.Create(filepath.Join(cfg.DstFolderPath, cfg.DstFileName))
	if err != nil {
		return err
	}

	t, err := template.ParseFS(stdTemplatesFS, "templates/*.gotmpl")
	if err != nil {
		return err
	}

	err = t.Funcs(TemplateFunctions).Execute(out, arg)
	if err != nil {
		return err
	}

	return out.Close()
}

func generateTemplateArg(services ...desc.Service) *Arg {
	arg := newArg()
	for _, s := range services {
		arg.extractService(s)
	}

	return arg
}
