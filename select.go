package arrayutil

type SelectFunc func(interface{}) bool

/**
Invokes `SelectFunc` for each element in the array, keeping elements for which
the function returns true. Opposite of Reject().

Example:

	func Even(v interface{}) bool {
		return v.(int)%2 == 0
	}
	var myArray = []interface{}{1, 2, 3, 4}
	result := Select(myArray, Even)  // [2, 4]
*/
func Select(arr []interface{}, selectFunc SelectFunc) []interface{} {
	result := make([]interface{}, 0)
	for _, v := range arr {
		if selectFunc(v) {
			result = append(result, v)
		}
	}
	return result
}
