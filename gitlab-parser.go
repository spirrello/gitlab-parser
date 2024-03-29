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
func getAPIJSONFileData(apiJSONFileName string) []Request {

	//creat slice for collecting JSON objects
	var jsonFileSlice []Request
	// Open our jsonFile
	apiJSONFile, err := ioutil.ReadFile(apiJSONFileName)
	if err != nil {
		fmt.Println(err)
	}

	//Loop through the file and decode the json objects
	for _, line := range bytes.Split(apiJSONFile, []byte{'\n'}) {
		var request Request
		json.Unmarshal([]byte(line), &request)
		//append the json objects to a slice
		jsonFileSlice = append(jsonFileSlice, request)
	}

	return jsonFileSlice

}

func getRequestList(apiJSONFileVar *os.File) {

}

//apiJsonLogField print particular fields
func apiJSONLogField(jsonItemValue *Request, field *string) {

	switch *field {
	case "severity":
		fmt.Println(jsonItemValue.Severity)
	case "duration":
		fmt.Println(jsonItemValue.Duration)
	case "db":
		fmt.Println(jsonItemValue.Db)
	case "view":
		fmt.Println(jsonItemValue.View)
	case "status":
		fmt.Println(jsonItemValue.Status)
	case "method":
		fmt.Println(jsonItemValue.Method)
	case "path":
		fmt.Println(jsonItemValue.Path)
	case "params":
		fmt.Println(jsonItemValue.Params)
	case "host":
		fmt.Println(jsonItemValue.Host)
	case "ip":
		fmt.Println(jsonItemValue.IP)
	case "ua":
		fmt.Println(jsonItemValue.Ua)
	default:
		fmt.Println(jsonItemValue)
	}

}

func main() {

	//Get user intput
	apiJSONFileName := flag.String("apijsonlog", "api_json.log", "Log file containing API calls.")
	field := flag.String("field", "", "Log file containing API calls.")

	flag.Parse()
	//Extract json data from the file
	apiFileData := getAPIJSONFileData(*apiJSONFileName)

	//Loop through the returned file data
	for jsonItem := range apiFileData {
		//dereference the struct from the slice.  Printing Severity
		jsonItemValue := &apiFileData[jsonItem]

		//print interesting fields
		apiJSONLogField(jsonItemValue, field)

	}

}
