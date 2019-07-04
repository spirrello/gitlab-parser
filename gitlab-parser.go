//Used for parsing Gitlab logs

package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

//Request struct for parsing requests from api_json.log
type Request struct {
	Time     time.Time `json:"time"`
	Severity string    `json:"severity"`
	Duration float64   `json:"duration"`
	Db       float64   `json:"db"`
	View     float64   `json:"view"`
	Status   int       `json:"status"`
	Method   string    `json:"method"`
	Path     string    `json:"path"`
	Params   struct {
		Filepath string `json:"filepath"`
	} `json:"params"`
	Host string `json:"host"`
	IP   string `json:"ip"`
	Ua   string `json:"ua"`
}

//getJSONFileData extracts the JSON formated data from the file
func getJSONFileData(apiJSONFileName string) { //[]Request {

	//creat slice for collecting JSON objects
	//var jsonFileSlice []map[string]interface{}
	//var jsonFileSlice []Request

	// Open our jsonFile
	//apiJSONFile, err := os.Open(apiJSONFileName)
	apiJSONFile, err := ioutil.ReadFile(apiJSONFileName)

	if err != nil {
		fmt.Println(err)
	}

	for _, line := range bytes.Split(apiJSONFile, []byte{'\n'}) {
		var request Request

		json.Unmarshal([]byte(line), &request)
		//fmt.Println(err)

		fmt.Println(request)
	}

	//return jsonFileSlice

}

func getRequestList(apiJSONFileVar *os.File) {

}

func main() {

	//Get user intput
	apiJSONFileName := flag.String("apijsonlog", "api_json.log", "Secret file in json format, 'user', 'secret'")

	//Extract json data from the file
	//apiFileData := getJSONFileData(*apiJSONFileName)
	getJSONFileData(*apiJSONFileName)

	//COMMENTING OUT FOR NOW
	//Loop through the returned file data
	// for value := range apiFileData {
	// 	println(value.path)
	// 	//Print
	// 	//fmt.Println(reflect.TypeOf(value))
	// }

}
