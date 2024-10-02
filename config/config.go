package config

import "os"

// obtiene las variables de entorno con valor default
func GetEnv(key, fallback string) string {
    value := os.Getenv(key)
    if len(value) == 0 {
        return fallback
    }
    return value
}
