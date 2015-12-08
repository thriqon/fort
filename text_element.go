package fort

// A TextElement is an element that shows a message.
// Make sure that the length of the text is not too long, as
// long messages may be shortened.
type TextElement struct {
	Message string
}

func (c *Container) AddTextElement(id string, msg string) *TextElement {
	te := TextElement{Message: msg}
	c.elements = append(c.elements, &te)
	c.elementIds[id] = &te
	te.renderFor(c)
	c.Output.Write([]byte("\n"))
	return &te
}

func (t *TextElement) renderFor(c *Container) {
	c.Output.Write([]byte(t.Message))
}
