package algorithm

func BinaryAllConstraints(a *Algorithm) {
	noThreeInARowConstraint(a)
	uniqueConstraint(a)
	equalSumConstraint(a)
}

func noThreeInARowConstraint(a *Algorithm) {
	addCheck := func(f1, f2, f3 *Field) {
		constraint := func() bool {
			if f1.val != -1 && f1.val == f2.val && f1.val == f3.val {
				return false
			}
			return true
		}
		f1.Constraints = append(f1.Constraints, constraint)
		f2.Constraints = append(f2.Constraints, constraint)
		f3.Constraints = append(f3.Constraints, constraint)
	}

	for i := 0; i < a.edgeSize; i++ {
		for j := 0; j < a.edgeSize; j++ {
			if i <= a.edgeSize-3 {
				defer addCheck(a.fields[i*a.edgeSize+j], a.fields[(i+1)*a.edgeSize+j], a.fields[(i+2)*a.edgeSize+j])
			}
			if j <= a.edgeSize-3 {
				defer addCheck(a.fields[i*a.edgeSize+j], a.fields[i*a.edgeSize+j+1], a.fields[i*a.edgeSize+j+2])
			}
		}
	}
}

func fieldsFromColumn(a *Algorithm, col int) []*Field {
	values := make([]*Field, a.edgeSize)
	for i := 0; i < a.edgeSize; i++ {
		values[i] = a.fields[col*a.edgeSize+i]
	}
	return values
}

func fieldsFromRow(a *Algorithm, row int) []*Field {
	values := make([]*Field, a.edgeSize)
	for i := 0; i < a.edgeSize; i++ {
		values[i] = a.fields[i*a.edgeSize+row]
	}
	return values
}

func uniqueConstraint(a *Algorithm) {
	slicesEqual := func(s1, s2 []*Field) bool {
		for i, v := range s1 {
			if v.val == -1 || s2[i].val == -1 {
				return false
			}
			if v.val != s2[i].val {
				return false
			}
		}
		return true
	}

	addCheck := func(f1, f2 []*Field) {
		constraint := func() bool {
			return !slicesEqual(f1, f2)
		}

		for _, f := range f1 {
			f.Constraints = append(f.Constraints, constraint)
		}

		for _, f := range f2 {
			f.Constraints = append(f.Constraints, constraint)
		}
	}

	for i := 0; i < a.edgeSize; i++ {
		row := fieldsFromRow(a, i)
		column := fieldsFromColumn(a, i)
		for j := i + 1; j < a.edgeSize; j++ {
			defer addCheck(row, fieldsFromRow(a, j))
			defer addCheck(column, fieldsFromColumn(a, j))
		}
	}
}

func equalSumConstraint(a *Algorithm) {
	equalSum := func(f1, f2 []*Field) bool {
		var s1, s2 int
		for i := 0; i < len(f1); i++ {
			if f1[i].val == -1 || f2[i].val == -1 {
				return true
			}
			s1 += f1[i].val
			s2 += f2[i].val
		}
		return s1 == s2
	}

	addCheck := func(f1, f2 []*Field) {
		constraint := func() bool {
			return equalSum(f1, f2)
		}

		for _, f := range f1 {
			f.Constraints = append(f.Constraints, constraint)
		}

		for _, f := range f2 {
			f.Constraints = append(f.Constraints, constraint)
		}
	}

	for i := 0; i < a.edgeSize; i++ {
		row := fieldsFromRow(a, i)
		column := fieldsFromColumn(a, i)
		for j := 0; j < a.edgeSize; j++ {
			if i != j {
				defer addCheck(row, fieldsFromRow(a, j))
				defer addCheck(column, fieldsFromColumn(a, j))
			}
		}
	}
}
