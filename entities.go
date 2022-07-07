package main

type dog struct {
	weight float64
}

func (d dog) String() string {
	return "dog"
}

func (d dog) Weight() float64 {
	return d.weight
}

func (d dog) calculateFood() float64 {
	const (
		weightStep        float64 = 5
		foodForWeightStep float64 = 10
	)

	if d.weight <= 0 {
		return 0
	}

	foodNeededForMonth := d.weight / weightStep * foodForWeightStep
	return foodNeededForMonth
}

type cat struct {
	weight float64
}

func (c cat) String() string {
	return "cat"
}

func (c cat) Weight() float64 {
	return c.weight
}

func (c cat) calculateFood() float64 {
	const (
		foodFor1KgWeight float64 = 7
	)

	if c.weight <= 0 {
		return 0
	}

	foodNeededForMonth := c.weight * foodFor1KgWeight
	return foodNeededForMonth
}

type cow struct {
	weight float64
}

func (c cow) String() string {
	return "cow"
}

func (c cow) Weight() float64 {
	return c.weight
}

func (c cow) calculateFood() float64 {
	const foodFor1KgWeight float64 = 25

	if c.weight <= 0 {
		return 0
	}

	foodNeededForMonth := c.weight * foodFor1KgWeight
	return foodNeededForMonth
}
