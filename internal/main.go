package main

import (
	"face-recognition/service"
	"github.com/gin-gonic/gin"
)

func main() {
	// -------------------------- INIT START--------------------------
	service.NewRecognise()
	// -------------------------- INIT END--------------------------
	r := gin.Default()
	r.POST("/enter/face-data", service.EnterFaceData)
	r.POST("/recognise/face", service.RecogniseFace)
	r.Run()
}
