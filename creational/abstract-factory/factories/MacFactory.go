package factories

import (
	"github.com/aladhims/go-patterns-examples/creational/abstract-factory/button"
	"github.com/aladhims/go-patterns-examples/creational/abstract-factory/checkbox"
)

type MacFactory struct{}

func (m *MacFactory) CreateButton() button.Buttoner {
	return &button.MacButton{}
}

func (m *MacFactory) CreateCheckbox() checkbox.Checkboxer {
	return &checkbox.MacCheckbox{}
}
