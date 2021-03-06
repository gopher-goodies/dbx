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

package xform

import (
	"gopkg.in/spacemonkeygo/dbx.v1/ast"
	"gopkg.in/spacemonkeygo/dbx.v1/ir"
)

func Transform(ast_root *ast.Root) (root *ir.Root, err error) {
	lookup := newLookup()

	models, err := transformModels(lookup, ast_root.Models)
	if err != nil {
		return nil, err
	}
	models = ir.SortModels(models)

	root = &ir.Root{
		Models: models,
	}

	for _, ast_cre := range ast_root.Creates {
		cre, err := transformCreate(lookup, ast_cre)
		if err != nil {
			return nil, err
		}
		root.Creates = append(root.Creates, cre)
	}

	for _, ast_read := range ast_root.Reads {
		reads, err := transformRead(lookup, ast_read)
		if err != nil {
			return nil, err
		}
		root.Reads = append(root.Reads, reads...)
	}

	for _, ast_update := range ast_root.Updates {
		upd, err := transformUpdate(lookup, ast_update)
		if err != nil {
			return nil, err
		}
		root.Updates = append(root.Updates, upd)
	}

	for _, ast_del := range ast_root.Deletes {
		del, err := transformDelete(lookup, ast_del)
		if err != nil {
			return nil, err
		}
		root.Deletes = append(root.Deletes, del)
	}

	return root, nil
}
