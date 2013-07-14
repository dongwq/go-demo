/**
 * Comment: ??
 * User: DongWQ
 * Date: 13-1-18，上午1:53
 * @version 0.1
 */
package main

import (
	"encoding/csv"
	"io"
	"log"
	"strings"
	"os"
)

var text = `aaa,bbb,ccc
ddd,eee, dd
ggg,hhh,iii
`

func main() {
	csvReader := csv.NewReader(strings.NewReader(text))
	// csvReader.TrailingComma = truea

	f , e := os.Create("toXls.csv")
	if e != nil {
		log.Fatalln("create file err", e)
	}
	defer f.Close()

	for {
		record, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}

		//		log.Println(string(a))
		for _, x := range record {
			log.Println(x)
			f.Write([]byte(x))
		}

//		log.Print(record)
		//		log.Print(type( record))

	}
}

