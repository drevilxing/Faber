package client

const EventServiceTypeDeliver = "deliver"
const EventServiceTypeEventHub = "eventhub"

type EventServiceTimeout struct {
	Connection           string `json:"connection"`
	RegistrationResponse string `json:"registrationResponse"`
}

type EventService struct {
	Method  string               `json:"type"`
	Timeout *EventServiceTimeout `json:"timeout"`
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
