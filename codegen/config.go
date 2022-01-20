package codegen

import "github.com/ronaksoft/ronydesc/desc"

type Config struct {
	// DstFolderPath is the base folder which the generated codes will be there.
	DstFolderPath string
	// DstPkgName is the generated package name.
	DstPkgName string
	// Tags is the list of the tags which will be extracted from the schema.
	Tags []string
	// Message is the list of the messages which will be generated
	Messages []desc.IMessage
	// Services is the list of the services which will be generated
	Services []desc.IService
	// ExternalTemplates is the list of template filenames which will be used to generate the code.
	ExternalTemplates []string
	// ExternalTemplatesOnly if is set then codegen does not execute internal templates.
	// This flag has no effect if ExternalTemplates is empty.
	ExternalTemplatesOnly bool
}
