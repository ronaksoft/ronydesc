package codegen

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed templates/*
var stdTemplatesFS embed.FS

// Generate generates the template
func Generate(cfg Config) error {
	if len(cfg.Tags) == 0 {
		cfg.Tags = []string{"json", "proto", "paramName", "sql", "cql"}
	}
	arg := newArg(cfg.DstPkgName, cfg.Tags)
	for _, m := range cfg.Messages {
		arg.extractMessage(m)
	}
	for _, s := range cfg.Services {
		arg.extractService(s)
	}

	_ = os.MkdirAll(cfg.DstFolderPath, os.ModePerm|644)

	if !cfg.ExternalTemplatesOnly || len(cfg.ExternalTemplates) == 0 {
		dirEntries, err := stdTemplatesFS.ReadDir("templates")
		if err != nil {
			return err
		}

		for _, de := range dirEntries {
			if de.IsDir() {
				continue
			}

			t, err := template.New(de.Name()).
				Funcs(TemplateFunctions).
				ParseFS(stdTemplatesFS, fmt.Sprintf("templates/%s", de.Name()))
			if err != nil {
				return err
			}

			out, err := os.Create(filepath.Join(cfg.DstFolderPath, getFilename(de.Name())))
			if err != nil {
				return err
			}
			err = t.Execute(out, arg)
			if err != nil {
				return err
			}

			_ = out.Close()
		}
	}

	for _, templateFile := range cfg.ExternalTemplates {
		t, err := template.New(templateFile).Funcs(TemplateFunctions).ParseFiles(templateFile)
		if err != nil {
			return err
		}

		out, err := os.Create(filepath.Join(cfg.DstFolderPath, getFilename(templateFile)))
		if err != nil {
			return err
		}
		err = t.Execute(out, arg)
		if err != nil {
			return err
		}

		_ = out.Close()
	}

	return nil
}

func getFilename(name string) string {
	return fmt.Sprintf("gen.%s", strings.TrimSuffix(filepath.Base(name), filepath.Ext(name)))
}

func getPkgPathSuffix(base, pkgPath string) string {
	return strings.TrimPrefix(pkgPath, base)
}
