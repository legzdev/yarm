package yarm_test

import (
	"testing"

	"github.com/legzdev/yarm"
)

func TestGenerateRandomID(t *testing.T) {
	randomID := yarm.GenerateRandomID()
	t.Log("randomID", randomID)
}
