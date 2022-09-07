package templates

import (
	"os"
	"text/template"
)

func CreateTemplate(name, t string) *template.Template {
	return template.Must(template.New(name).Parse(t))
}

func Run() {

	// simple template
	t1 := template.New("template-1")

	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		panic(err)
	}

	t1.Execute(os.Stdout, "some val")

	// destructing from other data structures
	t2 := CreateTemplate("t2", "First Name: {{.FName}}, Last Name: {{.LName}}\n")
	type personType struct {
		FName string
		LName string
	}
	t2.Execute(os.Stdout, personType{"From", "Explicit Struct"})
	t2.Execute(os.Stdout, struct {
		FName string
		LName string
	}{"From", "Anonymous Struct"})
	t2.Execute(
		os.Stdout,
		map[string]string{
			"FName": "From",
			"LName": "map",
		},
	)

	// conditionals
	//  - trims whitespace
	conditionalTemplate := CreateTemplate(
		"conditional",
		"{{if . -}}  yes {{else -}} no {{end}}\n",
	)

	conditionalTemplate.Execute(
		os.Stdout,
		2+2 == 4,
	)

	// loops
	loopTemplate := CreateTemplate(
		"loop",
		"Looping {{range .}}\n{{.}}{{end}}\n",
	)
	loopTemplate.Execute(
		os.Stdout,
		[]string{
			"a", "b", "c",
		},
	)

}
