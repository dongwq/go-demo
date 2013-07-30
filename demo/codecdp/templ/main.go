package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

type Person struct {
	UserName string
}

func genDao() {
	/*

		So, name your file the same as the template object, (probably not generally practical) or else use ExecuteTemplate instead of just Execute.
	*/

	t := template.New("CPfT005Dao.java")

	t, err := t.ParseFiles("../CPfT005Dao.java")

	if err != nil {
		log.Fatalln(err)
	}
	p := Person{UserName: "Astaxie"}
	fmt.Println(t)
	err = t.Execute(os.Stdout, p)

	log.Fatalln(err)

}
func main() {

	genDao()

}
