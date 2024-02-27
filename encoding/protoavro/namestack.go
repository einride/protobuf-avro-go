package protoavro

import (
	"strings"
)

type nameStack struct {
	names []string
}

func (n *nameStack) push(name string) {
	n.names = append(n.names, strings.ToLower(name))
}

func (n *nameStack) pop() {
	n.names = n.names[:len(n.names)-1]
}

func (n *nameStack) current() string {
	return strings.Join(n.names, ".")
}

func newNameStack() nameStack {
	return nameStack{names: []string{"root"}}
}
