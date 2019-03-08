package factories

import (
	"github.com/aladhims/go-patterns-examples/creational/abstract-factory/button"
	"github.com/aladhims/go-patterns-examples/creational/abstract-factory/checkbox"
)

type GUIFactory interface {
	CreateButton() button.Buttoner
	CreateCheckbox() checkbox.Checkboxer
}
