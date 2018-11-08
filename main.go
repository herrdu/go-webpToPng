package main

import (
	"fmt"
	"image/png"
	"io/ioutil"
	"os"
	"regexp"

	"golang.org/x/image/webp"
)

func main() {
	var folderDir = "./src"
	fileInfoList, err := ioutil.ReadDir(folderDir)
	if err != nil {
		fmt.Println(err)
	}

	for _, fileInfo := range fileInfoList {
		var fileName = fileInfo.Name()
		fmt.Println(fileName)
		isMatch, _ := regexp.MatchString(".webp$", fileName)
		if !isMatch {
			continue
		}
		webpFilename := folderDir + "/" + fileName
		f0, err := os.Open(webpFilename)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f0.Close()
		img0, err := webp.Decode(f0)
		if err != nil {
			fmt.Println(err)
			fmt.Println(err)
			return
		}
		pngFile, err := os.Create("./output/" + fileName + ".png")
		if err != nil {
			fmt.Println(err)
		}
		err = png.Encode(pngFile, img0)
		if err != nil {
			fmt.Println(err)
		}
	}
}
