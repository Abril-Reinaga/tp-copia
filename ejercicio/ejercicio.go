package ejercicio

import (
	"errors"
	"fmt"
)

type DescripcionEjerc string
type TiempoEnSegEjerc int
type TipoEjerc string
type PuntosEjerc int
type Dificultad string

type Ejercicio struct {
	Nombre      string
	Descripcion DescripcionEjerc
	Tiempo      TiempoEnSegEjerc
	Tipo        TipoEjerc
	Puntos      PuntosEjerc
	Dificultad  Dificultad
}

type GestorDeEjercicios struct {
	ejercicios []Ejercicio
}

func NuevoGestorDeEjercicios() *GestorDeEjercicios {
	return &GestorDeEjercicios{
		ejercicios: []Ejercicio{},
	}
}

func (manager *GestorDeEjercicios) AniadirEjercicio(nombre string, descripcion DescripcionEjerc, tiempo TiempoEnSegEjerc, tipo TipoEjerc, puntos PuntosEjerc, dificultad Dificultad) {
	manager.ejercicios = append(manager.ejercicios, Ejercicio{
		Nombre:      nombre,
		Descripcion: descripcion,
		Tiempo:      tiempo,
		Tipo:        tipo,
		Puntos:      puntos,
		Dificultad:  dificultad,
	})
}

func (manager *GestorDeEjercicios) QuitarEjercicio(nombre string) error {
	for i, ejercicio := range manager.ejercicios {
		if ejercicio.Nombre == nombre {
			manager.ejercicios = append(manager.ejercicios[:i], manager.ejercicios[i+1:]...)
			return nil
		}
	}

	return errors.New("Ejercicio no encontrado.")
}

func (manager *GestorDeEjercicios) ModificarEjercicio(nombre string, descripcion DescripcionEjerc, tiempo TiempoEnSegEjerc, tipo TipoEjerc, puntos PuntosEjerc, dificultad Dificultad) error {
	for i, ejercicio := range manager.ejercicios {
		if ejercicio.Nombre == nombre {
			manager.ejercicios[i].Descripcion = descripcion
			manager.ejercicios[i].Tiempo = tiempo
			manager.ejercicios[i].Tipo = tipo
			manager.ejercicios[i].Puntos = puntos
			manager.ejercicios[i].Dificultad = dificultad
			return nil
		}
	}
	return errors.New("Ejercicio no encontrado.")
}

func (manager *GestorDeEjercicios) ObtenerEjercicio(nombre string) (*Ejercicio, error) {
	for _, ejercicio := range manager.ejercicios {
		if ejercicio.Nombre == nombre {
			return &ejercicio, nil
		}
	}
	return nil, errors.New("Ejercicio no encontrado")
}

func (manager *GestorDeEjercicios) ListaDeEjercicios() string {
	lista := "Lista de ejercicios: \n"
	for _, ejercicio := range manager.ejercicios {
		lista += fmt.Sprintf("%s : \n (Descripción: %s \n Tiempo: %d \n Tipo: %s \n Puntos: %d \n Dificultad: %s) \n", ejercicio.Nombre, ejercicio.Descripcion, ejercicio.Tiempo, ejercicio.Tipo, ejercicio.Puntos, ejercicio.Dificultad)
	}
	return lista
}
