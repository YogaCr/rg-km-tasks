package laptop

// Ini adalah concrete implementation dari state On
type On struct {
	Laptop *Laptop
}

func (o On) Press() {
	o.Laptop.CurrentState = "Off"
	o.Laptop.ChangeState(&Off{o.Laptop})
}

func (o On) CanTurnOnLaptop() bool {
	return true
}

func (o On) Sleep() {
	o.Laptop.ChangeState(&Sleeping{o.Laptop})
	o.Laptop.ChangeCurrentState("Sleeping")
	// TODO: answer here
}
