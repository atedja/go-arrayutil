package arrayutil

/*
Returns a new array that is one-dimensional flat.
Example:
		var arr1 = []interface{}{1, 2, 3, 4}       // [1, 2, 3, 4]
		var arr2 = []interface{}{5, 6, 7, arr1}    // [5, 6, 7, [1, 2, 3, 4]]
    result := arrayutil.Flatten(arr2)          // [5, 6, 7, 1, 2, 3, 4]
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
