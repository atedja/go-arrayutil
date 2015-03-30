package arrayutil

type MapFunc func(interface{}) interface{}

/*
Invokes `mapFunc` for each element in the array. Creates a new array containing
the values returned by `mapFunc`
Example:
		func MyMapFunc(v interface{}) interface{} {
			return v.(int) * 3
		}
		var myArray = []interface{}{1, 2, 3, 4}
		result := arrayutil.Map(myArray, MyMapFunc) // [3, 6, 9, 12]
*/
func Map(arr []interface{}, mapFunc MapFunc) []interface{} {
	result := make([]interface{}, 0)
	for _, v := range arr {
		result = append(result, mapFunc(v))
	}
	return result
}
