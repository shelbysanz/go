package main

type notification interface {
	importance() int
}

type directMessage struct {
	senderUsername string
	messageContent string
	priorityLevel  int
	isUrgent       bool
}

type groupMessage struct {
	groupName      string
	messageContent string
	priorityLevel  int
}

type systemAlert struct {
	alertCode      string
	messageContent string
}

func (d directMessage) importance() int {
	if d.isUrgent {
		return 50
	} else {
		return d.priorityLevel
	}
}

func (g groupMessage) importance() int {
	return g.priorityLevel
}

func (s systemAlert) importance() int {
	return 100
}

func processNotification(n notification) (string, int) {
	switch value := n.(type) {
	case directMessage:
		return value.senderUsername, value.importance()
	case groupMessage:
		return value.groupName, value.importance()
	case systemAlert:
		return value.alertCode, value.importance()
	default:
		return "", 0
	}
}
