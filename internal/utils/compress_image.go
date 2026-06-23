package utils

import (
	"image"
	"image/jpeg"
	"os"
)

// Quality is the quality of compressed image, range from 1-100
func CompressImage(inputPath, outputPath string, quality int) error {
	file, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	options := &jpeg.Options{Quality: quality}

	return jpeg.Encode(outFile, img, options)
}
