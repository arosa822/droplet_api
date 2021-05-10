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

func (d Devices) describeData() {
	for i := 0; i < len(d.Devices); i++ {
		fmt.Println("Device Type: " + d.Devices[i].Name)
		fmt.Println("state: " + strconv.Itoa(d.Devices[i].State))
	}
}

func (d *Device) updateState(desiredState int) {
	d.State = desiredState
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

	devices.describeData()
	devices.Devices[0].updateState(0)
	fmt.Println("changing device state")
	devices.describeData()

	updatedData, _ := json.MarshalIndent(devices, "", "    ")
	err = ioutil.WriteFile("output.json", updatedData, 0644)

}
