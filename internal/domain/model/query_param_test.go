package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var invalidQueries = []*QueryParam{&QueryParam{
	StartTime: "2017-01-01 08:59:22",
	EndTime:   "2017-01-01 09:59:22",
},

	&QueryParam{
		Hostname:  "",
		StartTime: "2017-01-01 08:59:22",
		EndTime:   "2017-01-01 09:59:22",
	},
	&QueryParam{
		Hostname:  "  ",
		StartTime: "2017-01-01 08:59:22",
		EndTime:   "2017-01-01 09:59:22",
	},

	&QueryParam{
		Hostname: "hostname",
		EndTime:  "2017-01-01 09:59:22",
	},

	&QueryParam{
		Hostname:  "hostname",
		StartTime: "",
		EndTime:   "2017-01-01 09:59:22",
	},

	&QueryParam{
		Hostname:  "hostname",
		StartTime: " ",
		EndTime:   "2017-01-01 09:59:22",
	},
	&QueryParam{
		Hostname:  "hostname",
		StartTime: "2017-01-01 08:59:22",
	},

	&QueryParam{
		Hostname:  "hostname",
		StartTime: "2017-01-01 08:59:22",
		EndTime:   "",
	},
	&QueryParam{
		Hostname:  "hostname",
		StartTime: "2017-01-01 08:59:22",
		EndTime:   " ",
	},
	&QueryParam{
		Hostname:  "hostname",
		StartTime: "08:59:22",
		EndTime:   "2017-01-01 09:59:22",
	},
	&QueryParam{
		Hostname:  "hostname",
		StartTime: "2017-01-01 08:59:22",
		EndTime:   "2017-01-01",
	},
}

func TestValidateShouldPass(t *testing.T) {
	query := &QueryParam{
		Hostname:  "hostname",
		StartTime: "2017-01-01 08:59:22",
		EndTime:   "2017-01-01 09:59:22",
	}
	err := query.Validate()
	assert.Nil(t, err)
	_, err = query.GetStartTime()
	assert.Nil(t, err)
	_, err = query.GetEndTime()
	assert.Nil(t, err)

}

func TestValidateInvalidQueries(t *testing.T) {
	for _, query := range invalidQueries {
		err := query.Validate()
		assert.NotNil(t, err)
	}
}

func TestValidateInvalidStartTime(t *testing.T) {
	query := &QueryParam{
		Hostname:  "hostname",
		StartTime: "2017-01-01",
		EndTime:   "2017-01-01 09:59:22",
	}
	_, err := query.GetStartTime()
	assert.NotNil(t, err)
}

func TestValidateInvalidEndTime(t *testing.T) {
	query := &QueryParam{
		Hostname:  "hostname",
		StartTime: "2017-01-01 09:59:22",
		EndTime:   "09:59:22",
	}
	_, err := query.GetEndTime()
	assert.NotNil(t, err)
}
