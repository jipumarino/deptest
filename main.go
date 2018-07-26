package main

import (
	"fmt"
	"io/ioutil"
	"os"

	imagequant "github.com/ultimate-guitar/go-imagequant"
)

func main() {
	sourceFh, _ := os.OpenFile("orig.png", os.O_RDONLY, 0444)
	defer sourceFh.Close()
	srcImage, _ := ioutil.ReadAll(sourceFh)
	fmt.Println("Source size:", len(srcImage))
	targetImage, _ := imagequant.Crush(srcImage, 10, -3)
	fmt.Println("Target size:", len(targetImage))
}
