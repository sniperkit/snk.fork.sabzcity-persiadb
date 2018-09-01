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
	"database/sql/driver"
)

// ReqAddData : The response structure of "AddData()".
type ReqAddData struct {
	MetaData []*MetaData
	Data     []*Data
}

// AddData to key-value storage and index object data if needed.
// If Object UUID exist, It means developer need object history!
func AddData(logicRequest *ReqAddData) error {
	// Ready requests to send to storage engines.
	// If multi requestsend as one request it means developer need transaction!
	// !!!!!TODO!!!!! We can't just return error inside transaction requests!
	for i, metaData := range logicRequest.MetaData {
		// Check or calculate DataSize

		// Check if we can index data.
		// parse data by MediaType.

		// Send parsed data and metadata to index layer

		// !!!!!TODO
		// MySQL is just to start!
		result, err := DBC.Exec(`INSERT INTO {Table} (MetaData, Data) VALUES (?, ?);`, metaData, logicRequest.Data[i])
		if err == driver.ErrBadConn {
			return CantPrepareStatement
		} else if err != nil { // TODO : Check This case for other scenario!
			return ContentAlreadyExist
		}

		// Check storage engine responses
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			// Check error
		}
		if rowsAffected < 1 {
			return StoringDataNotComplete
		}

		// lastInsertID, err := result.LastInsertId()
		// if err != nil {
		//   // Check error
		// }
		// Check lastInsertID if need!!
	}

	return nil
}
