package yopmail

import (
	"fmt"
	"testing"
)

func TestOpenClose(t *testing.T) {
	mb := NewMailbox("toto")
	defer mb.Close()
	ttl, _ := mb.wd.Title()
	fmt.Printf("\nTitle : %v\n", ttl)
}

func TestReadOneMessage(t *testing.T) {
	mb := NewMailbox("taratata")
	defer mb.Close()

	m := mb.readMessage()

	fmt.Printf("\nMESSAGE DUMP\n%v\n", m)

}
