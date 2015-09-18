// Copyright 2015 Ghislain Guiot <gt.ghislain@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package termutil

import (
	"github.com/nsf/termbox-go"
)

type Cell struct {
	C  rune
	Fg termbox.Attribute
	Bg termbox.Attribute
}
