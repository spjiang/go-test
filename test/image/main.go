package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"io/ioutil"
	"os"
	"path/filepath"
)

const dir_to_scan string = "/home/da/to_merge"

func main() {
	files, _ := ioutil.ReadDir(dir_to_scan)
	for _, imgFile := range files {

		if reader, err := os.Open(filepath.Join(dir_to_scan, imgFile.Name())); err == nil {
			defer reader.Close()
			im, _, err := image.DecodeConfig(reader)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s: %v\n", imgFile.Name(), err)
				continue
			}
			fmt.Printf("%s %d %d\n", imgFile.Name(), im.Width, im.Height)
		} else {
			fmt.Println("Impossible to open the file:", err)
		}
	}
}
