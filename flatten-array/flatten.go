// Package flatten solves the Flatten Array problem from Exercism
package flatten

// Flatten returns a flat list with the values from an arbitrarily-deep nested list-like structure
func Flatten(input interface{}) []interface{} {
	result := make([]interface{}, 0)
	return flatten(input, result)
}

func flatten(input interface{}, result []interface{}) []interface{} {
	if slice, ok := input.([]interface{}); ok {
		for _, x := range slice {
			result = flatten(x, result)
		}
		return result
	}
	if input != nil {
		return append(result, input)
	}
	return result
}
