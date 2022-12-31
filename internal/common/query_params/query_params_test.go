package query_params

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFromFileFileNotSpecified(t *testing.T) {
	_, err := FromFile("")
	assert.NotNil(t, err)
}
