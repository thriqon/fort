package fort

import "strconv"
import "os"
import "golang.org/x/crypto/ssh/terminal"

// Container is the main element of Fort.
// It contains all displayed elements, and
// is the only instance allowed to modify this list.
type Container struct {
	elements     []Element
	elementIds   map[string]Element
	Output       *os.File
	TerminalSize struct {
		W, H int
	}
}

func NewContainer() (*Container, error) {
	c := Container{}
	c.Output = os.Stdout
	c.elementIds = make(map[string]Element)

	w, h, err := terminal.GetSize(int(c.Output.Fd()))
	if err != nil {
		return nil, err
	}
	c.TerminalSize.W = w
	c.TerminalSize.H = h

	return &c, nil
}

// An Element is shown on the screen.
// Each element occupies exactly one line of the screen.
type Element interface {
	renderFor(*Container)
}

func (c *Container) Refresh() {
	// move upwards
	c.csi(strconv.Itoa(len(c.elements)-1) + "F")
	for _, el := range c.elements {
		el.renderFor(c)
		// clear till end of line
		c.csi("K")
		// one line down, beginning of line
		c.csi("E")
	}
}
