/**
 * Comment: ??
 * User: DongWQ
 * Date: 13-1-9，上午12:24
 * @version 0.1
 */
package config

import (
	"io/ioutil"
	"log"
	"encoding/json"
)

func LoadConfig(filename string) ( map[string]interface {}, error ) {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("read file err", err)
		return nil, err
	}

//	var config map[string]interface {
//	}
	config:= make(map[string]interface {})
	json.Unmarshal(bs, &config)
	return config , nil
}


