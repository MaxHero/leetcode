package _0732

import "sort"

type Period struct {
	start int32
	k     int16
}

type MyCalendarThree struct {
	periods []Period
	kMax    int16
}

func max(a, b int16) int16 {
	if a >= b {
		return a
	}
	return b
}

func Constructor() MyCalendarThree {
	periods := make([]Period, 2, 802)
	periods[1].start = 1000000000
	return MyCalendarThree{periods: periods}
}

func (m *MyCalendarThree) Book(startInt int, endInt int) int {
	start := int32(startInt)
	end := int32(endInt)

	idx := sort.Search(len(m.periods), func(i int) bool {
		return start < m.periods[i].start
	})

	if m.periods[idx-1].start == start {
		idx--
	} else {
		m.periods = append(m.periods, Period{})
		copy(m.periods[idx:], m.periods[idx-1:])
		m.periods[idx].start = start
	}

	for ; m.periods[idx+1].start < end; idx++ {
		m.periods[idx].k++
		m.kMax = max(m.kMax, m.periods[idx].k)
	}

	if m.periods[idx].start < end {
		m.periods = append(m.periods, Period{})
		copy(m.periods[idx+1:], m.periods[idx:])
		m.periods[idx+1].start = end
		m.periods[idx].k++
		m.kMax = max(m.kMax, m.periods[idx].k)
	}

	return int(m.kMax)
}
