package main

import "fmt"

func main() {
	//hitung rata-rata
	scores := [8]float64{100, 80, 75, 92, 70, 93, 88, 67}
	res := 0.0
	for i := 0; i < len(scores); i++ {
		res += scores[i]
	}
	fmt.Println(res / float64(len(scores)))
	var goodScores []float64
	for i := 0; i < len(scores); i++ {
		if scores[i] >= 90 {
			goodScores = append(goodScores, scores[i])
		}
	}
	fmt.Println(goodScores)
}
