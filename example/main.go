package main

import (
	"github.com/FlowingSPDG/csgo-log"
	"github.com/FlowingSPDG/csgo-log-http"
	"github.com/gin-gonic/gin"
	"log"
)

// File log format
// L 11/12/2018 - 19:57:28: World triggered "Round_Start"
// csgolog.LogLinePattern = regexp.MustCompile(`L (\d{2}\/\d{2}\/\d{4} - \d{2}:\d{2}:\d{2}): (.*)`)

// HTTP log format
// 02/09/2020 - 22:42:32.000 - World triggered "Round_Start"
// csgolog.LogLinePattern = regexp.MustCompile(`(\d{2}\/\d{2}\/\d{4} - \d{2}:\d{2}:\d{2}.\d{3}) - (.*)`)

func main() {
	r := gin.Default()
	r.POST("/csgolog", csgologhttp.CSGOLogger(MessageHandler))
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello!"})
	})
	log.Panicf("Failed to listen port 3090 : %v\n", r.Run("0.0.0.0:3090"))
}

// MessageHandler handles message from CSGO Server and Gin middleware
func MessageHandler(msg csgolog.Message) {
	switch m := msg.(type) {
	case csgolog.PlayerEntered:
		log.Printf("PlayerEntered : %v\n", m)
	case csgolog.PlayerConnected:
		log.Printf("PlayerConnected : %v\n", m)
	case csgolog.WorldMatchStart:
		log.Printf("WorldMatchStart : %v\n", m)
	case csgolog.TeamScored:
		log.Printf("TeamScored : %v\n", m)
	case csgolog.GameOver:
		log.Printf("GameOver : %v\n", m)
	case csgolog.PlayerAttack:
		log.Printf("PlayerAttack : %v\n", m)
	case csgolog.PlayerKill:
		log.Printf("PlayerKill : %v\n", m)
		log.Printf("Meta : %v\n", m.Meta)
	case csgolog.PlayerPurchase:
		log.Printf("PlayerPurchase : %v\n", m)
	case csgolog.PlayerSay:
		log.Printf("PlayerSay : %v\n", m)
	case csgolog.ServerCvar:
		log.Printf("ServerCvar : %v\n", m)
	case csgolog.Get5Event:
		log.Printf("Get5Event : [%v]\n", m)
	case csgolog.PlayerKillOther:
		log.Printf("PlayerKillOther : [%v]\n", m)
	case csgolog.Unknown:
		log.Printf("Unknown : [%v]\n", m.Raw)

	default:
		log.Printf("type[%s] : [%v]\n", msg.GetType(), msg)
	}
}
