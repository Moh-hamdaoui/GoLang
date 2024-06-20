package exercices

import ("fmt")

func Maps(){

	capitales := map[string]string{
        "France": "Paris",
        "Maroc":  "Rabat",
        "Espagne": "Madrid",
    }

    pays := "Maroc"
    capitale, existe := capitales[pays]
    if existe {
        fmt.Printf("La capitale de %s est %s.\n", pays, capitale)
    } else {
        fmt.Printf("La capitale de %s n'est pas trouvée dans la map.\n", pays)
    }
}