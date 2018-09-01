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

// ReqQueryData : The request structure of "QueryData()"
type ReqQueryData struct {
	DataVersion []uint16
	offset      uint8
	limit       uint8
	Condition   []struct {
		FieldName         string
		FieldValue        string
		FieldOperator     uint8 // 0:= , 1:> , 2:< , 3:!= , 4:>= , 5:<= , 6:LIKE(regexp) ,
		ConditionOperator uint8 // 0:AND(Must have) , 1:OR(One of 2 condition) , 2:NOT(Must not have)
	}
}

// ResQueryData : The response structure of "QueryData()".
type ResQueryData struct {
	MetaData []*MetaData
	Data     []*Data
}

// QueryData is func to find & retrieve data with related condition in data or metadata.
// conditions can be any indexed data e.g. UUID, TAGS, ...
func QueryData(logicRequest *ReqQueryData) (*ResQueryData, error) {
	logicResponse := ResQueryData{}

	// limit can't be more than 100
	if logicRequest.limit > 100 {
		logicRequest.limit = 100
	}
	// if limit set to zero(0) it means not no limit!! It means just 30 last record!
	if logicRequest.limit == 0 {
		logicRequest.limit = 30
	}

	// We don't return error if no data found, instead we return with empty array

	// Use GetData() to retrive needed objects.

	return &logicResponse, nil
}
