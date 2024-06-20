package main

import (
	"fmt"
	"mohammed/goland/exercices"
	"mohammed/goland/person"
)

func main(){
	
	// Exercice 1
	exercices.Ex1()
	// Exercice 2
	fmt.Println(exercices.Add(10, 5))
	exercices.Greet("Mohammed")
	// Exercice 3
	exercices.CheckEvenOdd(5)
	// Exercice 4
	exercices.PrintNumbers()
	// Exercice 5
	exercices.Tab()
	
	// Exercice 6
	p := person.Personnage{
        Nom: "Mohammed",
        Age: 30,
    }

    fmt.Println("Nom:", p.Nom)
    fmt.Println("Age:", p.Age)

	// Exercice 7
	exercices.Maps()
}