
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"biblioteca/internal/application/autor"
	libroapp "biblioteca/internal/application/libro"
	"biblioteca/internal/infrastructure/database"
	"biblioteca/internal/infrastructure/persistence"
	"biblioteca/internal/interfaces/http/facade"
	"biblioteca/internal/interfaces/http/routes"
)

func main() {
	// Cargar las variables del archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Advertencia: no se pudo cargar el archivo .env")
	}

	// Conectar con Supabase PostgreSQL
	db, err := database.Connect(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Error al conectar con la base de datos: ", err)
	}
	defer db.Close()

	// Inicializar repositorios
	autorRepo := persistence.NewAutorRepository(db)
	libroRepo := persistence.NewLibroRepository(db)

	// Inicializar servicios de aplicación
	autorService := autor.NewService(autorRepo)
	libroService := libroapp.NewService(libroRepo, autorRepo)

	// Inicializar Facade
	bibliotecaFacade := facade.New(autorService, libroService)

	// Configurar puerto
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Configurar rutas HTTP
	handler := routes.Setup(bibliotecaFacade)

	// Iniciar servidor
	log.Printf("API de biblioteca escuchando en http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
