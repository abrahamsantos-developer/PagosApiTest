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

	// monfigurar modo release
	gin.SetMode(gin.ReleaseMode)

	// Conecta a la base de datos
	db := pkg.ConnectDB()
	if db == nil {
		log.Fatal("No se pudo conectar a la base de datos")
	}

	// migra modelo Merchant para crear la tabla si no existe
	db.AutoMigrate(&models.Merchant{}, &models.Transaction{})

	// Inicializa repo service y handler de merchants
	merchantRepo := repositories.NewMerchantRepository(db)
	merchantService := services.NewMerchantService(merchantRepo)
	merchantHandler := handlers.NewMerchantHandler(merchantService)

	// Inicializa repo service y handler de transactions
	transactionRepo := repositories.NewTransactionRepository(db)
	transactionService := services.NewTransactionService(transactionRepo, merchantRepo)
	transactionHandler := handlers.NewTransactionHandler(transactionService)

	// Inicializa service y handler de profits
	profitService := services.NewProfitService(transactionRepo)
	profitHandler := handlers.NewProfitHandler(profitService)

	// Inicializar el router de gin
	r := gin.Default()

	// sirve para verificar que el servidor funciona
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "¡Servidor corriendo correctamente, OK OK OK!",
		})
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// routes de merchants
	r.POST("/merchants", merchantHandler.CreateMerchantHandler)
	r.GET("/merchants", merchantHandler.GetAllMerchantsHandler)
	r.GET("/merchants/:id", merchantHandler.GetMerchantByIDHandler)
	r.PUT("/merchants/:id", merchantHandler.UpdateMerchantHandler)

	//routes de transactions
	r.POST("/transactions", transactionHandler.CreateTransactionHandler)
	r.GET("/transactions", transactionHandler.GetAllTransactionsHandler)
	r.GET("/transactions/:id", transactionHandler.GetTransactionByIDHandler)
	r.GET("/transactions/merchant/:merchant_id", transactionHandler.GetTransactionsByMerchantIDHandler)

	// routes de ganancias
	r.GET("/profits", profitHandler.GetTotalProfitsHandler)
	r.GET("/profits/merchant/:merchant_id", profitHandler.GetProfitsByMerchantIDHandler)

	// server port: 3000
	r.Run(":3000")
}
