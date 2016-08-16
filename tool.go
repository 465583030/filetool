package main

import "strconv"

func readableBytes(size int64) string {
	switch {
	case size>>40 > 0:
		return strconv.FormatInt(size>>40, 10) + "T"
	case size>>30 > 0:
		return strconv.FormatInt(size>>30, 10) + "G"
	case size>>20 > 0:
		return strconv.FormatInt(size>>20, 10) + "M"
	case size>>10 > 0:
		return strconv.FormatInt(size>>10, 10) + "K"
	}
	return strconv.FormatInt(size, 10) + "B"
}
