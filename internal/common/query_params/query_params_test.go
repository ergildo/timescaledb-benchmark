package query_params

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFromFileFileNotSpecified(t *testing.T) {
	_, err := FromFile("")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "query-file not specified")
}

func TestFromFileInvalidPath(t *testing.T) {
	queries, err := FromFile("unknown_file_path")
	assert.NotNil(t, err)
	assert.Nil(t, queries)
	assert.Contains(t, err.Error(), "unable to open unknown_file_path")
}
