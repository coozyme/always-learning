package customer

import "fmt"

type Customer struct {
	Id string
}

func (c *Customer) update(itemName string) {
	fmt.Printf("Sending email to customer %s for item %s\n", c.Id, itemName)
}

func (c *Customer) getID() string {
	return c.Id
}
