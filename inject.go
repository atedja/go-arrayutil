package arrayutil

type InjectFunc func(interface{}, interface{}) interface{}

/*
Combines all elements of array by applying an operation, specified by
`InjectFunc`.  For each element, the `InjectFunc` is passed a memo value and
the element. The result becomes the memo value for the next iteration.

If `nil` is passed to the `initial` memo value, it will use the first element.

Example:

	func Sum(memo interface{}, e interface{}) interface{} {
		return memo.(int) + e.(int)
	}

	var myArray = []interface{}{1, 2, 3, 4}
	result := Inject(myArray, nil, Sum)     // result is 10
	result := Inject(myArray, 10, Sum)      // result is 20
*/
func Inject(arr []interface{}, initial interface{}, injectFunc InjectFunc) interface{} {
	if arr == nil || injectFunc == nil {
		return arr
	}

	memo := initial
	startIndex := 0
	if memo == nil {
		memo = arr[0]
		startIndex = 1
	}
	for _, v := range arr[startIndex:] {
		memo = injectFunc(memo, v)
	}
	return memo
}
