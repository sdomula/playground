package accumulate

const testVersion = 1

// Accumulate returns a new
// collection containing the result of applying that operation to each element of
// the input collection.
func Accumulate(collection []string, operation func(string) string) []string {
	for i, e := range collection {
		collection[i] = operation(e)
	}
	return collection
}
