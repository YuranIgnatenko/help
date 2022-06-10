package Helper

import (
	"fmt"
	"os"
	"strings"
)

type Help struct{}

//reader file
func (help *Help) read(namefile string) (string, error) {
	file, err := os.ReadFile(namefile)
	var data string
	if err == nil {
		data := fmt.Sprintln(string(file), err)
		return data, err
	} else {
		return data, err
	}
}

//get array strings names methods
func (help *Help) Dir(absPath string, nameModule string) []string {
	data, err := help.read(absPath)
	if err != nil {
		fmt.Errorf("error : ", absPath)
	}
	list := make([]string, 0)
	data = strings.ReplaceAll(data, "\n", " ")
	dataSplit := strings.Split(data, " ")
	for ind, val := range dataSplit {
		val = strings.Trim(val, " ")
		if val == "func" {
			if dataSplit[ind+2] == "*"+nameModule+")" {
				line := strings.Split(dataSplit[ind+3], "(")[0]
				list = append(list, line)
			}
		}
	}
	return list
}

//get verbose array strings names methods
func (help *Help) DirVerbose(absPath string, nameModule string) []string {
	data, err := help.read(absPath)
	if err != nil {
		fmt.Errorf("error : ", absPath)
	}
	list := make([]string, 0)
	data = strings.ReplaceAll(data, "\n", " ")
	dataSplit := strings.Split(data, "func")
	for _, value := range dataSplit {
		val := strings.Split(value, "{")[0]
		val = strings.Trim(val, " ")
		if strings.Index(val, "package") == -1 {
			if strings.Index(val, "*"+nameModule+")") != -1 {
				list = append(list, val)
			}
		}
	}
	return list
}
