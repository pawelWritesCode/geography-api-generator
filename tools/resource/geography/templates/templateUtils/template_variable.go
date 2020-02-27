package templateUtils

//TemplateVariables represents all variables required for rendering any available template
type TemplateVariables struct {
	Entity
	Property
}

//NewTemplateVariables returns new TemplateVariables struct
func NewTemplateVariables(e Entity, p Property) TemplateVariables {
	return TemplateVariables{
		Entity:   e,
		Property: p,
	}
}
