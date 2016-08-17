package main

import (
	"errors"
	"regexp"
	"strconv"
)

var (
	sizeReg = regexp.MustCompile("[0-9]{1,}[TtGgMmKkBb]?")
)

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

func parseSize(size string) (int64, error) {
	if !sizeReg.MatchString(size) {
		return 0, errors.New("invalid size:" + size)
	}
	return sizeToBytes(size), nil
}

func sizeToBytes(s string) int64 {
	switch l, c := len(s), s[len(s)-1]; c {
	case 'T', 't':
		i, err := strconv.ParseInt(s[:l-1], 10, 64)
		if err != nil {
			panic(err)
		}
		return i << 40
	case 'G', 'g':
		i, err := strconv.ParseInt(s[:l-1], 10, 64)
		if err != nil {
			panic(err)
		}
		return i << 30
	case 'M', 'm':
		i, err := strconv.ParseInt(s[:l-1], 10, 64)
		if err != nil {
			panic(err)
		}
		return i << 20
	case 'K', 'k':
		i, err := strconv.ParseInt(s[:l-1], 10, 64)
		if err != nil {
			panic(err)
		}
		return i << 10
	case 'B', 'b':
		i, err := strconv.ParseInt(s[:l-1], 10, 64)
		if err != nil {
			panic(err)
		}
		return i
	}
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}
