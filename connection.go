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

import (
	"database/sql"

	// For layers architecture we have to add MySQL package in here.
	_ "github.com/go-sql-driver/mysql"

	chaparkhane "github.com/SabzCity/ChaparKhane"
	"github.com/SabzCity/go-library/log"
)

// DBC is DataBase connections pools.
var DBC *sql.DB

// Use MySQL as storage for start. when PersiaOS released we switch to it.
// Open first connection to database and keep it in DBC for future use.
func init() {
	// Code Repo must authenticate itself and get token from related services.
	token := ""

	// Check replication number and make connection to needed storages.
	// MySQL Connection is ::> UserName + ":" + Password + "@" + IP + ":" + Port + "/" + DBName + "?parseTime=true&multiStatements=true"
	DBC, err := sql.Open("mysql", token+":"+token+"@"+"/sabzcity?parseTime=true&multiStatements=true")
	if err != nil {
		log.Fatal(chaparkhane.AddInformationToError(DatabaseConnectionError, map[string]interface{}{"ExtraInfo": err}))
	}

	err = DBC.Ping()
	if err != nil {
		log.Fatal(chaparkhane.AddInformationToError(DatabasePingOut, map[string]interface{}{"ExtraInfo": err}))
	}
}
