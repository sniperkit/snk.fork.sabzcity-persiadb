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
	"strings"

	"github.com/SabzCity/go-library/errors"
)

// Use MySQL as storage for start. when PersiaOS released we switch to it.

// AddData to key-value storage and index object data if needed.
func AddData(r *DBR, update bool, Priority uint8, ReplicationNum uint8) error {
	// Ready requests to send to storage engines.
	// If multi requestsend as one request it means developer need transaction!
	// !!!!!TODO!!!!! We can't just return error in transaction.
	for i, metaData := range r.MetaData {
		// Check if developer want to index data.
		if metaData.Indexed {
			// parse data by Mime.
		}

		// Send parsed data and metadata to index layer

		// Check priority and choose best storage

		// Check replication number and send object to needed storage.

		// !!!!!TODO
		// MySQL is just to start!
		result, err := DBC.Exec(strings.Replace(`INSERT INTO {Table} (MetaData, Data) VALUES (?, ?);`, "{Table}", metaData.TAGS[0], -1), metaData, r.Data[i])
		if err == driver.ErrBadConn {
			return errors.CantPrepareStatement
		} else if err != nil { // TODO : Check This case for other scenario!
			return errors.ContentAlreadyExist
		}

		// Check storage engine responses
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			// Check error
		}
		if rowsAffected < 1 {
			return errors.StoringDataNotComplete
		}

		// lastInsertID, err := result.LastInsertId()
		// if err != nil {
		//   // Check error
		// }
		// Check lastInsertID if need!!
	}

	return nil
}
