package event

type event struct {
	name, routingKey, exchange string
}

func (e *event) Name() string {
	return e.name
}

func (e *event) RoutingKey() string {
	return e.name
}
func (e *event) Exchange() string {
	return e.name
}

func New(name, routingKey, exchange string) Event {
	return &event{name, routingKey, exchange}
}
