package env

import (
	"os"
	"strings"
)

// Variable sdfs
type Variable struct {
	Name  string `json:"name" form:"name" query:"name"`
	Value string `json:"value" form:"value" query:"value"`
}

func New(name string, value string) Variable {
	return Variable{Name: name, Value: value}
}

// PrintVars asdfs
func CurrentEnv() []Variable {
	vars := make([]Variable, 0)
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		vars = append(vars, New(pair[0], pair[1]))
	}
	return vars
}
