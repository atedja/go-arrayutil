package arrayutil

type MapFunc func(interface{}) interface{}

/**
Invokes func for each element in the array, deleting elements for which the func
return false.
*/
func Map(arr []interface{}, mapFunc MapFunc) []interface{} {
	result := make([]interface{}, 0)
	for _, v := range arr {
		result = append(result, mapFunc(v))
	}
	return result
}
