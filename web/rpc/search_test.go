package rpc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFakeSearch_globalSearch1(t *testing.T) {
	search := FakeSearch{}
	result, _ := search.GlobalSearch("Whatever")

	assert.Equal(t, "Whatever", result.OriginalRequest.GetSearchTerm())
	assert.Equal(t, "Whatever author", result.Entry[0].Author)
}
