package main

import (
	"embed"
	"fmt"
	"os"
	"os/exec"
	"text/template"

	"github.com/ronaksoft/ronydesc/codeparse"
	"github.com/spf13/cobra"
)

//go:embed main.gotmpl
var generatorTemplate embed.FS

var RootCmd = &cobra.Command{
	Use: "rony",
	RunE: func(cmd *cobra.Command, args []string) error {
		messages, services, err := codeparse.ExtractMessages("./internal/example/ex1/schema.go")
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
		cfg.BasePkgPath, _ = cmd.Flags().GetString("basePkgPath")
		cfg.SrcFolderPath, _ = cmd.Flags().GetString("srcFolderPath")
		cfg.DstPkgName, _ = cmd.Flags().GetString("dstPkgName")
		cfg.DstFolderPath, _ = cmd.Flags().GetString("dstFolderPath")
		cfg.Messages = messages
		cfg.Services = services

		err = template.Must(template.ParseFS(generatorTemplate, "main.gotmpl")).
			Execute(out, cfg)
		if err != nil {
			return err
		}

		err = out.Close()
		if err != nil {
			return err
		}

		err = exec.Command("go", "run", outName).Run()
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
