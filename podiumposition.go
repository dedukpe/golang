package piscine

func PodiumPosition(podium [][]string) [][]string {
	if len(podium) == 4 {
		return [][]string{podium[3], podium[2], podium[1], podium[0]}
	}
	if len(podium) == 3 {
		return [][]string{podium[2], podium[1], podium[0]}
	}
	if len(podium) == 2 {
		return [][]string{podium[1], podium[0]}
	}
	if len(podium) == 1 {
		return [][]string{podium[0]}
	}
	return [][]string{}
}
