// Copyright 2015 Ghislain Guiot <gt.ghislain@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package termutil

import (
	"errors"
	"sync"
	"time"

	"github.com/nsf/termbox-go"
)

type Screen struct {
	sync.Mutex

	current *Window
	Windows []*Window

	timeout time.Duration

	Fg, Bg termbox.Attribute

	SizeX, SizeY int // size of screen

	EventFunc EventFunc // global event function

	quit chan bool
}

func New(timeout time.Duration) *Screen {

	if err := termbox.Init(); err != nil {
		panic(err)
	}

	sX, sY := termbox.Size()

	return &Screen{
		timeout: timeout,
		quit:    make(chan bool),
		SizeX:   sX,
		SizeY:   sY,
	}
}

func (s *Screen) End() {
	termbox.Close()
}

func (s *Screen) NewWindow() *Window {
	s.Lock()
	defer s.Unlock()

	sX, sY := termbox.Size()

	win := &Window{
		screen:     s,
		AutoResize: true,
		Fg:         s.Fg,
		Bg:         s.Bg,
		SizeX:      sX,
		SizeY:      sY,
	}

	s.Windows = append(s.Windows, win)
	s.current = win

	return win
}

func (s *Screen) Quit() {
	go func() {
		s.quit <- true
	}()
}

func (s *Screen) Focus(w *Window) {
	s.Lock()
	defer s.Unlock()

	s.current = w
}

func (s *Screen) Update() {

	for _, w := range s.Windows {
		w.update()
	}
}

func (s *Screen) Draw() error {

	if err := termbox.Clear(s.Fg, s.Bg); err != nil {
		return err
	}

	for _, w := range s.Windows {
		w.draw()
	}

	return nil
}

func (s *Screen) Run() (err error) {

	var ev termbox.Event
	tc := time.Tick(s.timeout)
	draw := true

	pe := make(chan termbox.Event)
	go func(pe chan termbox.Event) {
		for {
			pe <- termbox.PollEvent()
		}
	}(pe)

	for {
		if draw {
			s.Update()
			if err = s.Draw(); err != nil {
				return
			}
		}
		termbox.Flush()
		draw = false

		win := s.current
		if win == nil {
			return errors.New("no current window")
		}

		select {
		case <-s.quit:
			return

		case <-tc:
			draw = true

		case ev = <-pe:
			if ev.Type == termbox.EventError {
				return ev.Err
			}

			if ev.Type == termbox.EventResize {
				s.SizeX, s.SizeY = termbox.Size()
				if win.AutoResize {
					win.resize()
					draw = true
				}
			}

			if s.EventFunc != nil {
				s.EventFunc(ev)
				draw = true
			}
			if win.EventFunc != nil {
				win.EventFunc(ev)
				draw = true
			}
		}
	}
}
