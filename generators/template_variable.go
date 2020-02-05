package generators

import "strings"

type RandomVariables struct {
	Entity Entity
	Property Property
}

//EntityFU returns Entity value with first letter uppercase
func (t RandomVariables) EntityFU() string  {
	return strings.Title(string(t.Entity))
}

//PropertyFU returns Property value with first letter uppercase
func (t RandomVariables) PropertyFU() string  {
	return strings.Title(string(t.Property))
}
