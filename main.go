package main

import "fmt"

type Student struct {
	Name   string
	Grades []float64
}

func (student Student) AverageGrade() float64 {
	if len(student.Grades) == 0 { // Обрабатываю случай с пустым срезом избегая деления на ноль
		return 0.0
	}

	sum := 0.0
	for _, grade := range student.Grades {
		sum += grade
	}

	avg := sum / float64(len(student.Grades))
	return avg
}

func main() {
	student := Student{
		Name: "Jhon Doe",
		Grades: []float64{
			4.0,
			5.0,
			4.9,
			4.5,
		},
	}
	avg := student.AverageGrade()
	fmt.Printf("Средний бал студента %s: %.2f\n", student.Name, avg)
}
