package service

import "github.com/Kagami/go-face"

const modelsDir = "models"

type Face struct {
	Samples []face.Descriptor
	Ids     []int32
	Names   []string
}

var Rec *face.Recognizer
var FaceData Face

func NewRecognise() {
	rec, err := face.NewRecognizer(modelsDir)
	if err != nil {
		panic("[REC INIT ERROR] : " + err.Error())
	}
	Rec = rec
}
