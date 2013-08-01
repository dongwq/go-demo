package main

import (
	"text/template"
	"os"
	"log"
)

func main() {

	type Inventory struct {
		Material string
		Count    uint
	}

	sweaters := Inventory{"wool", 17}
	tmpl, err := template.New("test").Parse("1.{{.Count}} items are made of {{.Material}}\n")
	if err != nil { panic(err) }
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil { panic(err) }


	t, err := template.New("foo").Parse(`2. {{define "T"}}Hello, {{.}}!{{end}}\n`)
	if err != nil{
		log.Fatalf("parse err",err)
		return
	}
	out := os.Stdout
	err = t.ExecuteTemplate(out, "T", "3. <script>alert('you have been pwned')</script>\n")
	if err != nil{
		log.Fatalf("parse err",err)
		return
	}


}
