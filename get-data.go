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

// GetData is func to retrieve data with related metadata.
// args can be any indexed data e.g. UUID, TAGS, ...
func GetData(args ...interface{}) (*DBR, error) {
	response := DBR{}
	return &response, nil
}

// GetDataWithoutMetaData is func to retrieve data without related metadata.
// args can be any indexed data e.g. UUID, TAGS, ...
func GetDataWithoutMetaData(args ...interface{}) (*DBR, error) {
	response := DBR{}
	return &response, nil
}

// GetMetaData is func to retrieve related metadata to object.
// args can be any indexed data e.g. UUID, TAGS, ...
func GetMetaData(args ...interface{}) (*DBR, error) {
	response := DBR{}
	return &response, nil
}
