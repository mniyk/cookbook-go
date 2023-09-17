package main

import (
	"fmt"
	"image/png"
	"os"
	"time"

	"github.com/kbinani/screenshot"
)

func main() {
	now := time.Now()
	tm := now.Format("20060102150405")

	if err := os.Mkdir(tm, 0777); err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(1)
	}

	n := screenshot.NumActiveDisplays()

	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)

		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			fmt.Printf("Error: %s", err)
			os.Exit(1)
		}

		x := bounds.Dx()
		y := bounds.Dy()

		fn := fmt.Sprintf("%s/%d_%dx%d.png", tm, i, x, y)
		f, _ := os.Create(fn)

		png.Encode(f, img)

		f.Close()
	}
}
