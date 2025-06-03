package type_

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MyErr struct {
	msg string
}

func (e MyErr) Error() string {
	return e.msg
}

type ErrorInterface interface {
    Error() string
    MyMethod()
}

func TestErrorPackage(t *testing.T) {
	err := MyErr{msg: "error"}
	var tgtErr0 MyErr
	assert.True(t, errors.As(err, &tgtErr0))
	var tgtErr1 ErrorInterface
	assert.False(t, errors.As(err, &tgtErr1))
}