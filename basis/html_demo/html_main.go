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
	tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
	if err != nil { panic(err) }
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil { panic(err) }


	t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	if err != nil{
		log.Fatalf("parse err",err)
		return
	}
	out := os.Stdout
	err = t.ExecuteTemplate(out, "T", "<script>alert('you have been pwned')</script>")
	if err != nil{
		log.Fatalf("parse err",err)
		return
	}


}
