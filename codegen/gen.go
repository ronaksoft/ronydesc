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

func Generate(cfg Config) error {
	arg := newArg(cfg.DstPkgName)
	for _, m := range cfg.Messages {
		arg.extractMessage(m)
	}
	for _, s := range cfg.Services {
		arg.extractService(s)
	}

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

		_ = os.MkdirAll(cfg.DstFolderPath, os.ModePerm|644)
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

	return nil
}

func getFilename(name string) string {
	return fmt.Sprintf("%s", strings.TrimSuffix(filepath.Base(name), filepath.Ext(name)))
}

func getPkgPathSuffix(base, pkgPath string) string {
	return strings.TrimPrefix(pkgPath, base)
}
