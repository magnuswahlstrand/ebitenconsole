package main

import (
	"errors"
	"fmt"
	"image/color"
	"log"
	"math"

	"golang.org/x/image/colornames"

	"github.com/hajimehoshi/ebiten/ebitenutil"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/kyeett/ebitenconsole"
)

var flip bool
var name, clr string

func update(screen *ebiten.Image) error {
	if inpututil.IsKeyJustPressed(ebiten.Key1) {
		return errors.New("Exit")
	}

	var c color.Color
	switch clr {
	case "red":
		c = colornames.Red
	case "blue":
		c = colornames.Blue
	case "black":
		c = colornames.Black
	case "green":
		c = colornames.Green
	case "yellow":
		c = colornames.Yellow
	default:
		c = color.White
	}
	screen.Fill(c)
	tmpImg, _ := ebiten.NewImage(200, 200, ebiten.FilterDefault)

	// Add this
	ebitenconsole.CheckInput()
	ebitenutil.DebugPrintAt(tmpImg, ebitenconsole.String(), 0, 0)

	ebitenutil.DebugPrintAt(tmpImg, "Press enter and type:\nset color=red\n", 5, 40)

	ebitenutil.DebugPrintAt(tmpImg, fmt.Sprintf("name: %s\ncolor: %s\nflip: %t\n", name, clr, flip), 5, 140)
	op := &ebiten.DrawImageOptions{}
	if flip {

		op.GeoM.Translate(-100, -100)
		op.GeoM.Rotate(math.Pi)
		op.GeoM.Translate(100, 100)
	}
	screen.DrawImage(tmpImg, op)
	return nil
}

func main() {
	fmt.Println("Start")

	name = "Testman"
	clr = "black"

	ebitenconsole.StringVar(&name, "name", "of player")
	ebitenconsole.StringVar(&clr, "color", "background color")
	ebitenconsole.BoolVar(&flip, "flip", "show/hide something")

	if err := ebiten.Run(update, 200, 200, 2, "ebitenconsole demo"); err != nil {
		log.Fatal(err)
	}
}
