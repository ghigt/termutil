// Copyright 2015 Ghislain Guiot <gt.ghislain@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package termutil

import (
	"errors"
	"time"

	"github.com/nsf/termbox-go"
)

// MainWindow controls all windows inside it.
type MainWindow struct {
	current *Window
	windows []*Window

	timeout time.Duration

	Fg, Bg termbox.Attribute

	EventFunc EventFunc // global event function

	quit chan bool
}

var Screen *MainWindow

func Init(timeout time.Duration) error {

	if err := termbox.Init(); err != nil {
		return err
	}

	Screen = &MainWindow{
		timeout: timeout,
		quit:    make(chan bool),
	}

	return nil
}

func End() {
	termbox.Close()
}

func Quit() {
	go func() {
		Screen.quit <- true
	}()
}

func Focus(w *Window) {
	Screen.current = w
}

func update() {
	for _, w := range Screen.windows {
		w.update()
	}
}

func draw() error {

	if err := termbox.Clear(Screen.Fg, Screen.Bg); err != nil {
		return err
	}

	for _, w := range Screen.windows {
		w.draw()
	}

	return nil
}

func resize(ws []*Window) (redraw bool) {
	for _, w := range ws {
		if w.AutoResize {
			w.resize()
			redraw = true
		}
	}
	return
}

func Run() (err error) {

	var ev termbox.Event
	tc := time.Tick(Screen.timeout)
	redraw := true

	pe := make(chan termbox.Event)
	go func(pe chan termbox.Event) {
		for {
			pe <- termbox.PollEvent()
		}
	}(pe)

	for {
		if redraw {
			update()
			if err = draw(); err != nil {
				return
			}
		}
		termbox.Flush()
		redraw = false

		win := Screen.current
		if win == nil {
			return errors.New("no current window")
		}

		select {
		case <-Screen.quit:
			return

		case <-tc:
			redraw = true

		case ev = <-pe:
			if ev.Type == termbox.EventError {
				return ev.Err
			}

			if ev.Type == termbox.EventResize {
				if resize(Screen.windows) {
					redraw = true
				}
			}

			if Screen.EventFunc != nil {
				Screen.EventFunc(ev)
				redraw = true
			}
			if win.EventFunc != nil {
				win.EventFunc(ev)
				redraw = true
			}
		}
	}
}
