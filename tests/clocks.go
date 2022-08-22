package test

import "fmt"

// Define the Clock type here.
type Clock struct {
	h int
	m int
}

func New(h, m int) Clock {

	var countToAddToHour int
	var countToSubToHour int

	if m < 0 {
		for m < -59 {
			m += 60
			countToSubToHour++
		}
		if m != 0 {
			countToSubToHour++
		}
		if m < 0 {
			m = 60 + m
		}
		if h > countToSubToHour {
			h -= countToSubToHour
		} else {
			h = (24 - countToSubToHour) + h
		}
	} else if m > 59 {
		for m > 59 {
			m -= 60
			countToAddToHour++
		}
		h += countToAddToHour
	}

	if h < 0 {
		h = 0 - h
		if h < 24 {
			h = 24 - h
		} else {
			h = 24 - (h % 24)
		}
	} else if h >= 24 {
		if h%24 == 0 {
			h = 0
		} else {
			h = h % 24
		}
	}

	instancePur := Clock{h: h, m: m}
	return instancePur
}

func (c Clock) Add(m int) Clock {
	return New(c.h, c.m+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(c.h, c.m-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}
