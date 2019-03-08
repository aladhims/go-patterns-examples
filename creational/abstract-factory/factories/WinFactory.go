package factories

import (
	"github.com/aladhims/go-patterns-examples/creational/abstract-factory/button"
	"github.com/aladhims/go-patterns-examples/creational/abstract-factory/checkbox"
)

type WinFactory struct{}

func (w *WinFactory) CreateButton() button.Buttoner {
	return &button.WinButton{}
}

func (w *WinFactory) CreateCheckbox() checkbox.Checkboxer {
	return &checkbox.WinCheckbox{}
}
