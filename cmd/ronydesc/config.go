package main

type Config struct {
	ImportPath string
	DstPkgName string
	// DstFolderPath should be RELATIVE to the root folder of the project.
	DstFolderPath string
	Messages      []string
	Services      []string
}
