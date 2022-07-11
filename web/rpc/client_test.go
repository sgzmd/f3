package rpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGlobalSearch_Fake(t *testing.T) {
	client, _ := NewClient(nil)
	resp, _ := client.GlobalSearch(nil)

	assert.Equal(t, resp.OriginalRequest.GetSearchTerm(), "Маски")
}
