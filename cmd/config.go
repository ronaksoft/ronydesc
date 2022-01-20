package main

type Config struct {
	// BasePkgPath is the golang package path of the project. We use it and the SrcFolder to create
	// the package path for the schema packages.
	BasePkgPath string
	// SrcFolderPath should be RELATIVE to the root folder of the project. Otherwise, the outcome is
	// not guaranteed.
	SrcFolderPath string
	// DstPkgName is the name of the package which will be generated. It is useful for generated
	// Golang codes.
	DstPkgName string
	// DstFolderPath should be RELATIVE to the root folder of the project.
	DstFolderPath string
	Messages      []string
	Services      []string
}
