package day4

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

type Guard struct {
	ID      int
	Entries []Entry
}

func (g *Guard) getTotalTimeSlept() int {
	var total time.Duration

	for _, entry := range g.Entries {
		total += entry.getDuration()
		fmt.Printf("Slept %d minutes\n", int(entry.getDuration().Minutes()))
	}

	fmt.Printf("Total Slept: %d minutes\n", int(total.Minutes()))

	return int(total.Minutes())
}

func (g *Guard) getMostSleptMinute() int {
	var minutes [60]int

	for _, entry := range g.Entries {
		for i := entry.Start.Minute(); i < entry.End.Minute(); i++ {
			minutes[i]++
		}
	}

	var greatest int
	for minute, total := range minutes {
		if total > minutes[greatest] {
			greatest = minute
		}
	}

	return greatest
}

func (g *Guard) getMostSleptMinuteTimes() int {
	var minutes [60]int

	for _, entry := range g.Entries {
		for i := entry.Start.Minute(); i < entry.End.Minute(); i++ {
			minutes[i]++
		}
	}

	var greatest int
	for minute, total := range minutes {
		if total > minutes[greatest] {
			greatest = minute
		}
	}

	return minutes[greatest]
}

type Entry struct {
	Start time.Time
	End   time.Time
}

func (e *Entry) getDuration() time.Duration {
	return e.End.Sub(e.Start) - time.Minute
}

type Result struct {
	Guard Guard
	Minute int
}

var DateRegex = regexp.MustCompile("^\\[(.*?)]")
var GuardRegex = regexp.MustCompile("Guard #(\\d+) begins shift$")

func Strategy1(entries []string) Result {
	guards := LoadGuards(entries)
	fmt.Println()

	fmt.Println("CALCULATING MOST ASLEEP GUARD")
	fmt.Println("-----------------------------")
	var mostSleepy int
	var maxSlept int
	for id, guard := range guards {
		timeSlept := guard.getTotalTimeSlept()

		fmt.Printf("Guard #%d sleeps %d minutes\n", id, timeSlept)
		fmt.Println()

		if maxSlept < timeSlept {
			maxSlept = timeSlept
			mostSleepy = id
		}
	}

	fmt.Printf("Guard Most Asleep: #%d\n", mostSleepy)
	fmt.Printf("Most Likely Asleep: %d\n", guards[mostSleepy].getMostSleptMinute())

	fmt.Printf("Result: %d\n", mostSleepy*guards[mostSleepy].getMostSleptMinute())

	return Result{
		Guard:  *guards[mostSleepy],
		Minute: guards[mostSleepy].getMostSleptMinute(),
	}
}

func Strategy2(entries []string) Result {
	guards := LoadGuards(entries)
	fmt.Println()

	fmt.Println("CALCULATING GUARD MOST CONSISTENTLY ASLEEP")
	fmt.Println("------------------------------------------")
	var guardId int
	var minute int
	var amount int
	for id, guard := range guards {
		m := guard.getMostSleptMinute()
		a := guard.getMostSleptMinuteTimes()


		fmt.Printf("Guard #%d sleeps most often on %d (%d times)\n", id, m, a)
		fmt.Println()

		if amount < a {
			amount = a
			minute = m
			guardId = id
		}
	}

	fmt.Printf("Guard Most Asleep: #%d\n", guardId)
	fmt.Printf("Most Likely Asleep: %d\n", minute)

	fmt.Printf("Result: %d\n", guardId * minute)

	return Result{
		Guard:  *guards[guardId],
		Minute: guards[guardId].getMostSleptMinute(),
	}
}

func LoadGuards(entries []string) map[int]*Guard {
	var guards = make(map[int]*Guard)
	var guardId int

	fmt.Println("LOADING GUARDS DATA")
	fmt.Println("-------------------")
	for i := 0; i < len(entries); i++ {
		date := parseTime(entries[i])
		fmt.Printf("Date: %s\n", date.Format(time.RFC3339))

		if GuardRegex.MatchString(entries[i]) {
			guard := parseGuard(entries[i])
			guardId = guard.ID

			if guards[guardId] == nil {
				guards[guardId] = &guard
			}

			fmt.Printf("Current Guard: %d\n", guardId)

			i++
		}

		if !GuardRegex.MatchString(entries[i]) && !GuardRegex.MatchString(entries[i]) {
			fmt.Println(entries[i])
			fmt.Println(entries[i+1])

			guards[guardId].Entries = append(guards[guardId].Entries, Entry{
				Start: parseTime(entries[i]),
				End:   parseTime(entries[i+1]),
			})

			i += 1
		}

		fmt.Println()
	}

	return guards
}

func parseTime(entry string) time.Time {
	matches := DateRegex.FindStringSubmatch(entry)
	date, err := time.Parse("2006-01-02 15:04", matches[1])
	if err != nil {
		panic(err)
	}

	return date
}

func parseGuard(entry string) Guard {
	matches := GuardRegex.FindStringSubmatch(entry)
	id, err := strconv.Atoi(matches[1])
	if err != nil {
		panic(err)
	}

	return Guard{
		ID: id,
	}
}
