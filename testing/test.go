package main

import (
	"encoding/json"
	"fmt"
)

type Ports struct {
	HostPort      string
	Protocol      string
	ContainerPort string
}

func main() {
	fmt.Println("we will test ports here")

	a1 := Ports{"", "", ""}

	content, err := json.Marshal(a1)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))

	portMap := make(map[string]string)

	portMap[a1.ContainerPort+"/"+a1.Protocol] = a1.HostPort
	fmt.Println(portMap)

	container := make(map[string]interface{})

	container["ports"] = portMap
	fmt.Println(container)

}
