package templateUtils

import "strings"

//Entity is representation of symfony framework entity
type Entity string

//EntityFU returns Entity value with first letter uppercase
func (e Entity) EntityFU() string {
	return strings.Title(string(e))
}
