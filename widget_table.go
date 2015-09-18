// Copyright 2015 Ghislain Guiot <gt.ghislain@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package termutil

import (
	"github.com/nsf/termbox-go"
)

type WidgetTable struct {
	Header *Header

	Body [][]string

	Active int
}

type Header struct {
	Titles []HeaderTitle

	Fg termbox.Attribute
	Bg termbox.Attribute

	// Colors for the active tab
	FgActive termbox.Attribute
	BgActive termbox.Attribute
}

type HeaderTitle struct {
	Name string
	// Percentage of size
	// When equal to 0, it takes the bigger space available
	Per int
}

func (h Header) buildRow(winX int, active int) Row {
	row := make(Row, winX)
	var fg termbox.Attribute
	var bg termbox.Attribute

	for it, ir := 0, 0; it < len(h.Titles) && ir < winX; it++ {

		size := h.Titles[it].Per * winX / 100

		if active == it {
			fg = h.FgActive
			bg = h.BgActive
		} else {
			fg = h.Fg
			bg = h.Bg
		}

		name := h.Titles[it].Name

		for i := 0; i < size && ir < winX; i++ {
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

func (wg *WidgetTable) Update(win *Window) []Row {
	out := make([]Row, len(wg.Body)+1)

	wg.Header.fillSpace()

	out[0] = wg.Header.buildRow(win.SizeX, wg.Active)

	for i, b := range wg.Body {
		out[i+1] = buildRow(b, wg.Header, win.SizeX)
	}

	return out
}

func buildRow(row []string, his *Header, winX int) Row {
	rowf := make(Row, winX)
	var ir int

	for is, r := range row {

		if ir >= winX {
			break
		}

		size := his.Titles[is].Per * winX / 100

		for i := 0; i < size && ir < winX; i++ {
			if i < len(r) {
				rowf[ir].C = rune(r[i])
			} else {
				rowf[ir].C = ' '
			}
			ir++
		}
	}

	return rowf
}
