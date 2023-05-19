package mocktc

import (
	"testing"

	"github.com/taskcluster/taskcluster/v50/clients/client-go/tcindex"
)

type Index struct {
	t *testing.T
}

func NewIndex(t *testing.T) *Index {
	return &Index{
		t: t,
	}
}

/////////////////////////////////////////////////

func (index *Index) FindTask(indexPath string) (*tcindex.IndexedTaskResponse, error) {
	return &tcindex.IndexedTaskResponse{}, nil
}
