package arrayutil

type SelectFunc func(interface{}) bool

/**
Invokes func for each element in the array, deleting elements for which the func
return false.
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
