package arrayutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func MySelectFunc(v interface{}) bool {
	return v.(int) != 3
}

func TestSelect(t *testing.T) {
	var myArray = []interface{}{1, 2, 3, 3, 4}
	result := Select(myArray, MySelectFunc)
	assert.Equal(t, 3, len(result))
	assert.Equal(t, 1, result[0].(int))
	assert.Equal(t, 2, result[1].(int))
	assert.Equal(t, 4, result[2].(int))
}
