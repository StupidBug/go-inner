package pkg

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringsSplit(t *testing.T) {
	assert.Equal(t, 1, len(strings.Split("", ",")))
}
