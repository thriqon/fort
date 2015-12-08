package fort

func (c *Container) csi(seq string) {
	c.Output.Write([]byte{0x1B, 0x5B})
	c.Output.Write([]byte(seq))
}
