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

// httpLogPattern is a regexp to parse the log line on HTTP
var httpLogPattern = regexp.MustCompile(`(\d{2}\/\d{2}\/\d{4} - \d{2}:\d{2}:\d{2}.\d{3}) - (.*)`)

type handler func(ip string, msg cs2log.Message)

// NewLogHandler returns a new LogHandler. This function has side effect to override log line prefix.
func NewLogHandler(h handler) LogHandler {
	// Override log line prefix
	cs2log.LogLinePattern = httpLogPattern
	return LogHandler{
		handler: h,
	}
}

// LogHandler is a handler for cs2-log HTTP.
type LogHandler struct {
	handler handler
}

// Handle returns a gin.HandlerFunc to handle cs2-log HTTP.
func (l *LogHandler) Handle() gin.HandlerFunc {
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
			l.handler(c.ClientIP(), msg)
		}
		c.String(http.StatusOK, "OK")
	}
}
