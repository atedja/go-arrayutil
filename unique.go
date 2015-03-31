package arrayutil

/*
Returns a new array with duplicates removed.

Example:

	var myArray = []interface{}{1, 2, 3, 3, 4}
	result := Unique(myArray)  // [1, 2, 3, 4]
*/
func Unique(arr []interface{}) []interface{} {
	if arr == nil || len(arr) <= 1 {
		return arr
	}

	result := make([]interface{}, 0, len(arr))
	for _, v := range arr {
		if len(result) == 0 {
			result = append(result, v)
			continue
		}

		duplicate := false
		for _, r := range result {
			if r == v {
				duplicate = true
				break
			}
		}

		if !duplicate {
			result = append(result, v)
		}
	}
	return result
}
