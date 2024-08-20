package folders

import "github.com/gofrs/uuid"

type FetchFolderRequest struct {
	OrgID uuid.UUID
}

type FetchFolderResponse struct {
	Folders []*Folder
}

type PaginatedFetchFolderResponse struct {
	Folders []*Folder
	Cursor string
}
