package exercices

import ("fmt")

func Add(a int, b int)(int){
	var somme int = a + b
	return somme
}

func Greet(nom string){
	fmt.Println("Hello",nom)
}

