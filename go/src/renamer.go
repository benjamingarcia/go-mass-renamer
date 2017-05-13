package main

import (
	//"github.com/rwcarlsen/goexif/exif"
	//"github.com/xiam/exif"
	"os"
	"fmt"
	"io/ioutil"
	"path/filepath"
	//"log"
	"os/exec"
	"time"
	"encoding/json"
)

func main() {
	argsWithoutProg := os.Args[1:]
	folder := argsWithoutProg[0]
	pattern := argsWithoutProg[1]
	rename(folder, pattern)
}

func rename(folder string, pattern string) {
	files, _ := ioutil.ReadDir(folder)
	for _, f := range files {

		if f.IsDir() {
			rename(folder+"/"+f.Name(), pattern)
		} else {
			fileExt := filepath.Ext(f.Name())
			fmt.Println(fileExt)
			datefile, err := execExifTool(folder + "/" + f.Name())
			if err != 1 {
				newName := folder + "/" + datefile.Format(pattern) + fileExt
				fmt.Println("Rename : " + f.Name() + " => " + datefile.Format(pattern) + fileExt)
				os.Rename(folder+"/"+f.Name(), newName)
			} else {
				fmt.Println("No exif file")
			}
			fmt.Println("==================")
		}

	}
}

func execExifTool(filepath string) (time.Time, int) {
	exifToolCmd := exec.Command("exiftool", "-j", filepath)
	exifToolOut, err := exifToolCmd.Output()
	if err != nil {
		return time.Now(), 1
	}
	return parseDate(string(exifToolOut)), 0
}

func parseDate(exifTool string) time.Time {
	exifToolOutput := exifTool[1:len(exifTool)-2]

	bytexif := []byte(exifToolOutput)
	var datexif map[string]interface{}
	if err := json.Unmarshal(bytexif, &datexif); err != nil {
		panic(err)
	}
	exifTimeLayout := "2006:01:02 15:04:05"
	fmt.Println(datexif["CreateDate"].(string))
	dateInDate, err := time.Parse(exifTimeLayout, datexif["CreateDate"].(string))
	if err != nil {
		panic(err)
	}

	//fmt.Println(date)
	return dateInDate
}
