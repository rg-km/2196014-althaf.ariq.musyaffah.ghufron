package laptop

type Sleeping struct {
	Laptop *Laptop
}

func (s Sleeping) Press() {
	// TODO: answer here
	s.Laptop.CurrentState = "On"
	s.Laptop.ChangeState(On{s.Laptop})
}

func (s Sleeping) CanTurnOnLaptop() bool {
	return true
}

func (s Sleeping) Sleep() {
	// TODO: answer here
	s.Laptop.CurrentState = "Sleeping"
	s.Laptop.ChangeState(Sleeping{s.Laptop})
}
