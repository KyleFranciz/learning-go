// Package stringsandrunes: small part to understand how strings works and what runes are
// all though these arent runs this is how I can treat strings like an array/slice in go.
// I can take parts of a string that I want and slice out the parts that I want easily, can't do this in python
package stringsandrunes

func SlicesAndRuneExample() string {
	str := "Text Example"
	strExample := str[0:5]

	return strExample
}
