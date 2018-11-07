package uadmin

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/nfnt/resize"
)

func processUpload(r *http.Request, f *F, modelName string, session *Session, s *ModelSchema) (val string) {
	// Get file description from http request
	httpFile, handler, err := r.FormFile(f.Name)
	if err != nil {
		return ""
	}
	defer httpFile.Close()

	// return "", s if there is no file uploaded
	if handler.Filename == "" {
		return ""
	}

	if handler.Size > MaxUploadFileSize {
		f.ErrMsg = fmt.Sprintf("File is too large. Maximum upload file size is: %d Mb", MaxUploadFileSize/1024/1024)
		return ""
	}

	// Get the upload to path and create it if it doesn't exist
	uploadTo := "/media/" + f.Type + "s/"
	if f.UploadTo != "" {
		uploadTo = f.UploadTo
	}
	if _, err = os.Stat("." + uploadTo); os.IsNotExist(err) {
		err = os.MkdirAll("."+uploadTo, os.ModePerm)
		if err != nil {
			Trail(ERROR, "processForm.MkdirAll. %s", err)
			return ""
		}
	}

	// Generate locacl file name and create it
	fParts := strings.Split(handler.Filename, ".")
	fExt := strings.ToLower(fParts[len(fParts)-1])
	var fName string
	var pathName string
	pathName = "." + uploadTo + modelName + "_" + f.Name + "_" + GenerateBase64(10) + "/"
	if f.Type == cIMAGE && len(fParts) > 1 {
		fName = strings.TrimSuffix(handler.Filename, "."+fExt) + "_raw." + fExt
	} else {
		f.ErrMsg = "Image file with no extension. Please use png, jpg, jpeg or gif."
		return ""
	}
	if f.Type == cFILE {
		fName = handler.Filename
	}
	for _, err = os.Stat(pathName + fName); os.IsExist(err); {
		pathName = "." + uploadTo + modelName + "_" + f.Name + "_" + GenerateBase64(10) + "/"
		//fName = "." + uploadTo + modelName + "_" + f.Name + "_" + GenerateBase64(10) + "_raw." + fExt
	}
	fName = pathName + fName
	err = os.MkdirAll(pathName, os.ModePerm)
	if err != nil {
		Trail(ERROR, "processForm.MkdirAll. unable to create folder for uploaded file. %s", err)
		return ""
	}
	fRaw, err := os.OpenFile(fName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		Trail(ERROR, "processForm.OpenFile. unable to create file. %s", err)
		return ""
	}

	// Copy http file to local
	_, err = io.Copy(fRaw, httpFile)
	if err != nil {
		Trail(ERROR, "ProcessForm error uploading http file. %s", err)
		return ""
	}
	defer fRaw.Close()

	// store the file path to DB
	if f.Type == cFILE {
		val = fmt.Sprint(strings.TrimPrefix(fName, "."))
	} else {
		// If case it is an image, process it first
		fRaw, err = os.Open(fName)
		if err != nil {
			Trail(ERROR, "ProcessForm.Open %s", err)
			return ""
		}

		// decode jpeg,png,gif into image.Image
		var img image.Image
		if fExt == cJPG || fExt == cJPEG {
			img, err = jpeg.Decode(fRaw)
		} else if fExt == cPNG {
			img, err = png.Decode(fRaw)
		} else if fExt == cGIF {
			img, err = gif.Decode(fRaw)
		} else {
			f.ErrMsg = "Unknown image file extension. Please use, png, jpg/jpeg or gif"
			return ""
		}
		if err != nil {
			f.ErrMsg = "Unknown image format or image corrupted."
			Trail(WARNING, "ProcessForm.Decode %s", err)
			return ""
		}

		// Resize the image to fit max height, max width
		width := img.Bounds().Dx()
		height := img.Bounds().Dy()
		if height > MaxImageHeight {
			Ratio := float64(MaxImageHeight) / float64(height)
			// Ratio = math.Min(Ratio, 1)
			width = int(float64(width) * Ratio)
			height = int(float64(height) * Ratio)
			if width > MaxImageWidth {
				Ratio = float64(MaxImageWidth) / float64(width)
				// Ratio = math.Min(Ratio, 1)
				width = int(float64(width) * Ratio)
				height = int(float64(height) * Ratio)
			}
			img = resize.Resize(uint(width), uint(height), img, resize.Lanczos3)
		}

		// Store the active file
		fActiveName := strings.Replace(fName, "_raw", "", -1)
		fActive, err := os.Create(fActiveName)
		if err != nil {
			Trail(ERROR, "ProcessForm.Create unable to create file for resized image. %s", err)
			return ""
		}
		defer fActive.Close()

		fRaw, err = os.OpenFile(fName, os.O_WRONLY, 0644)
		if err != nil {
			Trail(ERROR, "ProcessForm.Open %s", err)
			return ""
		}
		defer fRaw.Close()

		// write new image to file
		if fExt == cJPG || fExt == cJPEG {
			err = jpeg.Encode(fActive, img, nil)
			if err != nil {
				Trail(ERROR, "ProcessForm.Encode active jpg. %s", err)
				return ""
			}

			err = jpeg.Encode(fRaw, img, nil)
			if err != nil {
				Trail(ERROR, "ProcessForm.Encode raw jpg. %s", err)
				return ""
			}
		}

		if fExt == cPNG {
			err = png.Encode(fActive, img)
			if err != nil {
				Trail(ERROR, "ProcessForm.Encode active png. %s", err)
				return ""
			}

			err = png.Encode(fRaw, img)
			if err != nil {
				Trail(ERROR, "ProcessForm.Encode raw png. %s", err)
				return ""
			}
		}

		if fExt == cGIF {
			o := gif.Options{}
			err = gif.Encode(fActive, img, &o)
			if err != nil {
				Trail(ERROR, "ProcessForm.Encode active gif. %s", err)
				return ""
			}

			err = gif.Encode(fRaw, img, &o)
			if err != nil {
				Trail(ERROR, "ProcessForm.Encode raw gif. %s", err)
				return ""
			}
		}
		val = fmt.Sprint(strings.TrimPrefix(fActiveName, "."))
	}
	return val
}
