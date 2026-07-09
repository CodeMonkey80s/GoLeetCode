package solution3921

func scoreValidator(events []string) []int {
	var score int
	var counter int
	for i := 0; i < len(events); i++ {
		switch {
		case events[i] >= "0" && events[i] <= "6":
			score += int(events[i][0] - '0')
		case events[i] == "W":
			counter++
		case events[i] == "WD" || events[i] == "NB":
			score++
		}
		if counter == 10 {
			break
		}
	}

	return []int{score, counter}
}
