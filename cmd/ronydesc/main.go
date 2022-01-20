package main

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/ronaksoft/ronydesc/codegen"
	"github.com/ronaksoft/ronydesc/codeparse"
	"github.com/spf13/cobra"
)

//go:embed main.gotmpl
var generatorTemplate embed.FS

var RootCmd = &cobra.Command{
	Use: "ronydesc",
	RunE: func(cmd *cobra.Command, args []string) error {
		srcFolderPath, _ := cmd.Flags().GetString("srcFolderPath")
		basePkgPath, _ := cmd.Flags().GetString("basePkgPath")

		var messages, services []string
		err := filepath.Walk(
			srcFolderPath,
			func(path string, info fs.FileInfo, err error) error {
				if info.IsDir() {
					return nil
				}
				if !strings.HasSuffix(info.Name(), ".go") {
					return nil
				}
				m, s, err := codeparse.Extract(path)
				if err != nil {
					return err
				}
				messages = append(messages, m...)
				services = append(services, s...)

				return nil
			},
		)
		if err != nil {
			return err
		}

		fmt.Println("Messages:", messages)
		fmt.Println("Services:", services)

		outName := fmt.Sprint("./gen-rony.go")
		out, err := os.Create(outName)
		if err != nil {
			return err
		}

		cfg := Config{}
		cfg.ImportPath = fmt.Sprintf("%s/%s", strings.TrimRight(basePkgPath, "/"), srcFolderPath)
		cfg.DstPkgName, _ = cmd.Flags().GetString("dstPkgName")
		cfg.DstFolderPath, _ = cmd.Flags().GetString("dstFolderPath")
		cfg.Messages = messages
		cfg.Services = services

		tplBytes, err := generatorTemplate.ReadFile("main.gotmpl")
		if err != nil {
			return err
		}
		err = template.Must(
			template.New("gen-rony").
				Funcs(codegen.TemplateFunctions).
				Parse(string(tplBytes)),
		).
			Execute(out, cfg)
		if err != nil {
			return err
		}

		err = out.Close()
		if err != nil {
			return err
		}

		cmdd := exec.Command("go", "run", outName)
		cmdd.Stderr = os.Stderr

		err = cmdd.Run()
		if err != nil {
			return err
		}

		return os.Remove(outName)
	},
}

func main() {
	RootCmd.Flags().String("basePkgPath", workingDir(), "")
	RootCmd.Flags().String("srcFolderPath", "./schema", "")
	RootCmd.Flags().String("dstFolderPath", "./service", "")
	RootCmd.Flags().String("dstPkgName", "service", "")

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
