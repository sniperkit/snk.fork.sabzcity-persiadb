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
	"net/http"

	chaparkhane "github.com/SabzCity/ChaparKhane"
)

// Declare Errors PersiaDB Code
const (
	contentAlreadyExist = 40500 + (iota + 1)

	cantPrepareStatement

	storingDataNotComplete

	databaseConnectionError

	databasePingOut
)

// Declare Errors Detials
var (
	ContentAlreadyExist = chaparkhane.NewError("This content was already exist", contentAlreadyExist, http.StatusMethodNotAllowed)

	CantPrepareStatement = chaparkhane.NewError("Can't prepare a new statement to database", cantPrepareStatement, http.StatusInternalServerError)

	StoringDataNotComplete = chaparkhane.NewError("We have some problem in storing your data in our databases. Send your request again! If error exist contact SabzCity platform administrators", storingDataNotComplete, http.StatusInternalServerError)

	DatabaseConnectionError = chaparkhane.NewError("Could not connect to database", databaseConnectionError, http.StatusServiceUnavailable)

	DatabasePingOut = chaparkhane.NewError("Error ocurred in Ping to database", databasePingOut, http.StatusServiceUnavailable)
)
