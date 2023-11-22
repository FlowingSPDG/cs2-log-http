package main

import (
	"log"

	cs2loghttp "github.com/FlowingSPDG/cs2-log-http"
	"github.com/gin-gonic/gin"
	cs2log "github.com/janstuemmel/cs2-log"
)

func main() {
	r := gin.Default()
	logHandler := cs2loghttp.NewLogHandler(messageHandler)
	r.POST("/csgolog", logHandler.Handle())
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello!"})
	})
	log.Panicf("Failed to listen port 3090 : %v\n", r.Run("0.0.0.0:3090"))
}

func messageHandler(msg cs2log.Message) {
	switch m := msg.(type) {
	case cs2log.PlayerEntered:
		log.Printf("PlayerEntered : %v\n", m)
	case cs2log.PlayerConnected:
		log.Printf("PlayerConnected : %v\n", m)
	case cs2log.WorldMatchStart:
		log.Printf("WorldMatchStart : %v\n", m)
	case cs2log.TeamScored:
		log.Printf("TeamScored : %v\n", m)
	case cs2log.GameOver:
		log.Printf("GameOver : %v\n", m)
	case cs2log.PlayerAttack:
		log.Printf("PlayerAttack : %v\n", m)
	case cs2log.PlayerKill:
		log.Printf("PlayerKill : %v\n", m)
		log.Printf("Meta : %v\n", m.Meta)
	case cs2log.PlayerPurchase:
		log.Printf("PlayerPurchase : %v\n", m)
	case cs2log.PlayerSay:
		log.Printf("PlayerSay : %v\n", m)
	case cs2log.Unknown:
		log.Printf("Unknown : [%v]\n", m.Raw)

	default:
		log.Printf("type[%s] : [%v]\n", msg.GetType(), msg)
	}
}
