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
		// Ajustado a tu local. Recuerda que 'geochat' debe existir en tu Postgres
		connStr = "postgresql://postgres:pass@localhost:5432/geochat?sslmode=disable"
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("❌ Error: No se pudo conectar a Postgres:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("❌ Postgres no responde:", err)
	}

	log.Println("🐘 IA 5: Memoria Relacional (Postgres) conectada con éxito.")
	DB = db

	// Crear el esquema soberano
	crearTablasSoberanas()
}

func crearTablasSoberanas() {
	query := `
    -- Tabla de Propuestas (El cerebro financiero y de infraestructura)
    CREATE TABLE IF NOT EXISTS propuestas (
        id TEXT PRIMARY KEY,
        modulo TEXT NOT NULL,
        descripcion TEXT,
        monto NUMERIC(15,2),
        impacto TEXT,
        estado TEXT DEFAULT 'ESPERANDO_FIRMA',
        fecha TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        requiere_firma BOOLEAN DEFAULT true,
        firma_digital TEXT
    );

    -- Tabla de Estadísticas (La salud del organismo vivo)
    CREATE TABLE IF NOT EXISTS estadisticas (
        id SERIAL PRIMARY KEY,
        plasticidad DOUBLE PRECISION,
        energia_vendida NUMERIC(15,2),
        buena_onda_count INTEGER,
        usuarios_activos INTEGER,
        fecha_actualizacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

    -- Tabla de Tareas Operativas (Tu backlog de desarrollo original)
    CREATE TABLE IF NOT EXISTS tasks (
        id SERIAL PRIMARY KEY,
        titulo TEXT NOT NULL,
        descripcion TEXT,
        estado TEXT DEFAULT 'propuesta_ceo'
    );`

	_, err := DB.Exec(query)
	if err != nil {
		log.Printf("⚠️ IA 5 Error al crear esquema: %v", err)
	} else {
		log.Println("🧬 IA 5: Esquema de GeoChat (Capa 4) verificado y listo.")
	}
}

// GuardarPropuesta inyecta la intención de la IA en la memoria física
func GuardarPropuesta(id, modulo, desc string, monto float64, impacto, estado string, firma bool) error {
    query := `INSERT INTO propuestas (id, modulo, descripcion, monto, impacto, estado, requiere_firma) 
              VALUES ($1, $2, $3, $4, $5, $6, $7)
              ON CONFLICT (id) DO UPDATE SET estado = EXCLUDED.estado`
    
    _, err := DB.Exec(query, id, modulo, desc, monto, impacto, estado, firma)
    return err
}