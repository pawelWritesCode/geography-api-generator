package generator

import "strings"

//Property represents property name
type Property string

//PropertyFU returns Property value with first letter uppercase
func (p Property) PropertyFU() string {
	return strings.Title(string(p))
}
