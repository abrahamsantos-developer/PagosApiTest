package main

import (
    "log"
    "myPagosApp/pkg"     // Aquí está la conexión a PostgreSQL
    "github.com/gin-gonic/gin"
)

func main() {
    // Conectarse a la base de datos
    db := pkg.ConnectDB()
    if db == nil {
        log.Fatal("No se pudo conectar a la base de datos")
    }

    // Inicializar el router de Gin
    r := gin.Default()

    // Ruta simple para verificar que el servidor funciona
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "¡Servidor corriendo correctamente!",
        })
    })

    // Iniciar el servidor en el puerto 3000
    r.Run(":3000")
}
