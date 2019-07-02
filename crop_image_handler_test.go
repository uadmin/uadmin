package uadmin

import (
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
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

	os.MkdirAll("./media/user", 0744)

	// Save to iamge.png
	f1, _ := os.OpenFile("./media/user/image_raw.png", os.O_WRONLY|os.O_CREATE, 0600)
	png.Encode(f1, img)
	f1.Close()

	// Save to iamge.png
	f2, _ := os.OpenFile("./media/user/image_raw.jpg", os.O_WRONLY|os.O_CREATE, 0600)
	jpeg.Encode(f2, img, nil)
	f2.Close()

	// Save to iamge.png
	f3, _ := os.OpenFile("./media/user/image_raw.gif", os.O_WRONLY|os.O_CREATE, 0600)
	gif.Encode(f3, img, nil)
	f3.Close()

	// Save Bad Images
	f1, _ = os.OpenFile("./media/user/bad_image.png", os.O_WRONLY|os.O_CREATE, 0600)
	f1.Close()
	f1, _ = os.OpenFile("./media/user/bad_image.jpg", os.O_WRONLY|os.O_CREATE, 0600)
	f1.Close()
	f1, _ = os.OpenFile("./media/user/bad_image.gif", os.O_WRONLY|os.O_CREATE, 0600)
	f1.Close()

	s1 := &Session{
		UserID: 1,
		Active: true,
	}
	s1.GenerateKey()
	s1.Save()
	Preload(s1)

	u1 := User{
		Username: "u1",
		Password: "u1",
		Active:   true,
	}
	u1.Save()

	s2 := &Session{
		UserID: u1.ID,
		Active: true,
	}
	s2.GenerateKey()
	s2.Save()
	Preload(s2)

	examples := []struct {
		img     string
		top     int
		left    int
		bottom  int
		right   int
		status  string
		session *Session
	}{
		{img: "/media/user/image_raw.png", top: 10, left: 20, bottom: 40, right: 90, status: "ok", session: s1},
		{img: "/media/user/image_raw.jpg", top: 10, left: 20, bottom: 40, right: 90, status: "ok", session: s1},
		{img: "/media/user/image_raw.gif", top: 10, left: 20, bottom: 40, right: 90, status: "ok", session: s1},
		{img: "image_raw.gif", top: 10, left: 20, bottom: 40, right: 90, status: "error", session: s1},
		{img: "/media/user/image_raw.gif", top: 10, left: 20, bottom: 40, right: 90, status: "error", session: s2},
		{img: "/media/user/nofile.png", top: 10, left: 20, bottom: 40, right: 90, status: "error", session: s1},
		{img: "/media/user/bad_image.png", top: 10, left: 20, bottom: 40, right: 90, status: "error", session: s1},
		{img: "/media/user/bad_image.jpg", top: 10, left: 20, bottom: 40, right: 90, status: "error", session: s1},
		{img: "/media/user/bad_image.gif", top: 10, left: 20, bottom: 40, right: 90, status: "error", session: s1},
	}

	w := httptest.NewRecorder()
	URL := "/?img=%s&top=%d&left=%d&bottom=%d&right=%d"

	var croppedImg image.Image
	var err error
	var rect image.Rectangle
	var f *os.File

	for _, e := range examples {
		r := httptest.NewRequest("GET", fmt.Sprintf(URL, e.img, e.top, e.left, e.bottom, e.right), nil)
		cropImageHandler(w, r, e.session)

		// Read the response
		buf, _ := ioutil.ReadAll(w.Body)
		jObj := map[string]string{}
		err = json.Unmarshal(buf, &jObj)
		if err != nil {
			t.Errorf("cropImageHandler response is not valid JSON. %s", err)
			continue
		}
		if jObj["status"] != e.status {
			t.Errorf("cropImageHandler returned invalid status. Got %s expected %s", jObj["status"], e.status)
			continue
		}
		if jObj["status"] == "error" {
			continue
		}

		f, err = os.Open("." + strings.Replace(e.img, "_raw", "", -1))
		if err != nil {
			t.Errorf("Error in cropImageHandler. Cannot open file %s. %s", e.img, err)
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

	// Clean up
	Delete(s1)
	Delete(u1)
	Delete(s2)
}
