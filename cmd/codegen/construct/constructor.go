package construct

import (
	_"fmt"
	"github.com/a-h/generate"
)

type Constructor struct {
	Schemas []*generate.Schema
	Structs map[string]Struct
	// References getReferences(*generate)
}

type Struct struct {
	Name string
	Fields map[string]Field
	Description string
	AdditionalProperties bool
}

type Field struct {
	Name string
	Type string
	Required bool
	Description string
	OneOf []interface{} 
	AllOf []interface{}
}

func (constructor *Constructor) CreateStruct() error{
	// for _, schema := range constructor.Schemas {
		// a := Field{
		// 	Name:        name,
		// 	JSONName:    "",
		// 	Type:        rootType,
		// 	Required:    false,
		// 	Description: schema.Description,
		// 	OneOf:       nil,
		// 	AllOf:       nil,
		}
	// }



	return nil
}

