package images

import (
	"encoding/base64"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// LoadRandomImages lee la carpeta y devuelve N imágenes en base64
func LoadRandomImages(folder string, n int) ([]string, error) {
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		return nil, err
	}

	// Filtrar solo extensiones válidas
	validFiles := []os.FileInfo{}
	for _, f := range files {
		if !f.IsDir() && isImageFile(f.Name()) {
			validFiles = append(validFiles, f)
		}
	}

	if len(validFiles) == 0 {
		return nil, nil // No hay imágenes
	}

	// Barajar (shuffle) y seleccionar máximo N
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(validFiles), func(i, j int) {
		validFiles[i], validFiles[j] = validFiles[j], validFiles[i]
	})

	if n > len(validFiles) {
		n = len(validFiles)
	}

	// Convertir a Base64
	var imagesBase64 []string
	for _, file := range validFiles[:n] {
		data, err := os.ReadFile(filepath.Join(folder, file.Name()))
		if err != nil {
			continue
		}
		encoded := base64.StdEncoding.EncodeToString(data)
		imagesBase64 = append(imagesBase64, encoded)
	}

	return imagesBase64, nil
}

// Verifica si el archivo es .png, .jpg o .jpeg
func isImageFile(name string) bool {
	ext := strings.ToLower(filepath.Ext(name))
	return ext == ".png" || ext == ".jpg" || ext == ".jpeg"
}
