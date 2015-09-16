// Copyright 2015 Ghislain Guiot <gt.ghislain@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package termutil

type WidgetTable struct {
	Header []HeaderInfo
	Body   [][]string
}

type HeaderInfo struct {
	Title string
	// Percentage of size
	// When equal to 0, it takes the bigger space available
	Per int
}

func (wg *WidgetTable) Update(win *Window) []string {
	var out []string

	winX := win.SizeX
	per := calcPerc(wg.Header)

	head := make([]string, len(wg.Header))
	for i, h := range wg.Header {
		head[i] = h.Title
	}
	out = append(out, buildRow(head, wg.Header, per, winX))

	for _, b := range wg.Body {
		out = append(out, buildRow(b, wg.Header, per, winX))
	}

	return out
}

func buildRow(row []string, his []HeaderInfo, per int, winX int) string {
	var rowf string

	for is, r := range row {

		if his[is].Per == 0 {
			his[is].Per = per
		}
		size := his[is].Per * winX / 100

		for i := 0; i < size; i++ {
			if i < len(r) {
				rowf += string(r[i])
			} else {
				rowf += " "
			}
		}
	}

	return rowf
}

func calcPerc(his []HeaderInfo) int {
	var num int
	var tot int

	for _, hi := range his {
		tot += hi.Per
		if hi.Per == 0 {
			num += 1
		}
	}

	if num == 0 {
		return num
	}

	num = (100 - tot) / num
	return num
}
