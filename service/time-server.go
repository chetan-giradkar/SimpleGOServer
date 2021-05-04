package service

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetTimeEpoch(c *gin.Context) {
	var million int64 = 1000000

	epoch := time.Now().UnixNano()
	millis := epoch / million

	c.JSON(http.StatusOK, millis)
}

func GetDateTime(c *gin.Context) {
	t := time.Now().Format("Mon Jan 2 03:04:05PM -0700 MST 2006")
	c.JSON(http.StatusOK, t)
}

func GetTimeTZ(c *gin.Context) {
	tz := c.Param("timezone")
	log.Println("timezone = ", tz)

	loc, loadLocationError := time.LoadLocation(tz)
	if loadLocationError != nil {
		log.Println("error = ", loadLocationError.Error())
		c.JSON(http.StatusBadRequest, loadLocationError.Error())
		c.Abort()
		return
	}

	log.Println("loc = ", loc)
	now := time.Now().In(loc)
	formattedTime := now.Format("Mon Jan 2 03:04:05PM -0700 MST 2006")
	c.JSON(http.StatusOK, formattedTime)
}
