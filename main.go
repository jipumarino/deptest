package main

import (
	"fmt"
	"io/ioutil"
	"os"

	imagequant "github.com/ultimate-guitar/go-imagequant"
)

func main() {

	fmt.Println("starting")

	sourceFh, err := os.OpenFile("orig.png", os.O_RDONLY, 0444)
	if err != nil {
		return fmt.Errorf("os.OpenFile: %s", err.Error())
	}
	defer sourceFh.Close()

	image, err := ioutil.ReadAll(sourceFh)
	if err != nil {
		return fmt.Errorf("ioutil.ReadAll: %s", err.Error())
	}

	optiImage, err := imagequant.Crush(image, speed, compression)
	if err != nil {
		return fmt.Errorf("imagequant.Crush: %s", err.Error())
	}

	destFh, err := os.OpenFile("dest.png", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("os.OpenFile: %s", err.Error())
	}
	defer destFh.Close()

	destFh.Write(optiImage)
	return nil

	fmt.Println("finished")

}
