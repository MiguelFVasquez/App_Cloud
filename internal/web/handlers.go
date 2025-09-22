package web

import (
    "html/template"
    "net/http"
    "os"

    "github.com/MiguelFVasquez/App_Cloud/internal/images"
)

// Datos que pasamos al template
type PageData struct {
    Title  string   // Título de la página
    Host   string   // Nombre del host
    Images []string // Lista de imágenes en base64
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    // Obtiene el nombre del host del sistema
    host, _ := os.Hostname()

    // Cargar 4 imágenes aleatorias
    imgs, err := images.LoadRandomImages("images", 4)
    if err != nil {
        http.Error(w, "Error cargando imágenes", http.StatusInternalServerError)
        return
    }

    // Crea los datos que se pasarán al template
    data := PageData{
        Title:  "Servidor de imágenes",
        Host:   host,
        Images: imgs,
    }

    // Carga el archivo de plantilla HTML
    tmpl, err := template.ParseFiles("web/templates/index.html")
    if err != nil {
        http.Error(w, "Error cargando template", http.StatusInternalServerError)
        return
    }

    // Ejecuta la plantilla y envía la respuesta al cliente
    tmpl.Execute(w, data)
}
