// Copyright 2015 Ghislain Guiot <gt.ghislain@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package termutil

import (
	"github.com/nsf/termbox-go"
)

type UpdateFunc func() []string

type EventFunc func(termbox.Event)

type Window struct {
	screen *Screen

	X, Y       int
	Fg, Bg     termbox.Attribute
	AutoResize bool
	rows       []string

	UpdateFunc UpdateFunc
	EventFunc  EventFunc
}

func (w *Window) Quit() {
	go func() {
		w.screen.quit <- true
	}()
}

func (w *Window) update() {
	w.rows = w.UpdateFunc()
}

func (w *Window) draw() {
	for y := w.Y; y-w.Y < len(w.rows); y++ {
		for x := w.X; x-w.X < len(w.rows[y-w.Y]); x++ {
			termbox.SetCell(x, y, rune(w.rows[y-w.Y][x-w.X]), w.Fg, w.Bg)
		}
	}
}
