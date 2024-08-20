package folders

import (
	"errors"

	"github.com/gofrs/uuid"
)

/*
High Level Func Explanation:
Given an org_id, we return all the folders belonging to the org_id

Fixes and Improvements:
1. Remove unused variables - index in for loops and block variable declaration
2. Observed that function was returning the same object (duplicates)
	- after some research, I found out that the range function in Go returns a copy of the element at the index, and the copy is always reused. So in the old code, if we do:
	for _, v1 := range f {
    fp = append(fp, &v1)
  }
	&v1 will always be referring to the same memory address.
	To fix this, one way is we can use variable shadowing where we redeclare that copy variable.
	We can also fix it by doing the point below
3. I noticed that there were 2 `for loops doing` similar things of copying the return from FetchAllFoldersByOrgID() so I decided to simplify it by returning the result since FetchAllFoldersByOrgID() already formats it in the data type needed to be returned by GetAllFolders()

4. The error return parameter was not being utilized. I decided to use it by checking for invalid parameters (null) being passed (I modified getSampleData() to check for errors)
*/
func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	if (req == nil || req.OrgID == uuid.Nil) {
		return nil, errors.New("can't fetch null folder request")
	}
	res, err := FetchAllFoldersByOrgID(req.OrgID)

	var ffr *FetchFolderResponse = &FetchFolderResponse{Folders: res}
	return ffr, err
}

func FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {
	folders, err := GetSampleData()
	if err != nil {
		return nil, err
	}
	
	resFolder := []*Folder{}
	for _, folder := range folders {
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder)
		}
	}
	return resFolder, nil
}
