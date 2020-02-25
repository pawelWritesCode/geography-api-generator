package generator

//RandomVariables represents all variables required for rendering any available template
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
