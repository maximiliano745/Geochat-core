package models

type Task struct {
	ID          int    `json:"id"`
	Titulo      string `json:"titulo"`
	Descripcion string `json:"descripcion"`
	Estado      string `json:"estado"`
}

type Fase struct {
	ID             int    `json:"id"`
	TaskID         int    `json:"task_id"`
	Nombre         string `json:"nombre"`
	Explicacion    string `json:"explicacion"`
	CodigoSnippet  string `json:"codigo_snippet"`
	ArchivoDestino string `json:"archivo_destino"`
	Completada     bool   `json:"completada"`
}