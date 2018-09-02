/*
Sniperkit-Bot
- Status: analyzed
*/

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

// ReqDeleteData : The request structure of "DeleteData()"
type ReqDeleteData struct {
	ObjectUUID []string
	// Date uint64
}

// DeleteData is func to delete data with related metadata by ObjectUUID.
func DeleteData(logicRequest *ReqDeleteData) error {
	return nil
}
