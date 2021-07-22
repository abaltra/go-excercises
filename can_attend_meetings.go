package main

/*
	https://ttzztt.gitbooks.io/lc/content/sort/meeting-rooms.html

	Given an array of meeting time intervals consisting of start and end times[[s1,e1],[s2,e2],...](si< ei), determine if a person could attend all meetings.
*/

import (
	"sort"
)

func canAttendMeetings(intervals [][]int) bool {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for idx := 1; idx < len(intervals); idx++ {
		if intervals[idx-1][1] > intervals[idx][0] {
			return false
		}
	}

	return true
}
