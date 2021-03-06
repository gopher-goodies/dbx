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

package ir

type Create struct {
	Suffix string
	Model  *Model
	Raw    bool
}

func (cre *Create) Fields() (fields []*Field) {
	return cre.Model.Fields
}

func (cre *Create) InsertableFields() (fields []*Field) {
	if cre.Raw {
		return cre.Model.Fields
	}
	return cre.Model.InsertableFields()
}

func (cre *Create) AutoInsertableFields() (fields []*Field) {
	if cre.Raw {
		return nil
	}
	return cre.Model.AutoInsertableFields()
}
