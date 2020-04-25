package utils

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/pkg/errors"
	"github.com/quelhasu/notes-cli/parser"
	"github.com/spf13/viper"
)

// OpenEditor open the prefered editor in the env
func OpenEditor(filepath string) {
	editor := GoEnvVariable("EDITOR_NOTES_CLI")
	edit := exec.Command(editor, filepath)
	edit.Stdout = os.Stdout
	edit.Stdin = os.Stdin
	edit.Stderr = os.Stderr
	edit.Run()
}

// GoEnvVariable get the env value according given key
func GoEnvVariable(key string) string {
	if env, ok := os.LookupEnv(key); ok {
		return env
	}
	log.Fatal("Cannot locate home directory. Please set $" + key)
	return ""
}

// GoDotEnvVariable get the env value
func GoDotEnvVariable(key string) string {

	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	value, ok := viper.Get(key).(string)

	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	return value
}

// CreateDirIfNotExist create dir if not exist
func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}

// CreateFileIfNotExist create file and append the template if not exist
func CreateFileIfNotExist(homepath string, filename string) (f *os.File) {
	exist, err := Exists(filename)
	f = nil
	if exist {
		f, err = os.OpenFile(filename, os.O_RDONLY, 0755)
	} else {
		f, err = os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
		AppendTemplate(homepath, filename)
	}

	if err != nil {
		panic(err)
	}

	return f
}

// AppendTemplate append template to a given filename and date
func AppendTemplate(homepath string, filename string) {
	in, err := ioutil.ReadFile(homepath + "/template.md")

	in = parser.Parse(in)
	if err != nil {
		errors.Wrap(err, "Cannot find the template, please create template.md in the $HOME_NOTES_CLI dir")
	}

	ioutil.WriteFile(filename, in, 0666)
}

// Exists check if dir or file exist according given path
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
