package arrayutil

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func MyMapFunc(v interface{}) interface{} {
	return v.(int) * 3
}

func MyMapFunc2(v interface{}) interface{} {
	var buffer bytes.Buffer
	buffer.WriteString(v.(string))
	buffer.WriteString(" bar")
	return buffer.String()
}

func TestMap(t *testing.T) {
	var myArray = []interface{}{1, 2, 3, 4}
	result := Map(myArray, MyMapFunc)
	assert.Equal(t, 4, len(result))
	assert.Equal(t, 3, result[0].(int))
	assert.Equal(t, 6, result[1].(int))
	assert.Equal(t, 9, result[2].(int))
	assert.Equal(t, 12, result[3].(int))

	var myArray2 = []interface{}{"foo", "can't touch"}
	result = Map(myArray2, MyMapFunc2)
	assert.Equal(t, 2, len(result))
	assert.Equal(t, "foo bar", result[0].(string))
	assert.Equal(t, "can't touch bar", result[1].(string))
}
