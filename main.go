package main
 
import "fmt"
 
func main() {
	// Array bidimensional: 6 estudiantes x 4 materias
	notas := [6][4]float64{
		{8.5, 9.0, 7.5, 8.0},
		{6.0, 7.5, 8.0, 6.5},
		{9.5, 9.0, 10.0, 9.5},
		{5.5, 6.0, 7.0, 6.5},
		{7.0, 8.5, 7.5, 9.0},
		{8.0, 7.0, 6.5, 7.5},
	}
 
	sumaClase := 0.0
 
	fmt.Println("=== ANALISIS DE NOTAS ===")
	fmt.Println()
 
	for i := 0; i < len(notas); i++ {
		// Slice a partir del array: notas de un estudiante
		fila := notas[i][:]
 
		suma := 0.0
		maxima := fila[0]
		minima := fila[0]
 
		for _, n := range fila {
			suma += n
			if n > maxima {
				maxima = n
			}
			if n < minima {
				minima = n
			}
		}
 
		promedio := suma / float64(len(fila))
		sumaClase += promedio
 
		fmt.Printf("Estudiante %d\n", i+1)
		fmt.Println("  Notas:", fila)
		fmt.Printf("  Promedio:    %.2f\n", promedio)
		fmt.Printf("  Nota mas alta: %.2f\n", maxima)
		fmt.Printf("  Nota mas baja: %.2f\n", minima)
		fmt.Println()
	}
 
	promedioGeneral := sumaClase / float64(len(notas))
	fmt.Printf("Promedio general de la clase: %.2f\n", promedioGeneral)
}

