package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
)

var (
	uiImage       *ebiten.Image
	uiFont        font.Face
	uiFontMHeight int
)

func init() {
	uiImage = ebiten.NewImage(1, 1)
	uiImage.Fill(color.White)
}

type Editor struct{}

func (e *Editor) Update() error {
	return nil
}

func (e *Editor) Draw(screen *ebiten.Image) {
	screen.DrawImage(uiImage, nil)
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
