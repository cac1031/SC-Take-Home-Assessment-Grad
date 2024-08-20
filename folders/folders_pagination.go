package folders

import (
	"errors"

	"github.com/gofrs/uuid"
)

const DEFAULT_FETCH = 5 // change this according to page size option

// High Level Ideas
/*
payload will be the org_id uuid + token (cursor) string
if cursor is empty string, we return first x pages, where x is the number of things to return 

since each folder has a unique id, we can use it as the cursor to see where to start next

need to check for invalid/non existent cursors

we will return the cursor to the next item we want to read.

If there are no more items left (the extra read is null) OR we read < paginated items, we return a empty string as token

Potential Extensions: reverse pagination, where we specify a boolean if we want to read the next x folders from a cursor or the previous x folders from a cursor
*/

// Copy over the `GetFolders` and `FetchAllFoldersByOrgID` to get started
func GetPaginatedtFolders(req *FetchFolderRequest, cursor string) (*PaginatedFetchFolderResponse, error) {
	if (req == nil || req.OrgID == uuid.Nil) {
		return nil, errors.New("can't fetch null folder request")
	}

	uuidCursor := uuid.FromStringOrNil(cursor)
	r, next, err := FetchFoldersByOrgID(req.OrgID, uuidCursor)
	
	var fp []*Folder
	fp = append(fp, r...)

	var formattedNext string
	if next == uuid.Nil {
		formattedNext = ""
	} else {
		formattedNext = next.String()
	}

	var ffr *PaginatedFetchFolderResponse = &PaginatedFetchFolderResponse{Folders: fp, Cursor: formattedNext}
	return ffr, err
}

func FetchFoldersByOrgID(orgID uuid.UUID, cursor uuid.UUID) ([]*Folder, uuid.UUID, error) {
	folders, err := GetSampleData()	// this returns everything. Assume this is the database we query from
	if err != nil {
		return nil, cursor, err
	}
	
	resFolder := []*Folder{}
	var validOffset = cursor == uuid.Nil
	var nextHead uuid.UUID
	
	for _, folder := range folders {
		if folder.OrgId == orgID {
			if folder.Id == cursor {
				validOffset = true
			}

			if validOffset {
				if len(resFolder) < DEFAULT_FETCH {
					resFolder = append(resFolder, folder)
				} else {
					nextHead = folder.Id
					break
				}
			}
		}
	}
	return resFolder, nextHead, nil
}
