package arrayutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlatten(t *testing.T) {
	var arr1 = []interface{}{1, 2, 3, 4}       // [1, 2, 3, 4]
	var arr2 = []interface{}{5, 6, 7, arr1}    // [5, 6, 7, [1, 2, 3, 4]]
	var arr3 = []interface{}{8, 9, arr1, arr2} // [8, 9, [1, 2, 3, 4], [5, 6, 7, [1, 2, 3, 4]]]

	result := Flatten(arr3)
	assert.Equal(t, 13, len(result))
	assert.Equal(t, 8, result[0].(int))
	assert.Equal(t, 9, result[1].(int))
	assert.Equal(t, 1, result[2].(int))
	assert.Equal(t, 2, result[3].(int))
	assert.Equal(t, 3, result[4].(int))
	assert.Equal(t, 4, result[5].(int))
	assert.Equal(t, 5, result[6].(int))
	assert.Equal(t, 6, result[7].(int))
	assert.Equal(t, 7, result[8].(int))
	assert.Equal(t, 1, result[9].(int))
	assert.Equal(t, 2, result[10].(int))
	assert.Equal(t, 3, result[11].(int))
	assert.Equal(t, 4, result[12].(int))
}
