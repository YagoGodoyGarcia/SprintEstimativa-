package main

import (
	"github.com/YagoGodoyGarcia/SprintEstimativa-/backend/internal/auth"
	"github.com/YagoGodoyGarcia/SprintEstimativa-/backend/internal/room"
	"github.com/YagoGodoyGarcia/SprintEstimativa-/backend/internal/ws"
	"github.com/YagoGodoyGarcia/SprintEstimativa-/backend/pkg/config"
	db "github.com/YagoGodoyGarcia/SprintEstimativa-/backend/pkg/dbcontext"
	"github.com/YagoGodoyGarcia/SprintEstimativa-/backend/pkg/log"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"net/http"
	"os"
)

var (
	logger         = log.New()
	database       = db.SetupDatabaseConnection(config.GetDbClient())
	roomRepository = room.NewRoomRepository(database, logger)
	roomService    = room.NewRoomService(roomRepository, logger)
	roomController = room.NewRoomController(roomService, logger)
	wsServer       = ws.NewWSServer(roomService)
)

func main() {
	defer database.Close()

	r := gin.Default()

	c := cors.New(cors.Options{
		// Permite a origem configurada via ENV e também 127.0.0.1
		AllowedOrigins: []string{os.Getenv("CLIENT_URL"), "http://127.0.0.1:3001", "https://8082-cs-2a161208-33df-4cff-a8f6-60bc958ae240.cs-us-east1-dogs.cloudshell.dev/", "https://8081-cs-2a161208-33df-4cff-a8f6-60bc958ae240.cs-us-east1-dogs.cloudshell.dev/","34.148.199.162","https://3001-cs-2a161208-33df-4cff-a8f6-60bc958ae240.cs-us-east1-dogs.cloudshell.dev/"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodOptions, // Permite o método OPTIONS para preflight requests
			// Adicione outros métodos se sua API usar (ex: http.MethodPut, http.MethodDelete )
		},
		AllowedHeaders:   []string{"*"}, // Permite quaisquer cabeçalhos (ajuste se necessário para produção)
		AllowCredentials: true, // Permite envio de cookies/auth headers
		// Debug: true, // Descomente esta linha se precisar depurar mais o CORS
	})
	r.Use(c)
	

	roomRoutes := r.Group("api/room")
	{
		roomRoutes.POST("/", roomController.CreateRoom)
		roomRoutes.POST("/:id", roomController.JoinRoom)
		roomRoutes.GET("/:id", auth.IsUserAuthorizedInRoom, roomController.GetRoom)

		roomRoutes.GET("/ws/:token", func(ctx *gin.Context) {
			token := ctx.Param("token")
			userClaims, err := auth.GetUserClaimsFromToken(token)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, "Invalid token")
				return
			}

			ws.ServeWS(wsServer, ctx.Writer, ctx.Request, userClaims)
		})
	}

	r.Run()
}
