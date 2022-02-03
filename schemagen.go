package main

import (
	"encoding/json"
	"fmt"

	"github.com/juju/juju/apiserver"
	"github.com/juju/rpcreflect"

	"github.com/juju/schemagen/jsonschema"
)

// FacadeSchema describes the jsonschema of the RPC interface for a Facacde
type FacadeSchema struct {
	Name    string
	Version int
	Schema  *jsonschema.Schema
}

// DescribeFacadeSchemas returns the list of available Facades and their Versions
func DescribeFacadeSchemas() []FacadeSchema {
	registry := apiserver.AllFacades()
	facades := registry.List()
	result := make([]FacadeSchema, len(facades))
	for i, facade := range facades {
		result[i].Name = facade.Name
		version := facade.Versions[len(facade.Versions)-1]
		result[i].Version = version
		kind, err := registry.GetType(facade.Name, version)
		if err == nil {
			objtype := rpcreflect.ObjTypeOf(kind)
			result[i].Schema = jsonschema.ReflectFromObjType(objtype)
		}
	}
	return result
}

func main() {
	s := DescribeFacadeSchemas()
	b, _ := json.MarshalIndent(s, "", "  ")
	fmt.Printf("%s\n", b)
}
