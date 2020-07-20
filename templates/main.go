package main

import (
	"log"
	"os"
	"text/template"	
)

func main() {
	// temp, err := template.ParseFiles("basic_template.gohtml", "template_with_collection_variable.gohtml")
	temp, err := template.ParseGlob("./*.gohtml")

	if err != nil {
		log.Fatalln(err)
	}

	// err = temp.Execute(os.Stdout, "Bruce")
	err = temp.ExecuteTemplate(os.Stdout, "basic_template.gohtml", "Bruce")

	if err != nil {
		log.Fatalln(err)
	}
}