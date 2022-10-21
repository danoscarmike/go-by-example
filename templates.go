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

	t2 := Create("t2", "{{.Name}}: {{.DateOfBirth.Month}} {{.DateOfBirth.Day}}, {{.DateOfBirth.Year}}\n")
	data := struct {
		Name        string
		DateOfBirth time.Time
	}{"Dan", time.Date(1912, time.April, 14, 0, 0, 0, 0, time.UTC)}
	t2.Execute(os.Stdout, data)

	type animal struct {
		Real bool
		Name string
	}
	t3 := Create("t3", "{{if .Real}}{{.Name}} are real, seriously! {{else}}{{.Name}} are a myth, fool!{{end}}\n")
	quoka := animal{true, "quokas"}
	unicorn := animal{false, "unicorns"}
	t3.Execute(os.Stdout, quoka)
	t3.Execute(os.Stdout, unicorn)

	languages := []string{"C++", "Go", "Java", "Python"}
	t4 := Create("t4", "{{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout, languages)

}
