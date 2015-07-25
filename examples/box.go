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

	screen := termutil.New(time.Second)

	screen.EventFunc = func(ev termbox.Event) {
		switch ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				screen.Quit()
			}
		}
	}

	win1 := screen.NewWindow()
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

	win2 := screen.NewWindow()
	win2.X = 15
	win2.Y = 15

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

	err := screen.Run()
	screen.End()

	if err != nil {
		log.Fatal(err)
		return
	}
}
