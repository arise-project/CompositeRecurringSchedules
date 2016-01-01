package main

import (
    "text/template"
    "os")

type Inventory struct {
	Material string
	Count    uint
}

func main() {
    
sweaters := Inventory{"wool", 17}
tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}\n")
if err != nil { panic(err) }
err = tmpl.Execute(os.Stdout, sweaters)
if err != nil { panic(err) }
//TODO reverse with https://golang.org/pkg/text/template/parse/
    }

