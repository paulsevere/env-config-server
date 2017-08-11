package config

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func importSCRIPT() string {
	return "# ~~~ENVY~~~\nsource " + getConfig() + "\n# ~~~ENVY~~~"
}

var IMPORT_SCRIPT_FISH = "# ~~~ENVY~~~\nsource \"~/.envy\"\n# ~~~ENVY~~~"

func GetHome(path string) string {
	return os.Getenv("HOME") + "/" + path
}
func getConfig() string {
	return viper.GetString("defaultConfig")
}
func CheckForConfig() (*os.File, error) {
	shell := os.Getenv("SHELL")
	if shell == "/bin/bash" {
		shellSetup(".bashrc", importSCRIPT())
	} else {
		shellSetup(".config/fish/config.fish", IMPORT_SCRIPT_FISH)
	}
	return os.OpenFile(getConfig(), os.O_CREATE, os.ModePerm)
}

func shellSetup(loc string, script string) {
	bash, err := ioutil.ReadFile(GetHome(loc))
	if err != nil {
		return
	}

	if strings.Contains(string(bash), script) {
		return
	} else {
		ioutil.WriteFile(GetHome(loc), []byte(string(bash)+script), os.ModeAppend)
	}

}
