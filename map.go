package arrayutil

type MapFunc func(interface{}) interface{}

/**
Invokes func for each element in the array. Creates a new array containing
the values returned by the func.
*/
func Map(arr []interface{}, mapFunc MapFunc) []interface{} {
	result := make([]interface{}, 0)
	for _, v := range arr {
		result = append(result, mapFunc(v))
	}
	return result
}
