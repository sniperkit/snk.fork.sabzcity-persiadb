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

// DBR is request||response strucrure to||from database
type DBR struct {
	MetaData []*MetaData
	Data     []*Data
}

// MetaData :
type MetaData struct {
	ObjectUUID  string   // UUID of Record
	TAGS        []string // Like tables in RDBMS or folder structure in FS
	DataCreated uint     // Use
	Indexed     bool     //
	MIME        string   //
	Purge       bool     // Set it true if developer want to delete record!
}

// Data :
type Data interface{}
