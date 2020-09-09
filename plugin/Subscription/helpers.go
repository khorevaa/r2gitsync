package Subscription

import "strings"

func getEvent(end string) eventType {

	if strings.HasPrefix(end, string(OnEvent)) {
		return OnEvent
	}

	if strings.HasPrefix(end, string(AfterEvent)) {
		return AfterEvent
	}

	if strings.HasPrefix(end, string(BeforeEvent)) {
		return BeforeEvent
	}

	return UnknownEvent

}
