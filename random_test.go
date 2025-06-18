package yarm_test

import (
	"testing"

	"github.com/legzdev/yarm"
)

func TestGenerateRandomSuffix(t *testing.T) {
	suffix := yarm.GenerateRandomSuffix()
	t.Log("suffix:", suffix)
}
