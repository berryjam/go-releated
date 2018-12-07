package methodsinterfaces

import (
	"golang.org/x/tour/pic"
	"image/color"
	"image"
)

type Image struct {

}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 100, 100)
}

func (img Image) At(x, y int) color.Color {
	return color.Black
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
