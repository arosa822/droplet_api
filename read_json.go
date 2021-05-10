package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Devices struct {
	Devices []Device `json:"devices"`
}

type Device struct {
	Name  string `json:"name"`
	State int    `json:"state"`
}

func main() {

	// Open the json file
	jsonFile, err := os.Open("sample_json.json")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("File opened")

	defer jsonFile.Close()

	// read open jsonFile as a byte array
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// initialize users array
	var devices Devices

	json.Unmarshal(byteValue, &devices)

	for i := 0; i < len(devices.Devices); i++ {
		fmt.Println("Device Type: " + devices.Devices[i].Name)
		fmt.Println("state: " + strconv.Itoa(devices.Devices[i].State))
	}

}
