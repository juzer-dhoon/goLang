package mapping

import (
	"fmt"
)

func CalculateResult() {
	name := "Juzer Dhoon"

	subjectMarks := map[string]float64{
		"Math":      85,
		"Science":   90,
		"English":   78,
		"History":   88,
		"Geography": 92,
	}

	var total float64
	subjectCount := len(subjectMarks)

	// Calculate total marks
	for _, marks := range subjectMarks {
		total += marks
	}

	// Calculate average and percentage
	average := total / float64(subjectCount)
	percentage := (total / float64(subjectCount*100)) * 100

	// Print result
	fmt.Println("--- Result ---")
	fmt.Println("Name:", name)
	fmt.Println("Total Subjects:", subjectCount)

	for subject, marks := range subjectMarks {
		fmt.Printf("%s: %.2f\n", subject, marks)
	}
	fmt.Println("---------------")
	fmt.Printf("Average Marks: %.2f\n", average)
	fmt.Printf("Percentage: %.2f%%\n", percentage)
}
