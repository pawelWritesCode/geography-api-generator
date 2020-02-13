package generator

import "strings"

type RandomVariables struct {
	Entity   Entity
	Property Property
}

//New returns new RandomVariables struct
func New(e Entity, p Property) RandomVariables {
	return RandomVariables{
		Entity:   e,
		Property: p,
	}
}

//EntityFU returns Entity value with first letter uppercase
func (t RandomVariables) EntityFU() string {
	return strings.Title(string(t.Entity))
}

//PropertyFU returns Property value with first letter uppercase
func (t RandomVariables) PropertyFU() string {
	return strings.Title(string(t.Property))
}
