package httpmethod

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEntryProperties(t *testing.T) {
	assert := assert.New(t)

	// given
	relativeURL := "an_url"
	httpMethod := "a_method"

	// when
	sut := newEntry(relativeURL, httpMethod)

	// then
	assert.Equal(relativeURL, sut.RelativeURL())
	assert.Equal(httpMethod, sut.HTTPMethod())
}
