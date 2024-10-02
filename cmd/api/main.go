package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"myPagosApp/internal/handlers"
	"myPagosApp/internal/repositories"
	"myPagosApp/internal/services"
	"myPagosApp/internal/models"
	"myPagosApp/pkg" // Aquí está la conexión a PostgreSQL
)

func main() {

	// Configurar modo release en producción
	gin.SetMode(gin.ReleaseMode)

	// Conectarse a la base de datos
	db := pkg.ConnectDB()
	if db == nil {
		log.Fatal("No se pudo conectar a la base de datos")
	}

	// Migrar el modelo Merchant para crear la tabla si no existe
	db.AutoMigrate(&models.Merchant{})

	// Inicializar repositorio, servicio y handler para comercios
	merchantRepo := repositories.NewMerchantRepository(db)
	merchantService := services.NewMerchantService(merchantRepo)
	merchantHandler := handlers.NewMerchantHandler(merchantService)

	// Inicializar el router de Gin
	r := gin.Default()

	// Ruta simple para verificar que el servidor funciona
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "¡Servidor corriendo correctamente, OK OK OK!",
		})
	})

	// Definir rutas para comercios
	r.POST("/merchants", merchantHandler.CreateMerchantHandler)
	r.GET("/merchants", merchantHandler.GetAllMerchantsHandler)
	r.PUT("/merchants/:id", merchantHandler.UpdateMerchantHandler)

	// Iniciar el servidor en el puerto 3000
	r.Run(":3000")
}
