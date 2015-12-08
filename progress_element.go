package fort

//import "strings"
import "fmt"

type ProgressElement struct {
	Prefix  string
	Postfix string
	Total   uint64
	Current uint64
}

func (c *Container) AddProgressElement(id string, prefix, postfix string, total, current uint64) *ProgressElement {
	pe := ProgressElement{prefix, postfix, total, current}
	c.elements = append(c.elements, &pe)
	c.elementIds[id] = &pe
	return &pe
}

func (p *ProgressElement) renderFor(c *Container) {
	//widthProgressBar := c.TerminalSize.W - len(p.Prefix) - len(p.Postfix) - 2

	c.Output.Write([]byte(p.Prefix))

}

func progressAsBar(current, total uint64, width int) string {
	// |1111123333|
	delimitedWidth := uint64(width - 2)
	numberOfOnes := int(((current * delimitedWidth) / total))
	numberOfThrees := int(delimitedWidth) - numberOfOnes

	fmt.Println(numberOfOnes, numberOfThrees)
	return "asd"
	//	return "|" + strings.Repeat("=", numberOfOnes) + ">" + strings.Repeat("-", numberOfThrees) + "|"
}
