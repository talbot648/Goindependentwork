package main

import (
	"fmt"
)

type Student struct {
	name   string
	grades []float32
	age    int
}

func (firstStudent Student) getAverageGrade() float32 {
	var totalMarks float32
	totalMarks = 0
	for _, marks := range firstStudent.grades {
		totalMarks += marks
	}
	return totalMarks / float32(len(firstStudent.grades))
}

func main() {
	firstStudent := Student{"Charlie", []float32{70, 60, 80, 90, 50, 80}, 19}
	averageGrade := firstStudent.getAverageGrade()
	fmt.Println(averageGrade)
}
