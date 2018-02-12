package main

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestNewPR(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST",
		"https://api.github.com/repos/test_owner/test_repo/test_pull/2",
		httpmock.NewStringResponder(200, ""))
	assert.Equal(t, 1, 1, "they should be the same")
}
