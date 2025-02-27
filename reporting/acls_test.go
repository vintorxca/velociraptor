package reporting

import (
	"testing"

	"github.com/alecthomas/assert"
	"github.com/stretchr/testify/suite"
	api_proto "www.velocidex.com/golang/velociraptor/api/proto"
	"www.velocidex.com/golang/velociraptor/datastore"
	"www.velocidex.com/golang/velociraptor/file_store/test_utils"
	"www.velocidex.com/golang/velociraptor/paths"
)

type ACLTestSuite struct {
	test_utils.TestSuite
}

func (self *ACLTestSuite) TestNotebookPublicACL() {
	new_notebook := &api_proto.NotebookMetadata{
		NotebookId: "N.12345",
		Creator:    "Creator",
		Public:     true,
	}
	notebook_path_manager := paths.NewNotebookPathManager(new_notebook.NotebookId)
	db, err := datastore.GetDB(self.ConfigObj)
	assert.NoError(self.T(), err)

	err = db.SetSubject(self.ConfigObj, notebook_path_manager.Path(), new_notebook)
	assert.NoError(self.T(), err)

	// Check that everyone has access
	assert.True(self.T(), CheckNotebookAccess(new_notebook, "User1"))

	// Make the notebook not public.
	new_notebook.Public = false

	err = db.SetSubject(self.ConfigObj, notebook_path_manager.Path(), new_notebook)
	assert.NoError(self.T(), err)

	// User1 lost access.
	assert.False(self.T(), CheckNotebookAccess(new_notebook, "User1"))

	// The creator always has access regardless
	assert.True(self.T(), CheckNotebookAccess(new_notebook, "Creator"))

	// Explicitly share with User1
	new_notebook.Collaborators = append(new_notebook.Collaborators, "User1")
	err = db.SetSubject(self.ConfigObj, notebook_path_manager.Path(), new_notebook)
	assert.NoError(self.T(), err)

	err = UpdateShareIndex(self.ConfigObj, new_notebook)
	assert.NoError(self.T(), err)

	// User1 now has access
	assert.True(self.T(), CheckNotebookAccess(new_notebook, "User1"))

	// What notebooks does User1 have access to?
	notebooks, err := GetSharedNotebooks(self.ConfigObj, "User1", 0, 100)
	assert.NoError(self.T(), err)
	assert.Equal(self.T(), 1, len(notebooks))
	assert.Equal(self.T(), new_notebook.NotebookId, notebooks[0].NotebookId)

	// Check GetAllNotebooks without ACL checks
	all_notebooks, err := GetAllNotebooks(self.ConfigObj)
	assert.NoError(self.T(), err)
	assert.Equal(self.T(), 1, len(notebooks))
	assert.Equal(self.T(), new_notebook.NotebookId, all_notebooks[0].NotebookId)

	// test_utils.GetMemoryDataStore(self.T(), self.ConfigObj).Debug()
}

func TestACLs(t *testing.T) {
	suite.Run(t, &ACLTestSuite{})
}
