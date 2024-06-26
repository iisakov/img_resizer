package stl

import (
	"bufio"
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/nfnt/resize"
)

func GetImgPathSlice(path string) (result []string) {
	filepath.Walk(path, func(wPath string, info os.FileInfo, err error) error {

		if wPath == path {
			return nil
		}

		wPathSlice := strings.Split(wPath, string(os.PathSeparator))
		isImg, _ := regexp.MatchString(`.jpg$`, wPathSlice[len(wPathSlice)-1])
		isNotOriginal, _ := regexp.MatchString(`^(`+Prefix_b+`|`+Prefix_s+`)`, wPathSlice[len(wPathSlice)-1])

		if wPath != path && isImg && !isNotOriginal {
			result = append(result, wPath)
		}
		return nil
	})
	return
}

func ResizeImg(maxWidth, maxHight uint, imgPath, prefix string) (err error) {
	var imgFile, newImgFile *os.File

	imgFile, err = os.Open(imgPath)
	if err != nil {
		return
	}
	defer imgFile.Close()

	pathSlice := strings.Split(imgPath, string(os.PathSeparator))
	newImgPath := strings.Replace(imgPath, pathSlice[len(pathSlice)-1], prefix+pathSlice[len(pathSlice)-1], 1)

	newImgFile, err = os.Create(newImgPath)
	if err != nil {
		return
	}

	in := bufio.NewReader(imgFile)
	out := bufio.NewWriter(newImgFile)

	img, _, err := image.Decode(in)
	if err != nil {
		return
	}

	newImg := resize.Thumbnail(maxWidth, maxHight, img, resize.Lanczos3)

	err = jpeg.Encode(out, newImg, &jpeg.Options{Quality: 100})
	if err != nil {
		return
	}

	fmt.Println(MaxWidthBigImg, MaxHightBigImg, imgPath, Prefix_b)
	return
}

func DeleteImg(imgPath string) (err error) {
	err = os.Remove(imgPath)
	if err != nil {
		return
	}
	fmt.Printf("%s Удалён.\n\n", imgPath)
	return
}
