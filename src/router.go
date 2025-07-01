package main

import (
	"errors"
	"fmt"
	"github.com/MattiaGaspa/asciiImage/asciiConverter"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const filesDir = "./files/"

func setupRouter(threads int) *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://asciinator.netlify.app/"},
		AllowMethods:     []string{"POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	}))
	r.MaxMultipartMemory = 10 << 20 // 10 MiB

	r.GET("/healthcheck", func(c *gin.Context) {
		c.String(http.StatusOK, "")
	})

	r.POST("/generate", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		xPost := c.DefaultQuery("x", "10")
		yPost := c.DefaultQuery("y", "14")
		x, y, err := validate(file, xPost, yPost)
		if err != nil {
			log.Printf("Validation error: %v", err)
			c.String(http.StatusBadRequest, fmt.Sprintf("Invalid parameters: %v", err))
			return
		}

		if err := c.SaveUploadedFile(file, filesDir+file.Filename); err != nil {
			log.Printf("Error saving file: %v", err)
			c.String(http.StatusInternalServerError, "Failed to save file")
			return
		}
		defer os.Remove(filesDir + file.Filename)

		output, err := asciiConverter.ConvertToASCII(filesDir+file.Filename, x, y, threads)
		if err != nil {
			log.Printf("Error converting file: %v", err)
			c.String(http.StatusInternalServerError, "Failed to convert file")
			return
		}
		c.String(http.StatusOK, output)
	})

	return r
}

func validate(file *multipart.FileHeader, x, y string) (int, int, error) {
	if !strings.HasSuffix(file.Filename, ".png") &&
		!strings.HasSuffix(file.Filename, ".jpg") &&
		!strings.HasSuffix(file.Filename, ".jpeg") {
		return 0, 0, errors.New("file must be a PNG or JPG image")
	}

	xInt, err := strconv.Atoi(x)
	if err != nil || xInt < 1 {
		return 0, 0, errors.New("invalid x")
	}
	yInt, err := strconv.Atoi(y)
	if err != nil || yInt < 1 {
		return 0, 0, errors.New("invalid y")
	}
	return xInt, yInt, nil
}
