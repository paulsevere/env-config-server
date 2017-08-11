package env

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"
)

// Variable sdfs
type Variable struct {
	Name  string `json:"name" form:"name" query:"name"`
	Value string `json:"value" form:"value" query:"value"`
}

type varItem struct {
	Line int
	Var  Variable
}

type varItems []varItem

func (vs varItems) ToVars() []Variable {
	vars := make([]Variable, 0)
	for _, x := range vs {
		vars = append(vars, x.Var)
	}
	return vars
}

func New(name string, value string) Variable {
	return Variable{Name: name, Value: value}
}

func CurrentEnv() []Variable {
	cmd := exec.Command("/usr/bin/printenv")
	vars := make([]Variable, 0)
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		println(err)
		log.Fatal(err)
	}
	for _, pair := range strings.Split(string(stdout[:]), "\n") {
		parts := strings.Split(pair, "=")
		if len(parts) > 1 {
			vars = append(vars, New(parts[0], parts[1]))

		}
		// spew.Dump(Variable{Name: parts[0], Value: parts[1]})

	}
	return vars

}

func getConfig() string {
	return viper.GetString("defaultConfig")
}

func (v Variable) ToString() string {
	return v.Name + " = \"" + v.Value + "\""
}

func SetVar(v Variable) error {
	c, _ := ioutil.ReadFile(getConfig())
	current := string(c)
	for line := range strings.Split(current, "\n") {
		println(line)

	}
	println(current)
	return ioutil.WriteFile(getConfig(), []byte(v.Name+"="+v.Value), os.ModeAppend)

}

func ParseEnv() varItems {
	file, _ := os.Open(getConfig())
	scanner := bufio.NewScanner(file)
	vs := make(varItems, 0)

	lineCount := 1
	for scanner.Scan() {
		pars := strings.Split(scanner.Text(), "=")
		vs = append(vs, varItem{Line: lineCount, Var: New(pars[0], pars[1])})
		lineCount++
	}
	spew.Dump(vs.ToVars())
	return vs
}

// PrintVars asdfs
// func CurrentEnv() []Variable {
// 	vars := make([]Variable, 0)
// 	for _, e := range os.Environ() {
// 		pair := strings.Split(e, "=")
// 		vars = append(vars, New(pair[0], pair[1]))
// 	}
// 	return vars
// }

// func main() {
// 	spew.Dump(CurrentEnv())
// }
