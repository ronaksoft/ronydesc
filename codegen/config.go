package codegen

type Config struct {
	// DstPkgName the name of the generated package
	DstPkgName string
	// PkgPath is the package path of the generated package
	DstPkgPath string
	// FolderPath is the folder which the generated codes will be there.
	DstFolderPath string
	DstFileName   string
	// SrcFiles is source files
	SrcFiles []string
}
