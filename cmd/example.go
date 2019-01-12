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
	console := ebitenconsole.String()
	ebitenutil.DebugPrintAt(tmpImg, console, 5, 0)

	ebitenutil.DebugPrintAt(tmpImg, `
Press enter and type:
- color=red
- flip=true
- reset`, 5, 40)
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

	var i int
	reset := func() error {
		if i >= 3 {
			return errors.New("reset too many times (3)")
		}
		name = "Testman"
		clr = "black"
		flip = false
		i++
		return nil
	}
	ebitenconsole.FuncVar(reset, "reset", "reset all state")

	errorfunc := func() error { return errors.New("throws example") }
	ebitenconsole.FuncVar(errorfunc, "error", "throws error")

	if err := ebiten.Run(update, 200, 200, 2, "ebitenconsole demo"); err != nil {
		log.Fatal(err)
	}
}
