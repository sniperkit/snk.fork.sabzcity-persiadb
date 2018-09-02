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

// MediaType : MediaType or MimeType
type MediaType struct {
	ID             uint16 // Use ID instead of other data to improve efficiency!
	Name           string // Use as file extension in windows, ...!
	Type           string
	SubType        string
	Description    string
	Reference      string
	RegisteredDate uint
	ApprovedDate   uint
}

func init() {
	// Standrad MimeTypes list can found here
	// http://www.iana.org/assignments/media-types/media-types.xhtml

	// Read media-types folder and all .csv and index all data to use in any way by any service like ChaparKhane.
	// !!!! BE AWARE IN UPDATING .csv files in media-types folder !!!!
	// IANA publish update list in alphabetic order!
	// We can't change our order after assign ID to each type.
}
