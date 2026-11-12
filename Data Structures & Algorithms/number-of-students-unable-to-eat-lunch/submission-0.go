func countStudents(students []int, sandwiches []int) int {
    for {
		turn := len(students)
		numStudents := len(students)
		for turn > 0 && len(students) > 0 {
			student := students[0]
			sandwich := sandwiches[0]
			if student == sandwich{
				students = students[1:]
				sandwiches = sandwiches[1:]
			} else {
				students = students[1:]
				students = append(students, student)
			}
			turn--
		}
		if numStudents == len(students){
			return numStudents
		}
	}
	return 0
}