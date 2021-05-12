package models

import (
	"fmt"
	"strconv"
)

type Devices struct {
	Devices []Device `json:"devices"`
}

type Device struct {
	Name  string `json:"name"`
	State int    `json:"state"`
}

func (d Devices) GetIndex(targetName string) int {
	for i := 0; i < len(d.Devices); i++ {
		if d.Devices[i].Name == targetName {
			return i
		}
	}
	panic("device not found")
}

func (d Devices) DescribeData() {
	for i := 0; i < len(d.Devices); i++ {
		fmt.Println("Device Type: " + d.Devices[i].Name)
		fmt.Println("state: " + strconv.Itoa(d.Devices[i].State))
	}
}

func (d *Devices) UpdateState(name string, desiredState int) {
	index := d.GetIndex(name)
	d.Devices[index].State = desiredState
}

func (d *Devices) GetDevice(targetName string) Device {
	index := d.GetIndex(targetName)
	return d.Devices[index]
}
