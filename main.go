package main

import (
	"fmt"

	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
)

func main() {
	req := &folders.FetchFolderRequest{
		OrgID: uuid.FromStringOrNil(folders.DefaultOrgID),
	}

	res, err := folders.GetAllFolders(req)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	folders.PrettyPrint(res)

	// Pagination Implementation //
	resPag1, errPag := folders.GetPaginatedtFolders(req, "")
	if errPag != nil {
		fmt.Printf("%v", errPag)
		return
	}
	folders.PrettyPrint(resPag1)

	// Fetch next set
	resPag2, errPag := folders.GetPaginatedtFolders(req, resPag1.Cursor)
	if errPag != nil {
		fmt.Printf("%v", errPag)
		return
	}
	folders.PrettyPrint(resPag2)

	// When number of results is less than a page size of 5
	resPagEnd, errPag := folders.GetPaginatedtFolders(req, "d01bbaf2-05b5-4993-bdcf-0f6a442fb3e9")
	if errPag != nil {
		fmt.Printf("%v", errPag)
		return
	}
	folders.PrettyPrint(resPagEnd)

}
