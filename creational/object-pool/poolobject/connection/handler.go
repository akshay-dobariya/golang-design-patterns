package connection

type Connection struct {
	id int
}

func (c *Connection) GetId() int {
	return c.id
}

func (c *Connection) SetId(id int) {
	c.id = id
}
