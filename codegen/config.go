package codegen

import "github.com/ronaksoft/ronydesc/desc"

type Config struct {
	// DstFolderPath is the base folder which the generated codes will be there.
	DstFolderPath string
	// DstPkgName is the generated package name.
	DstPkgName string
	// Message is the list of the messages which will be generated
	Messages []desc.IMessage
	// Services is the list of the services which will be generated
	Services []desc.Service
}
