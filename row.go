// Copyright 2015 Ghislain Guiot <gt.ghislain@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package termutil

import (
	"github.com/nsf/termbox-go"
)

type Row []Cell

func StringToRow(s string, fg, bg termbox.Attribute) Row {
	r := make(Row, len(s))

	for i, c := range s {
		r[i] = Cell{c, fg, bg}
	}

	return r
}
func StringsToRows(s []string, fg, bg termbox.Attribute) []Row {
	rs := make([]Row, len(s))

	for i, s := range s {
		rs[i] = StringToRow(s, fg, bg)
	}

	return rs
}
