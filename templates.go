package main

import (
	"os"
	"text/template"
	"time"
)

func main() {
	t1, err := template.New("t1").Parse("Name: {{.}}\n")
	if err != nil {
		panic(err)
	}

	t1.Execute(os.Stdout, "Dan")

	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	t2 := Create("t2", "{{.Name}}, {{.DateOfBirth}}")
	data := struct {
		Name        string
		DateOfBirth time.Time
	}{"Dan", time.Date(1980, time.Month(4), 14, 0, 0, 0, 0)}
	t2.Execute(os.Stdout, data)
}
