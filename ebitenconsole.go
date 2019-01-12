// Copyright 2019 by Magnus Wahlstrand. All rights reserved.
// Contact me for usage :-)

package ebitenconsole

import (
	"strconv"
	"strings"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"

	"github.com/pkg/errors"
)

type Cmd struct {
	Name        string
	Description string
	Value       Value
}

var commands = CmdSet{
	Usage:  func() {},
	values: make(map[string]*Cmd),
	funcs:  make(map[string]func() error),
}

func Parse(input string) error {
	s := strings.SplitN(input, "=", 2)

	// Handle set command
	if len(s) == 2 {
		cmd, value := s[0], s[1]
		return commands.set(cmd, value)
	}

	// Handle single command
	cmd := input
	if _, ok := commands.funcs[input]; !ok {
		return errors.Errorf("no such command '%s'", cmd)
	}

	// Run command
	return commands.funcs[input]()
}

type CmdSet struct {
	// Usage is the function called when an error occured or help is entered
	Usage  func()
	values map[string]*Cmd
	funcs  map[string]func() error
}

func (c *CmdSet) addCmd(v Value, name string, description string) {
	c.values[name] = &Cmd{
		Name:        name,
		Value:       v,
		Description: description,
	}
}

func (c *CmdSet) set(name string, value string) error {
	if _, ok := c.values[name]; !ok {
		return errors.Errorf("cmd '%s' is not defined", name)
	}

	return c.values[name].Value.Set(value)
}

// Value is an interface to the value of a command. Used for changing a variable
type Value interface {
	Set(string) error
}

type boolVar bool

func BoolVar(p *bool, name string, description string) {
	commands.addCmd((*boolVar)(p), name, description)
}

func (b *boolVar) Set(s string) error {
	v, err := strconv.ParseBool(s)
	if err != nil {
		return err
	}
	*b = boolVar(v)
	return nil
}

type stringVar string

func StringVar(v *string, name string, description string) {
	commands.addCmd((*stringVar)(v), name, description)
}

func (s *stringVar) Set(v string) error {
	*s = stringVar(v)
	return nil
}

type floatVar float64

func FloatVar(v *float64, name string, description string) {
	commands.addCmd((*floatVar)(v), name, description)
}

func (f *floatVar) Set(s string) error {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}
	*f = floatVar(v)
	return nil
}

func FuncVar(f func() error, name string, description string) {
	commands.funcs[name] = f
}

// Input
var capturingInput bool
var input string = ""
var result string
var resultTime time.Time

func CheckInput() {
	if capturingInput {
		check()
		return
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		result = ""
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		capturingInput = true
	}
}

func check() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		if input == "" {
			stopCatching()
			return
		}

		err := Parse(input)
		if err != nil {
			result = "ERR: " + err.Error()
		} else {
			result = "OK"
		}
		resultTime = time.Now()
		stopCatching()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		stopCatching()
	}

	for _, r := range ebiten.InputChars() {
		input += string(r)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) && len(input) > 0 {
		input = input[:len(input)-1]
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		input = ""
	}
}

func stopCatching() {
	input = ""
	capturingInput = false
}

const blinkPeriodMs = 500 //ms
func String() string {
	if !capturingInput {
		if time.Since(resultTime) > 2*time.Second {
			return ""
		}
		return result
	}

	var postfix string
	ms := time.Now().Nanosecond() / 1e6
	if (ms/blinkPeriodMs)%2 == 0 {
		postfix = "_"
	}

	return "> " + input + postfix
}
