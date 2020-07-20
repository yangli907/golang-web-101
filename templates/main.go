package main

import (
	"log"
	"os"
	"text/template"	
	"strings"
)

var temp *template.Template

var funcMap = template.FuncMap {
	"toLower": strings.ToLower,
	"timesTwo": timesTwo,
}

func timesTwo(val int) int {
	return val * 2
}

func (p Player) IsPF() bool {
	return p.Position == "PF"
}

func init() {
	temp = template.Must(template.New("").Funcs(funcMap).ParseGlob("./*.gohtml"))
}

// Note: Must use capitalized field name, otherwise can't access from template
type Player struct {
	Fname string
	Lname string
	Position string
}

type Car struct {
	Make string
	Model string
	Year int
}

func main() {
	// temp, err := template.ParseFiles("basic_template.gohtml", "template_with_collection_variable.gohtml")
	// temp := template.Must(template.ParseGlob("./*.gohtml"))

	// err = temp.Execute(os.Stdout, "Bruce")
	err := temp.ExecuteTemplate(os.Stdout, "basic_template.gohtml", "Bruce")

	if err != nil {
		log.Fatalln(err)
	}

	// slice definition example
	// names := []string{"Harden", "Durant", "James", "Jordan", "Paul"}

	// map definition example
	// names := map[string]string{
	// 	"James": "Harden", 
	// 	"Kevin": "Durant", 
	// 	"LeBron": "James", 
	// 	"Michael": "Jordan", 
	// 	"Chris": "Paul",
	// }

	// struct definition example
	players := []Player {
		Player{"LeBron", "James", "PF"},
		Player{"Kevin", "Durant", "SF"},
	}

	cars := []Car {
		Car{"Honda", "Accord", 2013},
		Car{"Acura", "RDX", 2020},
	}

	data := struct {
		Person []Player
		Transport []Car
	} {
		players,
		cars,
	}

	// err = temp.ExecuteTemplate(os.Stdout, "template_with_collection_variable.gohtml", players)
	err = temp.ExecuteTemplate(os.Stdout, "template_with_collection_variable.gohtml", data)
}