package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type Form struct {
	url.Values
	Errors errors
}

func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

func (f *Form) Required(fields ...string) bool {
	errorsCount := len(f.Errors)

	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "Required field: "+field)
		}
	}

	return errorsCount != len(f.Errors)
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

func (f *Form) Get(field string) string {
	return f.Values.Get(field)
}

func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		return false
	}
	return true
}

func (f *Form) MinLength(field string, length int, r *http.Request) bool {
	x := r.Form.Get(field)
	if len(x) < length {
		f.Errors.Add(field, fmt.Sprintf("in nabayad kamtar az %d karakter bashe bashe", length))
		return false
	}
	return true
}

func (f *Form) MaxLength(field string, length int, r *http.Request) bool {
	x := r.Form.Get(field)
	if len(x) > length {
		f.Errors.Add(field, fmt.Sprintf("in nabayad bishtar az %d karakter bashe bashe", length))
		return false
	}
	return true
}
