package arrayutil

type RejectFunc func(interface{}) bool

/*
Invokes `RejectFunc` for each element in the array, deleting elements for which
the function returns true. Opposite of Select().

Example:

	func Even(v interface{}) bool {
		return v.(int)%2 == 0
	}
	var myArray = []interface{}{1, 2, 3, 4}
	result := Reject(myArray, Even)  // [1, 3]
*/
func Reject(arr []interface{}, rejectFunc RejectFunc) []interface{} {
	if arr == nil || rejectFunc == nil {
		return arr
	}

	result := make([]interface{}, 0, len(arr))
	for _, v := range arr {
		if !rejectFunc(v) {
			result = append(result, v)
		}
	}
	return result
}
