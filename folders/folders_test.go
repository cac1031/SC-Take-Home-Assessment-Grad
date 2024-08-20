package folders_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"

	// "github.com/georgechieng-sc/interns-2022/folders"
	"github.com/stretchr/testify/assert"
)

func Test_GetAllFolders(t *testing.T) {
	t.Run("nil FetchFolderRequest", func(t *testing.T) {
		_, err := folders.GetAllFolders(nil)
		if err == nil {
			t.Error("expected an error to occur")
		}
	})

	t.Run("FetchFolderRequest with null UUID", func(t *testing.T) {
		req := &folders.FetchFolderRequest{
			OrgID: uuid.Nil,
		}
		_, err := folders.GetAllFolders(req)
		if err == nil {
			t.Error("expected an error to occur")
		}
	})

	t.Run("Non Existent OrgID passed to FetchFolderRequest", func(t *testing.T) {
		req := &folders.FetchFolderRequest{
			OrgID: uuid.FromStringOrNil("83064af3-bb81-4514-a6d4-afba340825cd"),	// non existent UUID in json file
		}
		res, err := folders.GetAllFolders(req)
		assert.Nil(t, res.Folders)
		assert.Nil(t, err)
	})

	t.Run("FetchFolderRequest containing valid FetchFolderResponse", func(t *testing.T) {
		orgID1 := uuid.FromStringOrNil("5e14afd1-2f3c-467d-aab7-04c7f03ee57a")
		test1 := folders.Folder{
			Id: uuid.FromStringOrNil("83dabadc-7c94-4a55-aa5d-e48204594f3a"),
			Name: "relaxing-wraith",
			OrgId: orgID1,
			Deleted: false,
		}

		orgID2 := uuid.FromStringOrNil("65962443-4f54-45c2-b8a1-4af09c8db8bb")
		test2 := folders.Folder{
			Id: uuid.FromStringOrNil("59c98392-1110-46e3-8c8c-26289f82c518"),
			Name: "obliging-groot",
			OrgId: orgID2,
			Deleted: true,
		}
		var tests = []struct {
			in *folders.FetchFolderRequest
			out *folders.FetchFolderResponse
		}{
			{
				in: &folders.FetchFolderRequest{OrgID: orgID1},
				out: &folders.FetchFolderResponse{Folders: []*folders.Folder{&test1}},
			},
			{
				in:  &folders.FetchFolderRequest{OrgID: orgID2},
				out: &folders.FetchFolderResponse{Folders: []*folders.Folder{&test2}},
			},
		}

		for _, tt := range tests {
			t.Run("Test Case", func(t *testing.T) {
				res, err := folders.GetAllFolders(tt.in)
				assert.Equal(t, tt.out, res)
				assert.Nil(t, err)
			})
		}
	})
}
