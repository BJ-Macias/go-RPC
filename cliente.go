package main

import (
	"fmt"
	"net/rpc"
)

type Calificacion struct {
	Materia string
	Alumno  string
	Cali   float32
}

func main() {
	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	var op int64
	for {
		fmt.Println("1.Agregar calificacion")
		fmt.Println("2.Promedio de alumno")
		fmt.Println("3.Promedio general")
		fmt.Println("4.Promedio de materia")
		fmt.Println("5.Exit")
		fmt.Scanln(&op)

		switch op {
		case 1:
			var cal Calificacion
			fmt.Print("Ingrese materia: ")
			fmt.Scanln(&cal.Materia)

			fmt.Print("Ingrese alumno: ")
			fmt.Scanln(&cal.Alumno)

			fmt.Print("Ingrese calificacion: ")
			fmt.Scanln(&cal.Cali)

			var result string
			err = c.Call("Server.CalificaAlumno", cal, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(result)
			}
			break
		case 2:
			var nombre string
			fmt.Print("Ingrese alumno a promediar: ")
			fmt.Scanln(&nombre)

			var result float32
			err = c.Call("Server.PromedioAlumno", nombre, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Promedio de", nombre+":", result)
			}
			break
		case 3:
			var result float32
			err = c.Call("Server.PromedioGeneral", 0, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Promedio general:", result)
			}
			break
		case 4:
			var materia string
			fmt.Print("Ingrese materia a promediar: ")
			fmt.Scanln(&materia)

			var result float32
			err = c.Call("Server.PromedioMateria", materia, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Promedio de", materia+":", result)
			}
			break
		case 5:
			return
		}
		fmt.Println()
	}
}