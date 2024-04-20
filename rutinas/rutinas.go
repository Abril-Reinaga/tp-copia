package rutinas

type rutina struct {
	Nombre        string `csv:"nombre"`
	Duracion      int    `csv:"duracion"`
	TipoEjercicio string `csv:"tipoEjercicio"`
	Dificultad    string `csv:"dificultad"`
}
