// Los ejecutables usan siempre package main;
package main

import "fmt"

// Definimos el struct Persona
type Persona struct {
	Nombre string
	Edad   int
}

// Método para saber si es mayor de edad
func (p Persona) EsMayorDeEdad() bool {
	return p.Edad >= 18
}

func main() {
	var personas []Persona

	personas = append(personas,
		Persona{Nombre: "Ana", Edad: 20},
		Persona{Nombre: "Luis", Edad: 15},
	)

	for _, persona := range personas {
		if persona.EsMayorDeEdad() {
			fmt.Println(persona.Nombre, "es mayor de edad.")
		} else {
			fmt.Println(persona.Nombre, "es menor de edad.")
		}
	}
}
