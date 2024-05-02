package forms

import (
	"fmt"
	"net/url"
	"strings"
	"unicode/utf8"
)

// I will create a Custom Form struct that embeds a url.Values object
// to hold form data and an Erros field to hold any valildaion errors
// for the form data

type Form struct {
	url.Values
	Errors errors
}

// define a new function that initializes a custom Form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

func (f *Form) MaxLength(field string, d int) {
	value := f.Get(field)
	if value == "" {
		return
	}
	if utf8.RuneCountInString(value) > d {
		f.Errors.Add(field, fmt.Sprintf("This field is too long (maximum is %d characters)", d))
	}
}

func (f *Form) PermittedValues(field string, opts ...string) {
	value := f.Get(field)
	if value == "" {
		return
	}
	for _, opt := range opts {
		if value == opt {
			return
		}
	}
	f.Errors.Add(field, "This field is invalid")
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
