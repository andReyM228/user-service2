package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"testing"
)

func TestB64(t *testing.T) {
	imagePath := "C:\\Users\\admin\\Desktop\\images\\загружено1.jpg"
	imageBytes, err := ioutil.ReadFile(imagePath)
	if err != nil {
		log.Fatal(err)
	}

	// Конвертирование изображения в base64
	imageBase64 := base64.StdEncoding.EncodeToString(imageBytes)

	// Печать строки base64
	fmt.Println(imageBase64)
}
