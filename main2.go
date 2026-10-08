package main2

import "fmt"

func actividadGanadora(votos map[string]int) string {
	ganadora := ""
	maximo := -1

	for actividad, cantidad := range votos {
		if cantidad > maximo {
			maximo = cantidad
			ganadora = actividad
		}
	}

	return ganadora
}

func main() {
	votos := map[string]int{
		"deportes":   0,
		"videojuegos": 0,
		"cine":       0,
		"musica":     0,
	}

	fmt.Println("=== VOTACION DE ACTIVIDADES ===")
	fmt.Println("Opciones: deportes, videojuegos, cine, musica")

	for i := 1; i <= 5; i++ {
		var voto string
		fmt.Printf("Voto %d: ", i)
		fmt.Scan(&voto)

		if _, existe := votos[voto]; existe {
			votos[voto]++
		} else {
			fmt.Println("Actividad no valida, voto no contado.")
		}
	}

	fmt.Println()
	fmt.Println("=== RESULTADOS ===")
	for actividad, cantidad := range votos {
		fmt.Printf("%s: %d votos\n", actividad, cantidad)
	}

	fmt.Println()
	fmt.Println("Actividad con mayor aceptacion:", actividadGanadora(votos))
}