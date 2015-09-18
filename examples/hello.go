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

	win := termutil.NewWindow()

	win.UpdateFunc = func() []termutil.Row {
		return termutil.StringsToRows([]string{"hello"}, termbox.ColorCyan, termbox.ColorMagenta)
	}

	err = termutil.Run()
	termutil.End()

	if err != nil {
		log.Fatal(err)
		return
	}
}
