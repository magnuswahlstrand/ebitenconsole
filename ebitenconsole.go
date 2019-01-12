// Copyright 2019 by Magnus Wahlstrand. All rights reserved.
// Contact me for usage :-)

package ebitenconsole

import (
	"fmt"
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
}

func Parse(input string) error {
	if strings.HasPrefix(input, "set ") {
		s := strings.SplitN(input[4:], "=", 2)

		if len(s) < 2 {
			fmt.Println(s)
			return errors.New("invalid 'set' command. use 'set <variable>=<value>'")
		}

		cmd, value := s[0], s[1]
		return commands.set(cmd, value)
	}

	return errors.New("invalid command")
}

type CmdSet struct {
	// Usage is the function called when an error occured or help is entered
	Usage  func()
	values map[string]*Cmd
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

func BoolVar(p *bool, name string, description string) {
	commands.addCmd((*boolVar)(p), name, description)
}

type boolVar bool

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

var capturingInput bool
var input string = "set name=Magnus"
var result string
var resultTime time.Time

func CheckInput() {
	if capturingInput {
		check()
		return
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		fmt.Println("Start catching input")
		capturingInput = true
	}
}

func check() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		err := Parse(input)
		if err != nil {
			result = " ERR: " + err.Error()
		} else {
			result = " OK"
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

	if ebiten.IsKeyPressed(ebiten.KeyBackspace) && len(input) > 0 {
		input = input[:len(input)-1]
	}
}

func stopCatching() {
	fmt.Println("Stop catching input")
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
	fmt.Println(ms)
	if (ms/blinkPeriodMs)%2 == 0 {
		postfix = "_"
	}

	return " > " + input + postfix
}
