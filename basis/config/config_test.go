/**
 * Comment: ??
 * User: DongWQ
 * Date: 13-1-11，下午5:01
 * @version 0.1
 */
package config_test

import (
	"testing"
	"config"
	"log"
	"fmt"
)

func TestConfig(t *testing.T) {

	conf, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatal("load err", err)
	}

	fmt.Println(conf)

}

