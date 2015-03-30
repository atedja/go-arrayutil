package arrayutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func MyMapFunc(v interface{}) interface{} {
	return v.(int) * 3
}

func TestMap(t *testing.T) {
	var myArray = []interface{}{1, 2, 3, 4}
	result := Map(myArray, MyMapFunc)
	assert.Equal(t, 4, len(result))
	assert.Equal(t, 3, result[0].(int))
	assert.Equal(t, 6, result[1].(int))
	assert.Equal(t, 9, result[2].(int))
	assert.Equal(t, 12, result[3].(int))
}
