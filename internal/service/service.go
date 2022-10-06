package service

import (
	"face-recognition/define"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

// EnterFaceData 录入人脸数据
func EnterFaceData(c *gin.Context) {
	name := c.PostForm("name")
	if name == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "必填参不能为空",
		})
		return
	}
	fh, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "系统异常：" + err.Error(),
		})
		return
	}
	f, _ := fh.Open()
	b, _ := ioutil.ReadAll(f)

	face, err := Rec.RecognizeSingle(b)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "系统异常：" + err.Error(),
		})
		return
	}
	if face == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "未检出到人脸数据",
		})
		return
	}
	// 判断人脸数据是否存在
	id := Rec.ClassifyThreshold(face.Descriptor, define.Tolerance)
	if id > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据已存在，无需重复录入",
		})
		return
	}
	// 录入人脸数据
	FaceData.Samples = append(FaceData.Samples, face.Descriptor)
	FaceData.Ids = append(FaceData.Ids, int32(len(FaceData.Ids)+1))
	FaceData.Names = append(FaceData.Names, name)
	Rec.SetSamples(FaceData.Samples, FaceData.Ids)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "数据录入成功",
	})
}

// RecogniseFace 人脸识别
func RecogniseFace(c *gin.Context) {
	fh, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "系统异常：" + err.Error(),
		})
		return
	}
	f, _ := fh.Open()
	b, _ := ioutil.ReadAll(f)

	face, err := Rec.RecognizeSingle(b)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "系统异常：" + err.Error(),
		})
		return
	}
	id := Rec.ClassifyThreshold(face.Descriptor, define.Tolerance)
	if id < 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "人脸数据不存在",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "识别成功",
		"data": gin.H{
			"id":   id,
			"name": FaceData.Names[id-1],
		},
	})
}
