package fio

import (
	"io/ioutil"
	"os"

	"github.com/paulsevere/envy/env"
	"github.com/spf13/viper"
)

func getConfig() string {
	return viper.GetString("defaultConfig")
}

func ReadConfig() string {
	contents, _ := ioutil.ReadFile(getConfig())
	return string(contents)
}

func AppendToConfig(v env.Variable) error {
	curr := ReadConfig() + "\n" + v.ToString()

	return ioutil.WriteFile(getConfig(), []byte(curr), os.ModeAppend)

}
