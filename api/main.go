package main

import (
	"flag"
	"net/http"
	"os"
	"os/signal"
	"path"
	"runtime"
	"syscall"

	"github.com/Sirupsen/logrus"
	//"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

var (
	port string
)

var dummyFunc = func(c *gin.Context, userID, limit string) {
	c.JSON(200, nil)
}

// QueryFunc is mapping for controller
type QueryFunc struct {
	Type string
	Func func(c *gin.Context, param, limit string)
}

func execFunc(c *gin.Context, qfs []QueryFunc) {
	for i := range qfs {
		if q := c.Query(qfs[i].Type); q != "" {
			qfs[i].Func(c, q, c.Query("limit"))
			return
		}
	}
	return
}

func getTemplatePath(file string) string {
	return path.Join("templates", file)
}

func render(c *gin.Context, status int, file string, data interface{}) {
	c.HTML(status, file, data)
}

func main() {
	flag.StringVar(&port, "port", "3000", "Port number")
	flag.Parse()

	isProfiling := -1
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGUSR2)
	go func() {
		for {
			s := <-sigchan
			if s == syscall.SIGUSR2 {
				if isProfiling == -1 {
					logrus.Info("Start the profiling server")
					runtime.SetBlockProfileRate(1)
					isProfiling = 0
					go http.ListenAndServe(":6060", nil)
				} else if isProfiling == 0 {
					logrus.Info("Stop the profiling")
					runtime.SetBlockProfileRate(0)
					isProfiling = 1
				} else {
					logrus.Info("Restart the profiling")
					runtime.SetBlockProfileRate(1)
					isProfiling = 0
				}
			}
		}
	}()

	// Set gin mode to release
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	// Set redirect option
	r.RedirectFixedPath = true

	// Use gin recovery to catch panic error after controller execution
	r.Use(gin.Recovery())

	// Gzip compression
	// r.Use(middleware.Gzip(middleware.DefaultCompression))

	r.LoadHTMLGlob("templates/*")

	r.Use(gin.Logger())

	r.GET("/alive", GetAlive)

	r.GET("/sample.html", GetFriends)

	// FOR PRD
	r.GET("/searchUser", SearchUser)

	r.GET("/searchItem", SearchItem)

	r.GET("/searchPost")

	r.GET("/user/:userId", UserPage)

	logrus.WithField("port", port).Info("Starting the server")
	http.ListenAndServe(":"+port, r)
}
