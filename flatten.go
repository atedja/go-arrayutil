package arrayutil

/**
Returns a new array that is one-dimensional flat.
*/
func Flatten(arr []interface{}) []interface{} {
	result := make([]interface{}, 0)
	for _, v := range arr {
		switch v.(type) {
		case []interface{}:
			fl := Flatten(v.([]interface{}))
			for _, f := range fl {
				result = append(result, f)
			}
		default:
			result = append(result, v)
		}
	}
	return result
}
