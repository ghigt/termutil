// Copyright 2015 Ghislain Guiot <gt.ghislain@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package termutil

import "github.com/nsf/termbox-go"

type UpdateFunc func() []string

type EventFunc func(termbox.Event)

type Window struct {
	X, Y         int
	SizeX, SizeY int
	Fg, Bg       termbox.Attribute
	// AutoResize does not allow a fix SizeX and SizeY
	AutoResize bool
	rows       []string

	UpdateFunc UpdateFunc
	EventFunc  EventFunc

	parent     *Window
	SubWindows []*Window
}

func (w *Window) update() {

	if w.UpdateFunc != nil {
		w.rows = w.UpdateFunc()
	}

	for _, sub := range w.SubWindows {
		sub.update()
	}
}

func (w *Window) draw() {

	w.drawWin()

	for _, sub := range w.SubWindows {
		sub.draw()
	}
}

func (w *Window) drawWin() {

	ax := w.AbsX()
	ay := w.AbsY()

	termX, termY := termbox.Size()

	sizeY := termY - ay

	if sizeY > w.SizeY {
		sizeY = w.SizeY
	}
	if sizeY > len(w.rows) {
		sizeY = len(w.rows)
	}

	for y := ay; y-ay < sizeY; y++ {
		sizeX := termX - ax

		if sizeX > w.SizeX {
			sizeX = w.SizeX
		}
		if sizeX > len(w.rows[y-ay]) {
			sizeX = len(w.rows[y-ay])
		}
		for x := ax; x-ax < sizeX; x++ {
			termbox.SetCell(x, y, rune(w.rows[y-ay][x-ax]), w.Fg, w.Bg)
		}
	}
}

func (w *Window) resize() {

	if w.AutoResize {
		if w.parent == nil {
			w.SizeX, w.SizeY = termbox.Size()
		} else {
			w.SizeX = w.parent.SizeX - w.AbsX()
			w.SizeY = w.parent.SizeY - w.AbsY()
		}
	}

	for _, sub := range w.SubWindows {
		sub.resize()
	}
}

func (w *Window) AbsX() int {
	var x int

	if w.parent != nil {
		x = w.parent.AbsX()
	}
	return w.X + x
}

func (w *Window) AbsY() int {
	var y int

	if w.parent != nil {
		y = w.parent.AbsY()
	}
	return w.Y + y
}

func (w *Window) NewSubWindow() *Window {

	win := &Window{
		parent:     w,
		Fg:         w.Fg,
		Bg:         w.Bg,
		SizeX:      w.SizeX,
		SizeY:      w.SizeY,
		AutoResize: true,
	}

	w.SubWindows = append(w.SubWindows, win)

	return win
}

func NewWindow() *Window {

	win := &Window{
		AutoResize: true,

		Fg: Screen.Fg,
		Bg: Screen.Bg,
	}

	win.SizeX, win.SizeY = termbox.Size()

	Screen.windows = append(Screen.windows, win)
	Screen.current = win

	return win
}
