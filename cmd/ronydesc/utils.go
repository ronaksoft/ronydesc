package main

import "os"

func workingDir() string {
	wd, _ := os.Getwd()

	return wd
}
