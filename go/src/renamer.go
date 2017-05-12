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
	fmt.Println(pattern)
	files, _ := ioutil.ReadDir(folder)
	for _, f := range files {

		fileExt := filepath.Ext(f.Name())
		fmt.Println(fileExt)
		datefile := execExifTool(folder+"/"+f.Name())
		fmt.Println(datefile)
		newName := folder+"/"+datefile.Format(pattern)+fileExt
		os.Rename(folder+"/"+f.Name(), newName)
		fmt.Println("==================")

	}
}

func execExifTool(filepath string) time.Time {
	exifToolCmd := exec.Command("exiftool", "-j", filepath)
	exifToolOut, err := exifToolCmd.Output()
	if err != nil {
		panic(err)
	}
	return parseDate(string(exifToolOut))
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
