package main

import (
	"runtime"

	"github.com/aladhims/go-patterns-examples/creational/abstract-factory"
	"github.com/aladhims/go-patterns-examples/creational/abstract-factory/factories"
)

var AFactory *abstract.Factory
var GUIFactory factories.GUIFactory

func init() {
	OS := runtime.GOOS

	switch OS {
	case "windows":
		GUIFactory = &factories.WinFactory{}
		break
	case "darwin":
		GUIFactory = &factories.MacFactory{}
		break
	default:
		panic("Unsupported OS")
	}

	AFactory = abstract.NewFactory(GUIFactory)
}

func main() {
	AFactory.BPaint()
	AFactory.CPaint()
}
