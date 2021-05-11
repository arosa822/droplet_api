package jsonUtil

import (
	"droplet_api/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

////////////////////////////////////////////////////////////////////////////////

func ReadJsonFile(path string) []byte {
	// open the specified path
	jsonFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("File Opened")

	defer jsonFile.Close()

	// read open jsonFile as a byte array
	byteValue, _ := ioutil.ReadAll(jsonFile)

	return byteValue
}

func WriteToFile(object *models.Devices, path string) {
	updatedData, _ := json.MarshalIndent(&object, "", "    ")
	err := ioutil.WriteFile(path, updatedData, 0644)

	if err != nil {
		fmt.Println("Unable to write to file")
		panic(err)
	}

}

//func main() {
//
//	// Open the json file
//	jsonFile, err := os.Open("sample_json.json")
//
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	fmt.Println("File opened")
//
//	defer jsonFile.Close()
//
//	// read open jsonFile as a byte array
//	byteValue, _ := ioutil.ReadAll(jsonFile)
//
//	// initialize users array
//	var devices Devices
//
//	json.Unmarshal(byteValue, &devices)
//
//	devices.describeData()
//	devices.Devices[0].updateState(0)
//	fmt.Println("changing device state")
//	devices.describeData()
//
//	updatedData, _ := json.MarshalIndent(devices, "", "    ")
//	err = ioutil.WriteFile("output.json", updatedData, 0644)
//
//}
