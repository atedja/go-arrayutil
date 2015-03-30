package arrayutil

type RejectFunc func(interface{}) bool

/**
Deletes element from the array for which the func returns true.
Opposite of Select()
*/
func Reject(arr []interface{}, rejectFunc RejectFunc) []interface{} {
	result := make([]interface{}, 0)
	for _, v := range arr {
		if !rejectFunc(v) {
			result = append(result, v)
		}
	}
	return result
}
