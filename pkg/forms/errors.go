package forms

type errors map[string][]string

func (e errors) Add(field, message string) {
    e[field] = append(e[field], message)
}

// Retrieve the first error message from the map
func (e errors) Get(field string) string {
    es := e[field]
    if len(es) == 0 {
        return "No errors found"
    }
    return es[0]
}


