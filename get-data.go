// Copyright 2018 SabzCity

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//    http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package persiadb

// Condition :
type Condition struct {
	FieldName         string
	FieldValue        string
	FieldOperator     uint8 // = , > , < , != , >= , <= , LIKE , 
	ConditionOperator uint8 // AND, OR and NOT
}

// GetData is func to retrieve data with related metadata.
// args can be any indexed data e.g. UUID, TAGS, ...
func GetData(Tags []string, args []Condition) (*DBR, error) {
	response := DBR{}
	return &response, nil
}

// GetDataWithoutMetaData is func to retrieve data without related metadata.
// args can be any indexed data e.g. UUID, TAGS, ...
func GetDataWithoutMetaData(Tags []string, args []Condition) (*DBR, error) {
	response := DBR{}
	return &response, nil
}

// GetMetaData is func to retrieve related metadata to object.
// args can be any indexed data e.g. UUID, TAGS, ...
func GetMetaData(Tags []string, args []Condition) (*DBR, error) {
	response := DBR{}
	return &response, nil
}
