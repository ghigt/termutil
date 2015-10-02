// Copyright 2015 Ghislain Guiot <gt.ghislain@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package termutil

import (
	"github.com/nsf/termbox-go"
)

type WidgetTable struct {
	Header *Header
	Body   *Body
}

func (wg *WidgetTable) Update(win *Window) []Row {

	wg.Header.fillSpace()

	var out []Row

	out = append(out, wg.Header.buildRow(win.SizeX))
	out = append(out, wg.Body.buildRows(wg.Header.Titles, win.SizeX, win.SizeY-1)...)

	return out
}

func (wg *WidgetTable) EventFunc(ev termbox.Event) {
	switch ev.Type {
	case termbox.EventKey:
		switch ev.Key {
		case termbox.KeyArrowDown:
			if wg.Body.Active < len(wg.Body.Data)-1 {
				wg.Body.Active += 1
			}
		case termbox.KeyArrowUp:
			if wg.Body.Active > 0 {
				wg.Body.Active -= 1
			}
		}
	}
}

type Header struct {
	Titles []HeaderTitle

	Fg termbox.Attribute
	Bg termbox.Attribute

	// Colors for the active tab
	FgActive termbox.Attribute
	BgActive termbox.Attribute

	Active   int
	IsActive bool // Is a tab activated
}

func (h Header) buildRow(sizeX int) Row {
	var fg termbox.Attribute
	var bg termbox.Attribute

	row := make(Row, sizeX)

	for it, ir := 0, 0; it < len(h.Titles) && ir < sizeX; it++ {

		size := h.Titles[it].Per * sizeX / 100

		if h.IsActive && h.Active == it {
			fg = h.FgActive
			bg = h.BgActive
		} else {
			fg = h.Fg
			bg = h.Bg
		}

		name := h.Titles[it].Name

		for i := 0; i < size && ir < sizeX; i++ {
			cell := Cell{
				Fg: fg,
				Bg: bg,
			}

			if i < len(name) {
				cell.C = rune(name[i])
			} else {
				cell.C = ' '
			}
			row[ir] = cell

			ir++
		}
	}

	return row
}

// fillSpace replaces the automatic (0) percentages to fill the space
func (h *Header) fillSpace() {
	var num int
	var tot int

	for _, title := range h.Titles {
		tot += title.Per
		if title.Per == 0 {
			num += 1
		}
	}

	if num == 0 {
		return
	}

	num = (100 - tot) / num

	for it := 0; it < len(h.Titles); it++ {
		if h.Titles[it].Per == 0 {
			h.Titles[it].Per = num
		}
	}
}

type HeaderTitle struct {
	Name string
	// Percentage of size
	// When equal to 0, it takes the bigger space available
	Per int
}

type Body struct {
	Data   [][]string
	Active int

	Fg termbox.Attribute
	Bg termbox.Attribute

	// Colors for the active row
	FgActive termbox.Attribute
	BgActive termbox.Attribute
}

func (b Body) buildRows(h []HeaderTitle, sizeX, sizeY int) []Row {
	var fg termbox.Attribute
	var bg termbox.Attribute

	rows := make([]Row, sizeY)

	for ib := 0; ib < len(b.Data) && ib < sizeY; ib++ {

		d := b.Data[ib]

		row := make(Row, sizeX)

		if b.Active == ib {
			fg = b.FgActive
			bg = b.BgActive
		} else {
			fg = b.Fg
			bg = b.Bg
		}

		for it, ir := 0, 0; it < len(d) && ir < sizeX; it++ {

			size := h[it].Per * sizeX / 100

			name := d[it]

			for i := 0; i < size && ir < sizeX; i++ {
				cell := Cell{
					Fg: fg,
					Bg: bg,
				}

				if i < len(name) {
					cell.C = rune(name[i])
				} else {
					cell.C = ' '
				}
				row[ir] = cell

				ir++
			}
		}

		rows[ib] = row
	}

	return rows
}
