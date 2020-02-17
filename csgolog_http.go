package csgologhttp

import (
	"bufio"
	"github.com/FlowingSPDG/csgo-log"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func CSGOLogger(Handler func(csgolog.Message, *gin.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		raw, err := c.GetRawData()
		if err != nil {
			log.Printf("Failed to get raw data : %v\n", err)
			c.String(http.StatusInternalServerError, err.Error())
			c.Abort()
			return
		}

		csgolog.LogLinePattern = csgolog.HTTPLinePattern

		sl := strings.NewReader(string(raw))
		scanner := bufio.NewScanner(sl)
		for scanner.Scan() {
			// fmt.Println(scanner.Text())
			msg, err := csgolog.Parse(scanner.Text())
			if err != nil {
				log.Printf("Failed to parse data : %v\n", err)
				c.String(http.StatusInternalServerError, err.Error())
				c.Abort()
				return
			}
			Handler(msg, c)
		}
		c.String(http.StatusOK, "OK")
		c.Abort()
	}
}
