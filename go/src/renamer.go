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
		fmt.Println("==================")

	}
}

func execExifTool(filepath string) time.Time {
	exifToolCmd := exec.Command("exiftool", "-j", filepath)
	exifToolOut, err := exifToolCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("exifTool "+filepath+"-f")


	fmt.Println(string(exifToolOut))
	//return parseDate(string(exifToolOut))

	return time.Now()
}

//func parseDate(exifTool string) time.Time {
//	byt := []byte(exifTool)
//	var dat map[string]interface{}
//
//	if err := json.Unmarshal(byt, &dat); err != nil {
//		panic(err)
//	}
//	fmt.Println(dat)
//	date := dat["CreateDate"].(time.Time)
//	fmt.Println(date)
//	return date
//}