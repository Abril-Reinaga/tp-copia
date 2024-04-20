package ejercicio

import "testing"

func TesteoDeEjercicio(t *testing.T) {
	manager := NuevoGestorDeEjercicios()
	manager.AniadirEjercicio("Sentadillas", "Flexión de rodillas y bajar el cuerpo", 30, "Piernas / gluteos", 25, "Baja")
	manager.AniadirEjercicio("Zancada", "Lleva una pierna hacía delante, flexionando las rodillas", 30, "Piernas / Gluteos", 50, "Intermedia")

}
