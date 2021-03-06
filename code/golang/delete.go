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
	"bitbucket.org/pkg/inflect"
	"gopkg.in/spacemonkeygo/dbx.v1/ir"
	"gopkg.in/spacemonkeygo/dbx.v1/sql"
)

type Delete struct {
	Suffix string
	Args   []*Var
	Result *Var
	SQL    string
}

func DeleteFromIR(ir_del *ir.Delete, dialect sql.Dialect) *Delete {
	del := &Delete{
		Suffix: inflect.Camelize(ir_del.Suffix),
		SQL:    sql.RenderDelete(dialect, ir_del),
	}

	if ir_del.One() {
		del.Result = &Var{
			Name: "deleted",
			Type: "bool",
		}
	} else {
		del.Result = &Var{
			Name: "count",
			Type: "int64",
		}
	}

	for _, ir_where := range ir_del.Where {
		if ir_where.Right == nil {
			del.Args = append(del.Args, ArgFromField(ir_where.Left))
		}
	}

	return del
}
