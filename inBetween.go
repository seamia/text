// Copyright 2020 Seamia Corporation. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package text

import (
	"strings"
)

func FindInBetween(whole, start, end string) (string, bool) {

	if a := strings.Index(whole, start); a >= 0 {
		remains := whole[a+len(start):]
		if b := strings.Index(remains, end); a >= 0 {
			return remains[:b], true
		}
	}
	return "", false
}

