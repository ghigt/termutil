// Copyright 2015 Ghislain Guiot <gt.ghislain@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package main

import (
	"log"
	"time"

	"github.com/ghigt/termutil"
	"github.com/nsf/termbox-go"
)

var colors = []termbox.Attribute{
	termbox.ColorDefault, termbox.ColorBlack, termbox.ColorRed,
	termbox.ColorGreen, termbox.ColorYellow, termbox.ColorBlue,
	termbox.ColorMagenta, termbox.ColorCyan, termbox.ColorWhite,
}

func main() {

	var err error

	err = termutil.Init(time.Second)
	if err != nil {
		log.Fatal(err)
	}

	termutil.Screen.EventFunc = func(ev termbox.Event) {
		switch ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				termutil.Quit()
			}
		}
	}

	win1 := termutil.NewWindow()
	win1.X = 3
	win1.Y = 2

	win1.UpdateFunc = func() []string {
		return []string{
			"Put the text in the box:",
			"+---+",
			"|   |",
			"|   |",
			"|   |",
			"+---+",
		}
	}
	swin1 := win1.NewSubWindow()
	swin1.X = 1
	swin1.Y = 2
	swin1.UpdateFunc = func() []string {
		return []string{
			"+-+",
			"| |",
			"+-+",
		}
	}
	sswin := swin1.NewSubWindow()
	sswin.X = 1
	sswin.Y = 1
	sswin.UpdateFunc = func() []string {
		return []string{
			"*",
		}
	}

	win2 := termutil.NewWindow()
	win2.X = 15
	win2.Y = 15
	win2.SizeX = 3
	win2.SizeY = 3

	win2.UpdateFunc = func() []string {
		if win2.X == 4 && win2.Y == 4 {
			return []string{"you", "won", "!!!"}
		}
		return []string{"aaa", "bbb", "ccc"}
	}

	win2.EventFunc = func(ev termbox.Event) {
		switch ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowDown:
				win2.Y += 1
			case termbox.KeyArrowUp:
				win2.Y -= 1
			case termbox.KeyArrowLeft:
				win2.X -= 1
			case termbox.KeyArrowRight:
				win2.X += 1
			}
		}
	}

	err = termutil.Run()
	termutil.End()

	if err != nil {
		log.Fatal(err)
		return
	}
}
