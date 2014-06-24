package mclass

type Person struct {
	Name string
}

func (c Person) GetName() string {
	return c.Name
}
