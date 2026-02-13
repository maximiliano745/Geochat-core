package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq" 
)

var DB *sql.DB

func ConectarDB() {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		// Ajusta estos datos a tu DB local si no estás en Render
		connStr = "postgresql://postgres:pass@localhost:5432/geochat?sslmode=disable"
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("❌ Error: No se pudo conectar a Postgres:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("❌ Postgres no responde:", err)
	}

	log.Println("🐘 IA 5: Memoria Relacional (Postgres) conectada.")
	DB = db

	// Crear tablas iniciales si no existen
	crearTablasIniciales()
}

func crearTablasIniciales() {
	query := `
	CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		titulo TEXT NOT NULL,
		descripcion TEXT,
		estado TEXT DEFAULT 'propuesta_ceo'
	);

	CREATE TABLE IF NOT EXISTS fases (
		id SERIAL PRIMARY KEY,
		task_id INTEGER REFERENCES tasks(id) ON DELETE CASCADE,
		nombre TEXT,
		explicacion TEXT,
		codigo_snippet TEXT,
		archivo_destino TEXT,
		completada BOOLEAN DEFAULT false
	);`

	_, err := DB.Exec(query)
	if err != nil {
		log.Printf("⚠️ IA 5 Error al crear tablas: %v", err)
	} else {
		log.Println("🧬 IA 5: Esquema de base de datos verificado.")
	}
}