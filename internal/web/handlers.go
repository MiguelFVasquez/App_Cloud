package web

import (
	"html/template"
	"net/http"
	"os"

	"github.com/MiguelFVasquez/App_Cloud/internal/images"
)

// Datos que pasamos al template
type PageData struct {
	Title  string
	Host   string
	Images []images.Imagen
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	host, _ := os.Hostname()

	// Ahora devuelve []Imagen
	imgs, err := images.LoadRandomImages("images", 4)
	if err != nil {
		http.Error(w, "Error cargando imágenes", http.StatusInternalServerError)
		return
	}

	data := PageData{
		Title:  "Servidor de imágenes",
		Host:   host,
		Images: imgs,
	}

	tmpl, err := template.ParseFiles("web/templates/index.html")
	if err != nil {
		http.Error(w, "Error cargando template", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}
