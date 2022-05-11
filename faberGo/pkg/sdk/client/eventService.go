package client

const EventServiceTypeDeliver = "deliver"
const EventServiceTypeEventHub = "eventhub"

type EventServiceTimeout struct {
	Connection           string `json:"connection" yaml:"connection"`
	RegistrationResponse string `json:"registrationResponse" yaml:"registrationResponse"`
}

type EventService struct {
	Method  string               `json:"type" yaml:"method"`
	Timeout *EventServiceTimeout `json:"timeout" yaml:"timeout"`
}

func GenerateDefaultEventService(method string) *EventService {
	return &EventService{
		Method: method,
		Timeout: &EventServiceTimeout{
			Connection:           "15s",
			RegistrationResponse: "15s",
		},
	}
}
