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

var bodys = [][][]string{
	{
		{"Jean", "14", "I like sport...", "158"},
		{"Patrick", "46", "I travel", "190"},
		{"Jeremy", "54", "I read books", "179"},
	},
	{
		{"Jean", "53", "I like sport... a lot", "138"},
		{"Patrick", "102", "I don't travel", "185"},
		{"Jeremy", "43", "I eat books", "123"},
	},
	{
		{"Jean", "44", "I play football...", "160"},
		{"Patrick", "12", "I fly", "180"},
		{"Jeremy", "24", "I'm a books", "143"},
	},
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

	win.UpdateFunc = func() []string {
		wg := &termutil.WidgetTable{
			Header: []termutil.HeaderInfo{
				{"name", 20},
				{"age", 10},
				{"hobby", 0},
				{"height", 10},
			},
		}
		wg.Body = bodys[time.Now().Nanosecond()%3]
		return wg.Update(win)
	}

	err = termutil.Run()
	termutil.End()

	if err != nil {
		log.Fatal(err)
		return
	}
}
