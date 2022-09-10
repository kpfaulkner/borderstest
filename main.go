package main

import (
	"fmt"
	"log"

	"github.com/kpfaulkner/borders/border"
	"github.com/kpfaulkner/borders/converters"
	"github.com/kpfaulkner/borders/image"
)

func main() {
	img, err := border.LoadImage("florida-big.png", false)

	// erode/dilate used to trim off any stray pixels.
	img2, err := image.Erode(img, 1)
	if err != nil {
		panic("BOOM on erode")
	}

	img3, err := image.Dilate(img2, 1)
	if err != nil {
		panic("BOOM on dilate")
	}

	// save just to help diagnose any problems.
	border.SaveImage("border.png", img3)

	// find contours of image
	cont := border.FindContours(img3)

	// save the contours as an image. Again, just for debugging.
	border.SaveContourSliceImage("contour.png", cont, img3.Width, img3.Height, false, 0)

	// convert to a slice of polygons.
	// slippyX, slippyY are hardcoded based off the original imput image.
	slippyConverter := converters.NewSlippyToLatLongConverter(1139408, 1772861, 22)

	// convert to polygons.
	poly, err := converters.ConvertContourToPolygon(cont, true, true, slippyConverter)
	if err != nil {
		log.Fatalf("Unable to convert to polygon : %s", err.Error())
	}

	b, _ := poly.MarshalJSON()
	fmt.Printf("%s\n", string(b))
}
