// Copyright (C) 2016 Space Monkey, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package golang

import (
	"fmt"

	"bitbucket.org/pkg/inflect"
	"gopkg.in/spacemonkeygo/dbx.v1/ir"
)

type Struct struct {
	Name   string
	Fields []Field
}

type Field struct {
	Name string
	Type string
	Tags []Tag
}

type Tag struct {
	Key   string
	Value string
}

func FieldFromSelectable(selectable ir.Selectable) Field {
	var name string
	switch obj := selectable.(type) {
	case *ir.Model:
		name = inflect.Camelize(obj.Name)
	case *ir.Field:
		name = inflect.Camelize(obj.Name)
	default:
		panic(fmt.Sprintf("unhandled selectable type %T", obj))
	}
	return Field{
		Name: name,
		Type: name,
	}
}

func FieldsFromSelectables(selectables []ir.Selectable) (fields []Field) {
	for _, selectable := range selectables {
		fields = append(fields, FieldFromSelectable(selectable))
	}
	return fields
}
