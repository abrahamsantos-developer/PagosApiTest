package config

import "os"

// GetEnv obtiene las variables de entorno con un valor por defecto
func GetEnv(key, fallback string) string {
    value := os.Getenv(key)
    if len(value) == 0 {
        return fallback
    }
    return value
}
