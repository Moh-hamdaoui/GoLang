package exercices

import ("fmt")
	
	
func Tab(){

	array := [5] int{1, 2, 3, 4, 5}

	slice := make([]int, len(array))


	for i := 0; i < len(array); i++ {
		slice[i] = array[i]
	}

	fmt.Println("Votre slice:", slice)
}