package main

import "fmt"

func main() {
	/* err := CrearMateria(Materia{
		Descrip:         "Inglés",
		CantidadAlumnos: 15,
	})
	if err != nil {
		fmt.Println(err)
	}

	ms, err := LeerMaterias()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ms)
	} */

	/* err = BorrarMateria(3)
	if err != nil {
		fmt.Println(err)
	} */

	err := ActualizarMateria(Materia{
		IdMateria: 4,
		Descrip:   "Prueba",
	})
	if err != nil {
		fmt.Println(err)
	}
}
