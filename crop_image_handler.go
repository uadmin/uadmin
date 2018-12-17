package uadmin

import (
	"image"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type subImager interface {
	SubImage(r image.Rectangle) image.Image
}

func cropImageHandler(w http.ResponseWriter, r *http.Request, session *Session) {
	// image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)

	img := "." + r.FormValue("img")
	top, _ := strconv.ParseFloat(r.FormValue("top"), 32)
	left, _ := strconv.ParseFloat(r.FormValue("left"), 32)
	bottom, _ := strconv.ParseFloat(r.FormValue("bottom"), 32)
	right, _ := strconv.ParseFloat(r.FormValue("right"), 32)
	fType := strings.Split(img, ".")

	imageFile, err := os.Open(img)
	if err != nil {
		log.Println("ERROR: CropImage.Open1", err)
	}
	defer imageFile.Close()
	var myImage image.Image
	if strings.ToLower(fType[len(fType)-1]) == "jpg" || strings.ToLower(fType[len(fType)-1]) == "jpeg" {
		myImage, err = jpeg.Decode(imageFile)
		if err != nil {
			log.Println("ERROR: CropImage.Decode1", err)
		}
	}
	if strings.ToLower(fType[len(fType)-1]) == "png" {
		myImage, err = png.Decode(imageFile)
		if err != nil {
			log.Println("ERROR: CropImage.Decode2", err)
		}
	}

	if strings.ToLower(fType[len(fType)-1]) == "gif" {
		myImage, err = gif.Decode(imageFile)
		if err != nil {
			log.Println("ERROR: CropImage.Decode2", err)
		}
	}

	rect := image.Rect(int(left), int(top), int(right), int(bottom))

	mySubImage := image.NewRGBA(rect)
	draw.Draw(mySubImage, rect, myImage, image.Point{int(top), int(left)}, draw.Src)

	f, err := os.Create(strings.Replace(img, "_raw", "", -1))
	if err != nil {
		log.Println("ERROR: CropImage.Create", err)
	}
	defer f.Close()

	if strings.ToLower(fType[len(fType)-1]) == cJPG || strings.ToLower(fType[len(fType)-1]) == cJPEG {
		err = jpeg.Encode(f, mySubImage, nil)
		if err != nil {
			log.Println("ERROR: CropImage.Encode1", err)
		}
	}
	if strings.ToLower(fType[len(fType)-1]) == cPNG {
		err = png.Encode(f, mySubImage)
		if err != nil {
			log.Println("ERROR: CropImage.Encode2", err)
		}
	}

	if strings.ToLower(fType[len(fType)-1]) == cGIF {
		o := gif.Options{}
		err = gif.Encode(f, mySubImage, &o)
		if err != nil {
			log.Println("ERROR: CropImage.Encode2", err)
		}
	}

}
