package main

import "github.com/TheMarstonConnell/webframe/frame"

func main() {
	root := frame.Init()
	root.AppendToBody(root.P("paragraph", "", "this is a paragraph!"))
	b := root.Button("click", "", "Click Me!")
	frame.SetVariant(b, frame.VARIANT_SUCCESS)
	root.AppendToBody(b)
}
