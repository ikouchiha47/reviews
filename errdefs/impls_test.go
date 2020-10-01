package errdefs

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var errTest = errors.New("this is a test")

func TestNotFound(t *testing.T) {
	cases := []struct{
		err error
		expected bool
	}{
		{NotFoundError(errTest), true},
		{errTest, false},
	}

	for _, tcase := range cases {
		ok := IsNotFoundError(tcase.err)
	
		assert.Equal(t, tcase.expected, ok)
	}
}


func TestStatusError(t *testing.T) {
	errTest := NotFoundError(errTest)

	cases := []struct{
		err error
		expectedCode int
		expected bool
	}{
		{NewStatusError(100, errTest), 100, true},
		{errTest, 404, false},
	}

	for _, tcase := range cases {
		ok := IsHttpError(tcase.err)

		assert.Equal(t, tcase.expected, ok)
		assert.Equal(t, tcase.expectedCode, GetCode(tcase.err))
	}
}
