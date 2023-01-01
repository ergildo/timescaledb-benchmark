package repository

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckSearchByHostnameSqlQuery(t *testing.T) {
	const expected = "select host, max(usage) as max, min(usage) as min from cpu_usage where host = $1 and ts between $2 and $3 group by host"
	assert.Equal(t, expected, query)
}

//TODO add more tests
