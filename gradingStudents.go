func gradingStudents(grades []int32) []int32 {
    finalGrades := []int32{}
    for _, grade := range grades {
        if grade >= 38 {
            round := ((grade/5) * 5 + 5)
            diff := round - grade 
            if diff < 3 {
                finalGrades = append(finalGrades, round)
            } else {
                finalGrades = append(finalGrades, grade)
            }
        } else {
            finalGrades = append(finalGrades, grade)
        }
    }
    return finalGrades
}
