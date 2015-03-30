# arrayutil
--
    import "github.com/atedja/go-arrayutil"


## Usage

#### func  Flatten

```go
func Flatten(arr []interface{}) []interface{}
```
Returns a new array that is one-dimensional flat.

Example:

    var arr1 = []interface{}{1, 2, 3, 4}       // [1, 2, 3, 4]
    var arr2 = []interface{}{5, 6, 7, arr1}    // [5, 6, 7, [1, 2, 3, 4]]
    result := arrayutil.Flatten(arr2)          // [5, 6, 7, 1, 2, 3, 4]

#### func  Inject

```go
func Inject(arr []interface{}, initial interface{}, injectFunc InjectFunc) interface{}
```
Combines all elements of enum by applying an operation, specified by
`InjectFunc`. For each element, the `InjectFunc` is passed a memo value and the
element. The result becomes the memo value for the next iteration.

If `nil` is passed to the `initial` memo yvalue, it will use the first element.

Example:

    func Sum(memo interface{}, e interface{}) interface{} {
    	return memo.(int) + e.(int)
    }

    var myArray = []interface{}{1, 2, 3, 4}
    result := Inject(myArray, nil, Sum)				// result is 10
    result := Inject(myArray, 10, Sum)				// result is 20

#### func  Map

```go
func Map(arr []interface{}, mapFunc MapFunc) []interface{}
```
Invokes `MapFunc` for each element in the array. Creates a new array containing
the values returned by `MapFunc`

Example:

    func MyMapFunc(v interface{}) interface{} {
    	return v.(int) * 3
    }
    var myArray = []interface{}{1, 2, 3, 4}
    result := arrayutil.Map(myArray, MyMapFunc) // [3, 6, 9, 12]

#### func  Reject

```go
func Reject(arr []interface{}, rejectFunc RejectFunc) []interface{}
```
* Invokes `RejectFunc` for each element in the array, deleting elements for
which the function returns true. Opposite of Select().

Example:

    func Even(v interface{}) bool {
    	return v.(int)%2 == 0
    }
    var myArray = []interface{}{1, 2, 3, 4}
    result := Reject(myArray, Even) // [1, 3]

#### func  Select

```go
func Select(arr []interface{}, selectFunc SelectFunc) []interface{}
```
* Invokes `SelectFunc` for each element in the array, keeping elements for which
the function returns true. Opposite of Reject().

Example:

    func Even(v interface{}) bool {
    	return v.(int)%2 == 0
    }
    var myArray = []interface{}{1, 2, 3, 4}
    result := Select(myArray, Even) // [2, 4]

#### func  Unique

```go
func Unique(arr []interface{}) []interface{}
```
* Returns a new array with duplicates removed.

Example:

    var myArray = []interface{}{1, 2, 3, 3, 4}
    result := Unique(myArray) // [1, 2, 3, 4]

#### type InjectFunc

```go
type InjectFunc func(interface{}, interface{}) interface{}
```


#### type MapFunc

```go
type MapFunc func(interface{}) interface{}
```


#### type RejectFunc

```go
type RejectFunc func(interface{}) bool
```


#### type SelectFunc

```go
type SelectFunc func(interface{}) bool
```
