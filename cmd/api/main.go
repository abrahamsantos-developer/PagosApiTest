package main

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"       // Paquete Swagger Files
	"github.com/swaggo/gin-swagger" // Paquete Swagger
	"log"
	_ "myPagosApp/docs" // Importa la documentación generada
	"myPagosApp/internal/handlers"
	"myPagosApp/internal/models"
	"myPagosApp/internal/repositories"
	"myPagosApp/internal/services"
	"myPagosApp/pkg" // Aquí está la conexión a PostgreSQL
)

// @title MyPagosApp API
// @version 1.0
// @description Esta es la API de MyPagosApp para manejar comercios y transacciones.
// @host localhost:3000
// @BasePath /

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

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


	// Definir rutas para comercios
	r.POST("/merchants", merchantHandler.CreateMerchantHandler)
	r.GET("/merchants", merchantHandler.GetAllMerchantsHandler)
	r.PUT("/merchants/:id", merchantHandler.UpdateMerchantHandler)

	// Iniciar el servidor en el puerto 3000
	r.Run(":3000")
}
