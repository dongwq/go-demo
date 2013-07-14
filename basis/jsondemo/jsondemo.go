/**
 * Comment: ??
 * User: DongWQ
 * Date: 13-1-11，上午11:15
 * @version 0.1
 */
package jsondemo

import (
	"io/ioutil"
	"log"
	"encoding/json"
	"fmt"
)

type Config struct {
	Filename string
	Phone    int64
}

func LoadConfig() Config{
	bs, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal("read err", err)
	}
	var config Config
	json.Unmarshal(bs, &config)

	return config
}

func MakeConfig(){
	 config := Config{Filename:"wef",Phone:32424}

	 jsonConfig,err := json.Marshal(config)
	 if err!= nil{
		 log.Fatal("marshal err",err)
	 }
	 fmt.Print(string(jsonConfig))


}
