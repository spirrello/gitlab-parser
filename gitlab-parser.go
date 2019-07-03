//Used for parsing Gitlab logs

package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"reflect"
)

//getJSONFileData extracts the JSON formated data from the file
func getJSONFileData(apiJSONFileName string) []map[string]interface{} {

	//creat slice for collecting JSON objects
	var jsonFileSlice []map[string]interface{}

	// Open our jsonFile
	apiJSONFile, err := os.Open(apiJSONFileName)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer apiJSONFile.Close()

	scanVar := bufio.NewScanner(apiJSONFile)

	//Loop and append each line to a slice
	for scanVar.Scan() {
		var result map[string]interface{}
		if err := json.Unmarshal(scanVar.Bytes(), &result); err != nil {
			fmt.Println("JSON PARSING FAILED")
		}
		//fmt.Println(reflect.TypeOf(result))
		//fmt.Println("Appending:", result["path"])

		jsonFileSlice = append(jsonFileSlice, result)

	}

	if scanVar.Err() != nil {
		fmt.Println("scanVar.Scan() failed in getJSONFileData")
	}

	return jsonFileSlice

}

func getRequestList(apiJSONFileVar *os.File) {

}

func main() {

	//Get user intput
	apiJSONFileName := flag.String("apijsonlog", "api_json.log", "Secret file in json format, 'user', 'secret'")

	//Extract json data from the file
	apiFileData := getJSONFileData(*apiJSONFileName)

	//Loop through the returned file data
	for _, value := range apiFileData {
		println(value["path"])
		//Print
		fmt.Println(reflect.TypeOf(value))
	}

}
