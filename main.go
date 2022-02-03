package main

import "fmt"


func main()  {
	paises := make(map[string]string)
	fmt.Println(paises)

	paises["Mexico"]="D.F."
	paises["Argentina"]="Buenos Aires"

	fmt.Println(paises["Mexico"])
	fmt.Println(paises)

	//Ordena alfabeticamente por clave
	campeonato := map[string]int{
		"Barcelona": 39,
		"Real Madrid": 38,
		"Chivas": 37,
		"Boca Juniors": 30
	}

	fmt.Println(campeonato)

	//Agregamos un elemento al mapa campeonato
	campeonato["River Plate"] = 25

	//Modificamos el valor a la clave "Chivas" del mapa campeonato
	campeonato["Chivas"] = 55

	//Eliminamos el elemento con clave "Real Madrid" del mapa campeonato
	delete(campeonato, "Real Madrid")

	//range no ordena alfabeticamente
	for equipo, puntos := range campeonato{
		fmt.Printf("El equipo %s tiene una puntuación de %d \n", equipo, puntos)
	}

	//comprobar si un elemento con la clave "Mineiro" existe en el mapa
	puntos, existe := campeonato["Mineiro"]
	fmt.Printf("La puntuación es %d y el equipo existe %t \n", puntos, existe)

}	
