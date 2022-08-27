// Copyright 2020 Seamia Corporation. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package text

import (
	"strings"
)

/**
	Finds first occurence of text between specified markers, e.g.
	whole:		"prefix [ center ] suffix""
	start:		"["
	end:		"]"
	returns:	" center ", true
 */
func FindInBetween(whole, start, end string) (string, bool) {

	if left := strings.Index(whole, start); left >= 0 {
		remains := whole[left+len(start):]
		if right := strings.Index(remains, end); right >= 0 {
			return remains[:right], true
		}
	}
	return "", false
}

func RemoveInBetween(whole, start, end string, includingMarkers bool) string {

	if left := strings.Index(whole, start); left >= 0 {
		remains := whole[left+len(start):]
		if right := strings.Index(remains, end); right >= 0 {
			if includingMarkers {
				return whole[:left] + whole[left+len(start)+right+len(end):]
			} else {
				return whole[:left+len(start)] + whole[left+len(start)+right:]
			}
		}
	}
	return whole
}
