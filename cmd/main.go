package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "rony",
}

func main() {
	err := RootCmd.Execute()
	if err != nil {
		fmt.Println(err)
	}
}
