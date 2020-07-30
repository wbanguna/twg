package example_test

import (
	"fmt"
	"io"
	"testwithgo/example"

	// Needed for initialize side effect
	_ "image/png"
)

var file string = "this is not used"

func Example_crop() {
	var r io.Reader
	//f, err := os.Open("img.png")
	//if err != nil {
	//	panic(err)
	//}

	img, err := example.Decode(r)
	if err != nil {
		panic(err)
	}
	err = example.Crop(img, 0, 0, 20, 20)
	if err != nil {
		panic(err)
	}

	var w io.Writer
	//out, err := os.Create("out.jpg")
	//if err != nil {
	//	panic(err)
	//}
	err = example.Encode(img, w)
	if err != nil {
		panic(err)
	}
	fmt.Println("See out.jpg for the cropped image.")
	// Output:
	// See out.jpg for the cropped image.
}
