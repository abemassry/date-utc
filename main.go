package main

import (
  "net/http"
  "log"
  "regexp"
  "os/exec"

  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  r.LoadHTMLGlob("templates/*")
  r.GET("/", func(c *gin.Context) {
	  t := c.Query("time")
	  var validTime = regexp.MustCompile(`^\d\d:\d\d$`)
	  if !validTime.MatchString(t) {
		t = ""
	  }
	  output := ""
	  out, err := exec.Command("./date-utc-utc", t).Output()
	  if err != nil {
		  log.Println(err)
	  }
	  output = string(out[:])
	  c.HTML(http.StatusOK, "index.tmpl", gin.H{
		  "title": "date-utc",
		  "time": t,
		  "output": output,
	  })
  })
  r.GET("/utc", func(c *gin.Context) {
	  t := c.Query("time")
	  var validTime = regexp.MustCompile(`^\d\d:\d\d$`)
	  if !validTime.MatchString(t) {
		t = ""
	  }
	  output := ""
	  out, err := exec.Command("./date-utc-utc", t).Output()
	  if err != nil {
		  log.Println(err)
	  }
	  output = string(out[:])
	  c.HTML(http.StatusOK, "index.tmpl", gin.H{
		  "title": "date-utc",
		  "time": t,
		  "output": output,
	  })
  })
  r.GET("/est", func(c *gin.Context) {
	  t := c.Query("time")
	  var validTime = regexp.MustCompile(`^\d\d:\d\d$`)
	  if !validTime.MatchString(t) {
		t = ""
	  }
	  output := ""
	  out, err := exec.Command("./date-utc-est", t).Output()
	  if err != nil {
		  log.Println(err)
	  }
	  output = string(out[:])
	  c.HTML(http.StatusOK, "index.tmpl", gin.H{
		  "title": "date-utc",
		  "time": t,
		  "output": output,
	  })
  })
  r.GET("/pst", func(c *gin.Context) {
	  t := c.Query("time")
	  var validTime = regexp.MustCompile(`^\d\d:\d\d$`)
	  if !validTime.MatchString(t) {
		t = ""
	  }
	  output := ""
	  out, err := exec.Command("./date-utc-pst", t).Output()
	  if err != nil {
		  log.Println(err)
	  }
	  output = string(out[:])
	  c.HTML(http.StatusOK, "index.tmpl", gin.H{
		  "title": "date-utc",
		  "time": t,
		  "output": output,
	  })
  })
  r.GET("/ist", func(c *gin.Context) {
	  t := c.Query("time")
	  var validTime = regexp.MustCompile(`^\d\d:\d\d$`)
	  if !validTime.MatchString(t) {
		t = ""
	  }
	  output := ""
	  out, err := exec.Command("./date-utc-ist", t).Output()
	  if err != nil {
		  log.Println(err)
	  }
	  output = string(out[:])
	  c.HTML(http.StatusOK, "index.tmpl", gin.H{
		  "title": "date-utc",
		  "time": t,
		  "output": output,
	  })
  })
  r.GET("/api", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })
  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
  //r.Run(":80")
}
