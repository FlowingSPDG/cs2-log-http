package main

import (
	"log"

	cs2loghttp "github.com/FlowingSPDG/cs2-log-http"
	"github.com/gin-gonic/gin"
	cs2log "github.com/janstuemmel/cs2-log"
)

// File log format
// L 11/12/2018 - 19:57:28: World triggered "Round_Start"
// csgolog.LogLinePattern = regexp.MustCompile(`L (\d{2}\/\d{2}\/\d{4} - \d{2}:\d{2}:\d{2}): (.*)`)

// HTTP log format
// 02/09/2020 - 22:42:32.000 - World triggered "Round_Start"
// csgolog.LogLinePattern = regexp.MustCompile(`(\d{2}\/\d{2}\/\d{4} - \d{2}:\d{2}:\d{2}.\d{3}) - (.*)`)

func main() {
	r := gin.Default()
	r.POST("/csgolog", cs2loghttp.CS2Logger(MessageHandler))
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello!"})
	})
	log.Panicf("Failed to listen port 3090 : %v\n", r.Run("0.0.0.0:3090"))
}

// MessageHandler handles message from CS2 Server and Gin middleware
func MessageHandler(msg cs2log.Message, c *gin.Context) {
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
