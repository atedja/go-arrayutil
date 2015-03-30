package arrayutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnique(t *testing.T) {
	var myArray = []interface{}{1, 2, 3, 3}
	result := Unique(myArray)
	assert.Equal(t, 3, len(result))
	assert.Equal(t, 1, result[0].(int))
	assert.Equal(t, 2, result[1].(int))
	assert.Equal(t, 3, result[2].(int))

	var myArray2 = []interface{}{"something", 1, 1.0, "you know what", 1, "something"}
	result = Unique(myArray2)
	assert.Equal(t, 4, len(result))
	assert.Equal(t, "something", result[0].(string))
	assert.Equal(t, 1, result[1].(int))
	assert.Equal(t, 1.0, result[2].(float64))
	assert.Equal(t, "you know what", result[3].(string))

	assert.Nil(t, Unique(nil))

	var myArray3 = []interface{}{"lone"}
	assert.Equal(t, myArray3, Unique([]interface{}(myArray3)))

	var myArray4 = []interface{}{"lone", "lone"}
	result = Unique(myArray4)
	assert.Equal(t, 1, len(result))
	assert.Equal(t, "lone", myArray4[0].(string))
}
