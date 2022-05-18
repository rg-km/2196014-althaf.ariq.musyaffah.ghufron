package perusahaan

type VP struct {
	Subordinate []Employee
}

func (vp VP) GetSalary() int {
	return 20
}

func (vp VP) TotalDivisonSalary() int {
	total := 0
	for _, subordinate := range vp.Subordinate {
		total += subordinate.GetSalary()
	}
	totalDivisionSalary := vp.GetSalary() + total
	return totalDivisionSalary
}
