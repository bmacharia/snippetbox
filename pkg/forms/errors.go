package forms

// this new type will bs uesed to hold the validation error
// the name of the form field will be the key of this map
type errors map[string][]string

// Implemment ad Add() method to add an error message for a given form field
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Implement a Get() method retrieve the first message for a given filed map
func (e errors) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}
	return es[0]
}
