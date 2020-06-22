package configParser

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
)

func CheckFileIsExist(filename string) (exist bool) {
	exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return
}

func readFileData(path string) (data []byte, err error) {
	if CheckFileIsExist(path) {
		data, err = ioutil.ReadFile(path)
	} else {
		data = nil
		err = fmt.Errorf("File Not Exist. Path:%s", path)
	}
	return
}


//ReadConfig read the server config from toml file
func ReadConfig(path string, conf interface{}) error {
	confData, err := readFileData(path)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = toml.Unmarshal(confData, conf)
	if err != nil {
		fmt.Println("toml decode error: " + err.Error())
		return err
	}
	return nil
}
