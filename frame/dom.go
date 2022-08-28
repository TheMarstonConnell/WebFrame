package frame

import (
	"fmt"

	"honnef.co/go/js/dom/v2"
)

func (b Builder) Build(objType string, id string, classList string) dom.Element {
	p := b.Doc.CreateElement(objType)
	p.SetID(id)

	return p
}

func (b Builder) P(id string, classList string, text string) dom.Element {
	p := b.Build("p", id, classList)
	p.SetInnerHTML(text)
	return p
}

func (b Builder) Button(id string, classList string, text string) dom.Element {
	p := b.Build("sl-button", id, classList)
	p.SetInnerHTML(text)
	return p
}

func SetVariant(d dom.Element, variant string) {
	d.SetAttribute("variant", variant)
}

func (b Builder) AppendToBody(d dom.Element) {
	b.Body.AppendChild(d)
}

func Init() Builder {
	fmt.Println("Starting Dom")

	d := dom.GetWindow().Document()

	docbody := d.GetElementByID("docbody")

	b := Builder{
		Doc:  d,
		Body: docbody,
	}

	return b

}
