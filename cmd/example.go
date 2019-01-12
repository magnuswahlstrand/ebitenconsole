package main

import (
	"errors"
	"fmt"
	"image/color"
	"log"

	"golang.org/x/image/colornames"

	"github.com/hajimehoshi/ebiten/ebitenutil"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/kyeett/ebitenconsole"
)

var toggle bool
var name, clr string

func update(screen *ebiten.Image) error {
	if inpututil.IsKeyJustPressed(ebiten.Key1) {
		return errors.New("Exit")
	}

	var c color.Color
	switch clr {
	case "red":
		c = colornames.Red
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

	ebitenconsole.CheckInput()

	ebitenutil.DebugPrint(screen, ebitenconsole.String())
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("name: %s\ncolor: %s\ntoggle: %t\n", name, clr, toggle), 5, 140)
	return nil
}

func main() {
	fmt.Println("Start")

	name = "Testman"
	clr = "black"

	ebitenconsole.StringVar(&name, "name", "of player")
	ebitenconsole.StringVar(&clr, "color", "background color")
	ebitenconsole.BoolVar(&toggle, "toggle", "show/hide something")

	if err := ebiten.Run(update, 200, 200, 2, "ebitenconsole demo"); err != nil {
		log.Fatal(err)
	}
}
