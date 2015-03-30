package arrayutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Sum(v interface{}, w interface{}) interface{} {
	return v.(int) + w.(int)
}

func TestInject(t *testing.T) {
	var myArray = []interface{}{1, 2, 3, 4}
	result := Inject(myArray, nil, Sum)
	assert.Equal(t, 10, result.(int))

	result = Inject(myArray, 10, Sum)
	assert.Equal(t, 20, result.(int))
}
