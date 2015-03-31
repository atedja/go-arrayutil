package arrayutil

/**
Removes nil values from an array.

Example:

	var arr = []interface{}{1, 2, 3, 4, nil, 5}
	result := Compact(arr)  // [1, 2, 3, 4, 5]
*/
func Compact(arr []interface{}) []interface{} {
	if arr == nil {
		return arr
	}

	result := make([]interface{}, 0, len(arr))
	for _, v := range arr {
		if v != nil {
			result = append(result, v)
		}
	}
	return result
}
