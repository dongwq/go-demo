// Copyright 2013 Kelsey Hightower
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
)

type NetworkInterface struct {
	Name       string   `json:"name"`
	MacAddress string   `json:"macaddress"`
	Addrs      []string `json:"ipaddress"`
}

// NetworkInterfacesToMap returns a mapping of network interface names to
// interface attributes.
func NetworkInterfacesToMap() (networkInterfaceMap map[string]NetworkInterface) {
	networkInterfaceMap = make(map[string]NetworkInterface)

	networkInterfaces, err := net.Interfaces()
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, networkInterface := range networkInterfaces {
		var n NetworkInterface
		n.Name = networkInterface.Name
		n.MacAddress = networkInterface.HardwareAddr.String()

		addrs, _ := networkInterface.Addrs()
		for _, addr := range addrs {
			n.Addrs = append(n.Addrs, addr.String())
		}

		networkInterfaceMap[networkInterface.Name] = n
	}

	return
}

func main() {
	networkInterfaceMap := NetworkInterfacesToMap()

	// Use json.MarshalIndent() so we get pretty json output of our map
	networkInterfaceJSON, err := json.MarshalIndent(networkInterfaceMap, "", "  ")
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(string(networkInterfaceJSON))
}
