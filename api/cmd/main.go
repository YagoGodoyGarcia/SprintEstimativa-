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
		AllowAllOrigins:  true, // <-- PERMITE QUALQUER ORIGEM
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodOptions,
		},
		AllowedHeaders:   []string{"*"}, // Permite quaisquer cabeçalhos
		AllowCredentials: true, // Permite credenciais
		// Debug: true, // Descomente se ainda tiver problemas para ver logs detalhados do CORS
	} )
		
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
