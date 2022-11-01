package customer

type Observer interface {
	update(string)
	getID() string
}
