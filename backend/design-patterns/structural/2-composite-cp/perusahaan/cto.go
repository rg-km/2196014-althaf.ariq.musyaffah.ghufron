package perusahaan

type CTO struct {
	Subordinate []Employee
}

func (c CTO) GetSalary() int {
	return 30
}

func (c CTO) TotalDivisonSalary() int {
	total := 0
	for _, subordinate := range c.Subordinate {
		total += subordinate.TotalDivisonSalary()
	}
	totalDivisionSalary := c.GetSalary() + total
	return totalDivisionSalary

}
