package cs2loghttp

import (
	"bufio"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	cs2log "github.com/janstuemmel/cs2-log"
)

func CS2Logger(Handler func(cs2log.Message, *gin.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		raw, err := c.GetRawData()
		if err != nil {
			log.Printf("Failed to get raw data : %v\n", err)
			c.String(http.StatusInternalServerError, err.Error())
			c.Abort()
			return
		}

		// cs2log.LogLinePattern = cs2log.HTTPLinePattern

		sl := strings.NewReader(string(raw))
		scanner := bufio.NewScanner(sl)
		for scanner.Scan() {
			// fmt.Println(scanner.Text())
			msg, err := cs2log.Parse(scanner.Text())
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
