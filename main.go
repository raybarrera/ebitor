package main

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
)

var (
	uiImage       = ebiten.NewImage(3, 3)
	uiSubImage    = uiImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)
	uiFont        font.Face
	uiFontMHeight int
)

func init() {
	uiImage.Fill(color.Black)
}

type Editor struct{}

func (e *Editor) Update() error {
	return nil
}

func (e *Editor) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xff, 0xff, 0xff, 0xff})
	screen.DrawImage(uiSubImage, nil)
}

func (e *Editor) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Ebitengine Editor")
	if err := ebiten.RunGame(&Editor{}); err != nil {
		panic(err)
	}
}
