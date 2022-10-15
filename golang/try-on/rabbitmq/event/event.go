package event

type Events []Event

type Event interface {
	Name() string
	RoutingKey() string
	Exchange() string
}

type BodyMessage map[string]string
