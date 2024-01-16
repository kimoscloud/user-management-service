package utils

import (
	"fmt"
	"regexp"
	"strings"
)

func CreateSlug(input string) string {
	// Convertir el texto en minúsculas
	slug := strings.ToLower(input)

	// Reemplazar espacios con guiones
	slug = strings.ReplaceAll(slug, " ", "-")

	// Eliminar caracteres especiales excepto guiones y letras
	reg, err := regexp.Compile("[^a-zA-Z0-9-]+")
	if err != nil {
		fmt.Println(err)
	}
	slug = reg.ReplaceAllString(slug, "")

	// Recortar la longitud si es necesario (por ejemplo, a 50 caracteres)
	if len(slug) > 50 {
		slug = slug[:50]
	}

	return slug
}
