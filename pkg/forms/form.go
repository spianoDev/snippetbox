package forms

import (
    "fmt"
    "net/url"
    "strings"
    "unicode/utf8"
)

// Form struct that embeds a url.Values object and errors field
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

// Function to check for required fields
func (f *Form) Required(fields ...string) {
    for _, field := range fields {
        value := f.Get(field)
        if strings.TrimSpace(value) == "" {
            f.Errors.Add(field, "Please fill in required field!")
        }
    }
}

// MaxLength function to ensure the form doesn't have entries that are too long
func (f *Form) MaxLength(field string, d int) {
    value := f.Get(field)
    if value == "" {
        return
    }
    if utf8.RuneCountInString(value) > d {
        f.Errors.Add(field, fmt.Sprintf("&d are too many characters. Reduce and try again.", d))
    }
}

// Permitted Values function to ensure the value entered is allowed
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
    f.Errors.Add(field, "Invalid entry!")
}

func (f *Form) Valid() bool {
    return len(f.Errors) == 0
}
