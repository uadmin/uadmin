package uadmin

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

// TestCropImageHandler is a unit testing function for cropImageHandler() function
func TestCropImageHandler(t *testing.T) {
	// Create an 100 x 50 image
	img := image.NewRGBA(image.Rect(0, 0, 100, 50))
	c := color.RGBA{
		R: 255,
		G: 0,
		B: 0,
		A: 255,
	}
	for x := 0; x < 100; x++ {
		for y := 0; y < 50; y++ {
			img.Set(x, y, c)
		}
	}

	// Save to iamge.png
	f1, _ := os.OpenFile("./media/image_raw.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f1.Close()
	// defer os.Remove("./media/image.png")
	// defer os.Remove("./media/image_raw.png")
	png.Encode(f1, img)

	// Save to iamge.png
	f2, _ := os.OpenFile("./media/image_raw.jpg", os.O_WRONLY|os.O_CREATE, 0600)
	defer f2.Close()
	// defer os.Remove("./media/image.jpg")
	// defer os.Remove("./media/image_raw.jpg")
	jpeg.Encode(f2, img, nil)

	// Save to iamge.png
	f3, _ := os.OpenFile("./media/image_raw.gif", os.O_WRONLY|os.O_CREATE, 0600)
	defer f3.Close()
	// defer os.Remove("./media/image.gif")
	// defer os.Remove("./media/image_raw.gif")
	// o := gif.Options{}
	gif.Encode(f3, img, nil)

	examples := []struct {
		img    string
		top    int
		left   int
		bottom int
		right  int
	}{
		{img: "/media/image_raw.png", top: 10, left: 20, bottom: 40, right: 90},
		{img: "/media/image_raw.jpg", top: 10, left: 20, bottom: 40, right: 90},
		{img: "/media/image_raw.gif", top: 10, left: 20, bottom: 40, right: 90},
	}

	w := httptest.NewRecorder()
	URL := "/?img=%s&top=%d&left=%d&bottom=%d&right=%d"

	var croppedImg image.Image
	var err error
	var rect image.Rectangle
	var f *os.File
	for _, e := range examples {
		r := httptest.NewRequest("GET", fmt.Sprintf(URL, e.img, e.top, e.left, e.bottom, e.right), nil)
		cropImageHandler(w, r, nil)
		f, err = os.Open("." + strings.Replace(e.img, "_raw", "", -1))
		if err != nil {
			t.Errorf("Error in cropImageHandler. Cannot open file %s", e.img)
			continue
		}
		if strings.HasSuffix(e.img, "png") {
			croppedImg, err = png.Decode(f)
		}
		if strings.HasSuffix(e.img, "jpg") {
			croppedImg, err = jpeg.Decode(f)
		}
		if strings.HasSuffix(e.img, "gif") {
			croppedImg, err = gif.Decode(f)
		}
		if err != nil {
			t.Errorf("Error in cropImageHandler. Cannot decode file %s", e.img)
			continue
		}
		rect = croppedImg.Bounds()
		if rect.Dx() != (100 - e.left - (100 - e.right)) {
			t.Errorf("Invalid width in cropImageHandler for (%s)=%d expected %d", e.img, rect.Dx(), 100-e.left-(100-e.right))
		}
		if rect.Dy() != (50 - e.top - (50 - e.bottom)) {
			t.Errorf("Invalid height in cropImageHandler for (%s)=%d expected %d", e.img, rect.Dy(), 50-e.top-(50-e.bottom))
		}
	}
}
