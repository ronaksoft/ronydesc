package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "rony",
}

func main() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
