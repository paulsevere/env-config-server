package env

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

// Variable sdfs
type Variable struct {
	Name  string `json:"name" form:"name" query:"name"`
	Value string `json:"value" form:"value" query:"value"`
}

func New(name string, value string) Variable {
	return Variable{Name: name, Value: value}
}

func (v Variable) Set() {
	cmd := exec.Command("sh", "export", fmt.Sprintf("%v=%v", v.Name, v.Value))
	spew.Dump(cmd)
	err := cmd.Run()
	if err != nil {
		println("%v", err.Error())
	}

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

// func main() {
// 	v := New("MY", "400")
// 	v.Set()
// }
