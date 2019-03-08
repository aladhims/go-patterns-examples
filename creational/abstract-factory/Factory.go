package abstract

import (
	"github.com/aladhims/go-patterns-examples/creational/abstract-factory/button"
	"github.com/aladhims/go-patterns-examples/creational/abstract-factory/checkbox"
	"github.com/aladhims/go-patterns-examples/creational/abstract-factory/factories"
)

type Factory struct {
	button.Buttoner
	checkbox.Checkboxer
}

func NewFactory(factory factories.GUIFactory) *Factory {
	return &Factory{
		Buttoner:   factory.CreateButton(),
		Checkboxer: factory.CreateCheckbox(),
	}
}
