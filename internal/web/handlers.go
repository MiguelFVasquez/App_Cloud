package web

// Importa los paquetes necesarios
import (
	"html/template" // Para manejar plantillas HTML
	"net/http"      // Para manejar peticiones HTTP
	"os"            // Para obtener información del sistema
)

// Datos que pasamos al template
type PageData struct {
	Title  string   // Título de la página
	Host   string   // Nombre del host
	Images []string // Imágenes en base64
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Obtiene el nombre del host del sistema
	host, _ := os.Hostname()

	// Crea los datos que se pasarán al template
	data := PageData{
		Title:  "Servidor de imágenes", // Título mostrado en la página
		Host:   host,                   // Host actual
		Images: []string{},             // Lista de imágenes (vacía por ahora)
	}

	// Carga el archivo de plantilla HTML
	tmpl, err := template.ParseFiles("web/templates/index.html")
	if err != nil {
		// Si hay error al cargar la plantilla, responde con error 500
		http.Error(w, "Error cargando template", http.StatusInternalServerError)
		return
	}

	// Ejecuta la plantilla y envía la respuesta al cliente
	tmpl.Execute(w, data)
}
