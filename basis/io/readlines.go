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
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	// Make sure we close the file no matter what.
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	// The scanner will scan the file and return a single line for each
	// call to Text(). The line can be empty.
	//
	// The Scan() function will return false when it reaches the end of the
	// file.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// Because there could have been errors during the reading of the file.
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}