package common

type WithNameField struct {
	name string
}

func (n *WithNameField) Name() string {
	return n.name
}

func (n *WithNameField) SetName(name string) {
	n.name = name
}

type Namer interface {
	Name() string
	SetName(name string)
}
