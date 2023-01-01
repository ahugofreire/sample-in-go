package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	curso := Curso{"Go", 40}
	tmp := template.New("CursoTemplate")
	tmp, _ = tmp.Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}}")
	err := tmp.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}

	//Usando Must
	t := template.Must(template.New("CursoTemplate").Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}}"))
	err = t.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}

	//template externo
	t1 := template.Must(template.New("template.html").ParseFiles("template.html"))
	err = t1.Execute(os.Stdout, Cursos{
		{"Go", 80},
		{"Java", 85},
		{"Node", 65},
	})

	if err != nil {
		panic(err)
	}
}
