package structure

import "fmt"

type student struct {
	name string
	sub1 int
	sub2 int
	sub3 int
}

func StudentResult() {
	students := []student{
		{name: "JD", sub1: 85, sub2: 90, sub3: 78},
		{name: "Sid", sub1: 92, sub2: 88, sub3: 95},
		{name: "Mayur", sub1: 70, sub2: 75, sub3: 80},
	}
	highestPercentage := 0.0
	var topperName string
	for _, s := range students {
		total := s.sub1 + s.sub2 + s.sub3
		average := float64(total) / 3
		percentage := (float64(total) / 300) * 100
		if highestPercentage < percentage {
			highestPercentage = percentage
			topperName = s.name
		}

		fmt.Println("-------------")
		fmt.Println("Student Name:", s.name)
		fmt.Println("Subject 1:", s.sub1)
		fmt.Println("Subject 2:", s.sub2)
		fmt.Println("Subject 3:", s.sub3)
		fmt.Println("Total:", total)
		fmt.Printf("Average: %.2f\n", average)
		fmt.Printf("Percentage: %.2f%%\n", percentage)
	}
	fmt.Printf("Highest Percentage of %s is %.2f\n", topperName, highestPercentage)
}
