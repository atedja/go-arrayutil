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

#### func  Map

```go
func Map(arr []interface{}, mapFunc MapFunc) []interface{}
```
Invokes `mapFunc` for each element in the array. Creates a new array containing
the values returned by `mapFunc` Example:

    func MyMapFunc(v interface{}) interface{} {
    	return v.(int) * 3
    }
    var myArray = []interface{}{1, 2, 3, 4}
    result := arrayutil.Map(myArray, MyMapFunc) // [3, 6, 9, 12]

#### func  Reject

```go
func Reject(arr []interface{}, rejectFunc RejectFunc) []interface{}
```
* Deletes element from the array for which the func returns true. Opposite of
Select()

#### func  Select

```go
func Select(arr []interface{}, selectFunc SelectFunc) []interface{}
```
* Invokes func for each element in the array, deleting elements for which the
func return false. Opposite of Reject()

#### func  Unique

```go
func Unique(arr []interface{}) []interface{}
```
* Returns a new array with duplicates removed.

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
