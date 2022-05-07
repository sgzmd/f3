package rpc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGlobalSearch_Fake(t *testing.T) {
	client, _ := NewClient(nil)
	resp, _ := (*client).GlobalSearch(nil)

	assert.Equal(t, resp.OriginalRequest.GetSearchTerm(), "Маски")
}
