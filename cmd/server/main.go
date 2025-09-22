package main

// Importa los paquetes necesarios
import (
	"fmt"               // Para formato de texto
	"log"               // Para registro de errores
	"net/http"          // Para servidor HTTP
	"os"                // Para acceder a argumentos del sistema
	"github.com/MiguelFVasquez/App_Cloud/internal/web"
)

func main() {
	// Puerto recibido como argumento o valor por defecto
	port := "8000" // Valor por defecto del puerto
	if len(os.Args) > 1 {
		port = os.Args[1] // Si se pasa argumento, se usa como puerto
	}

	// Construye la dirección con el puerto
	addr := fmt.Sprintf(":%s", port)
	// Imprime en consola la dirección donde escucha el servidor
	fmt.Printf("Servidor escuchando en http://localhost%s\n", addr)

	// Registrar el handler para la ruta principal usando el paquete web
	http.HandleFunc("/", web.HomeHandler)

	// Inicia el servidor HTTP y muestra error si ocurre
	log.Fatal(http.ListenAndServe(addr, nil))
}
