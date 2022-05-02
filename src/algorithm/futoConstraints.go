package algorithm

func FutoAllConstraints(a *Algorithm, ineq []inequality) {
	sequenceConstraint(a)
	inequalitiesConstraint(a, ineq)
}

func sequenceConstraint(a *Algorithm) {
	check := func(f1 []*Field) bool {
		done := make([]int, a.edgeSize+1)
		for i := 0; i < a.edgeSize; i++ {
			if f1[i].val != -1 {
				if done[f1[i].val] != 0 {
					return false
				}
				done[f1[i].val] = 1
			}
		}
		return true
	}

	addCheck := func(f1 []*Field) {
		constraint := func() bool {
			return check(f1)
		}

		for _, f := range f1 {
			f.Constraints = append(f.Constraints, constraint)
		}
	}

	for i := 0; i < a.edgeSize; i++ {
		defer addCheck(fieldsFromRow(a, i))
		defer addCheck(fieldsFromColumn(a, i))
	}
}

func inequalitiesConstraint(a *Algorithm, inequalities []inequality) {
	check := func(f1, f2 *Field) bool {
		if !(f1.val > f2.val) &&
			f1.val != -1 &&
			f2.val != -1 {
			return false
		}
		return true
	}

	addCheck := func(f1, f2 *Field) {
		constraint := func() bool {
			return check(f1, f2)
		}
		f1.Constraints = append(f1.Constraints, constraint)
		f2.Constraints = append(f2.Constraints, constraint)
	}

	for _, ineq := range inequalities {
		defer addCheck(a.fields[ineq.bigger], a.fields[ineq.lower])
	}
}
