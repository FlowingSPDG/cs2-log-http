package cs2loghttp

import (
	"bufio"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	cs2log "github.com/janstuemmel/cs2-log"
)

var httpLogPattern = regexp.MustCompile(`(\d{2}\/\d{2}\/\d{4} - \d{2}:\d{2}:\d{2}.\d{3}) - (.*)`)

func CS2Logger(Handler func(cs2log.Message, *gin.Context)) gin.HandlerFunc {
	// Override log line prefix
	cs2log.LogLinePattern = httpLogPattern

	return func(c *gin.Context) {
		raw, err := c.GetRawData()
		if err != nil {
			log.Printf("Failed to get raw data : %v\n", err)
			c.String(http.StatusInternalServerError, err.Error())
			c.Abort()
			return
		}

		sl := strings.NewReader(string(raw))
		scanner := bufio.NewScanner(sl)
		for scanner.Scan() {
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
