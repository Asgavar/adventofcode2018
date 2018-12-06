package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type GuardEventInfo struct {
	Timestamp  time.Time
	FellAsleep bool
	WokeUp     bool
	BeganShift bool
	GuardId    int
}

func sleepDuration(fellAsleepMinute, wokeUp int) int {
	return wokeUp - fellAsleepMinute
}

func minutesDuringWhichGuardWasSleepingTonight(fellAsleepMinute, wokeUp int) []int {
	var tonightMinutes []int

	for minute := fellAsleepMinute; minute < wokeUp; minute++ {
		tonightMinutes = append(tonightMinutes, minute)
	}

	return tonightMinutes
}

func sleepiestGuardId(minutesSpentAsleep map[int]int) int {
	maxGuardId := 0
	maxMinutes := 0

	for k, v := range minutesSpentAsleep {
		if v > maxMinutes {
			maxGuardId = k
			maxMinutes = v
		}
	}

	return maxGuardId
}

func main() {
	f, _ := os.Open("input")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	dateTimePattern := regexp.MustCompile("[0-9]{4}(-[0-9]{2}){2} [0-9]{2}:[0-9]{2}")
	var events []GuardEventInfo

	for scanner.Scan() {
		eventInfoAsString := scanner.Text()
		timestampAsString := dateTimePattern.FindString(eventInfoAsString)
		timestamp, _ := time.Parse("2006-01-02 15:04", timestampAsString)
		eventInfo := GuardEventInfo{Timestamp: timestamp}

		switch {
		case strings.Contains(eventInfoAsString, "falls asleep"):
			eventInfo.FellAsleep = true
		case strings.Contains(eventInfoAsString, "wakes up"):
			eventInfo.WokeUp = true
		case strings.Contains(eventInfoAsString, "begins shift"):
			guardIdAsString := strings.Split(eventInfoAsString, " ")[3][1:]
			guardId, _ := strconv.Atoi(guardIdAsString)
			eventInfo.BeganShift = true
			eventInfo.GuardId = guardId
		}

		events = append(events, eventInfo)
	}

	totalMinutesSpentSleeping := make(map[int]int)
	minutesWhenAsleep := make(map[int]map[int]int)

	var currentGuideId int
	var minuteFellAsleep int
	var minuteWokeUp int

	sort.Slice(
		events,
		func(a, b int) bool { return events[a].Timestamp.Before(events[b].Timestamp) })

	for _, event := range events {
		switch {
		case event.BeganShift:
			currentGuideId = event.GuardId
		case event.FellAsleep:
			minuteFellAsleep = event.Timestamp.Minute()
		case event.WokeUp:
			minuteWokeUp = event.Timestamp.Minute()
			totalMinutesSpentSleeping[currentGuideId] += sleepDuration(minuteFellAsleep, minuteWokeUp)
			if minutesWhenAsleep[currentGuideId] == nil {
				minutesWhenAsleep[currentGuideId] = make(map[int]int)
			}
			for _, minute := range minutesDuringWhichGuardWasSleepingTonight(minuteFellAsleep, minuteWokeUp) {
				minutesWhenAsleep[currentGuideId][minute] += 1
			}
		}
	}

	maxGuardId := sleepiestGuardId(totalMinutesSpentSleeping)
	maxMinuteForThisGuard := sleepiestGuardId(minutesWhenAsleep[maxGuardId])

	fmt.Println("[PART I] Guard ID:", maxGuardId)
	fmt.Println("[PART I] His sleepiest minute:", maxMinuteForThisGuard)

	sleepiestMinuteSleepCount := 0
	sleepiestMinute := 0
	sleepiestGuardId := 0

	for guardId, sleepPattern := range minutesWhenAsleep {
		for minute, dayCount := range sleepPattern {
			if dayCount > sleepiestMinuteSleepCount {
				sleepiestMinute = minute
				sleepiestMinuteSleepCount = dayCount
				sleepiestGuardId = guardId
			}
		}
	}

	fmt.Println("[PART II] Minute", sleepiestMinute, "and guard", sleepiestGuardId)
}
