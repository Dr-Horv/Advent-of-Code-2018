package day04

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type entry struct {
	Timestamp time.Time
	Info      string
}

func (e entry) String() string {
	return fmt.Sprintf("%v %v", e.Timestamp, e.Info)
}

var entryFormat = regexp.MustCompile(`(?m)\[(.*)](.*)`)

func Solve(lines []string, partOne bool) string {
	var entries []entry

	for _, line := range lines {
		entries = append(entries, parseEntry(line))
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Timestamp.Before(entries[j].Timestamp)
	})

	guards := make(map[string][60]int)
	var startOfSleep time.Time
	var guardId = ""

	for _, entry := range entries {
		switch {
		case strings.Contains(entry.Info, "begins shift"):
			parts := strings.Split(entry.Info, " ")
			guardId = parts[1]
			_, found := guards[guardId]
			if !found {
				var minutes [60]int
				guards[guardId] = minutes
			}
		case strings.Contains(entry.Info, "falls asleep"):
			startOfSleep = entry.Timestamp
		case strings.Contains(entry.Info, "wakes up"):
			minutes := guards[guardId]
			for m := startOfSleep.Minute(); m < entry.Timestamp.Minute(); m++ {
				minutes[m] = minutes[m] + 1
			}
			guards[guardId] = minutes
		}
	}

	var sleepiestGuard = ""
	var minute = -1
	if partOne {
		var minutesSlept = -1
		for k, v := range guards {
			var sum = 0
			for _, m := range v {
				sum += m
			}
			if sum > minutesSlept {
				sleepiestGuard = k
				minutesSlept = sum
			}
		}

		var max = 0
		for i, m := range guards[sleepiestGuard] {
			if m > max {
				minute = i
				max = m
			}
		}
	} else {
		var max = -1
		for m := 0; m < 60; m++ {
			for k, v := range guards {
				if v[m] > max {
					max = v[m]
					sleepiestGuard = k
					minute = m
				}
			}
		}
	}

	guardIdNumber, _ := strconv.Atoi(sleepiestGuard[1:])

	answer := guardIdNumber * minute

	return fmt.Sprint(answer)
}

func parseEntry(s string) entry {
	groups := entryFormat.FindStringSubmatch(s)
	const form = "2006-01-02 15:04"
	t, _ := time.Parse(form, groups[1])
	return entry{t, strings.TrimSpace(groups[2])}

}
