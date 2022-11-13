package notifications

import (
	"bytes"
	"fmt"
	"strings"
)

type Notifications struct {
	errors map[string][]string
}

func NewNotification() *Notifications {
	return &Notifications{make(map[string][]string, 10)}
}

func (n *Notifications) AddError(context string, msg string) {
	n.errors[context] = append(n.errors[context], msg)
}

func (n *Notifications) HasErrors() bool {
	return len(n.errors) != 0
}

func (n *Notifications) GetErrors(context ...string) string {
	b := new(bytes.Buffer)

	hasContext := len(context) != 0

	for key, value := range n.errors {
		if hasContext && key != context[0] {
			continue
		}
		fmt.Fprintf(b, "%s: %s\n", key, strings.Join(value, ", "))
	}

	return b.String()
}
