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
	mb := NewMailbox("test")
	defer mb.Close()

	m := mb.ReadMessage(1)

	fmt.Printf("\n%v\n", m)
}

func TestCountMessage1(t *testing.T) {
	mb := NewMailbox("taratata")
	defer mb.Close()
	m := mb.CountMessages()
	fmt.Printf("\nThere are %v message in the inbox (first page) of %v. \n", m, mb.user)
}

func TestCountMessage2(t *testing.T) {
	mb := NewMailbox("1232164464221122")
	defer mb.Close()
	m := mb.CountMessages()
	fmt.Printf("\nThere are %v message in the inbox (first page) of %v. \n", m, mb.user)
}

func TestPrintEmptyMessage(t *testing.T) {
	var m Message
	fmt.Println(m.String())
}

func TestReadMessageNo(t *testing.T) {
	mb := NewMailbox("rodney")
	defer mb.Close()

	m := mb.ReadMessage(1)
	fmt.Printf("\n--------------#1------------------ %v\n", m)

	m = mb.ReadMessage(5)
	fmt.Printf("\n--------------#5------------------ %v\n", m)

	m = mb.ReadMessage(25)
	fmt.Printf("\n--------------#25----------------- %v\n", m)

}

func TestScrollPage(t *testing.T) {

	mb := NewMailbox("arthur")
	defer mb.Close()

	var e error
	for i := 1; e == nil; i++ {
		fmt.Printf("\nPage %d : %d messages", i, mb.CountMessages())
		e = mb.nextPage()
	}
	fmt.Println()
}
