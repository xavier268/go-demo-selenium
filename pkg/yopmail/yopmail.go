package yopmail

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/tebeka/selenium"
)

// Yopmail provides access to anonymous mail
// based on the public webmail www.yopmail.com

// Message contains a single message
type Message struct {
	from       string
	topic      string
	content    string
	received   time.Time
	downloaded time.Time
}

func (m *Message) String() string {
	return fmt.Sprintf("\nMessage dump\n\tfrom       : %v\n\ttopic      : %v\n\treceived   : %v\n\tdownloaded : %v\n\tcontent    : %v\n",
		m.from,
		m.topic,
		m.content,
		m.received,
		m.downloaded)
}

// HubURL defines the url to connect to the selenium Hub
var HubURL = "http://127.0.0.1:4444/wd/hub"

var debug = false
var yopLocation, _ = time.LoadLocation("Europe/Berlin")

// YopURL is the URL to the webmail site
const YopURL = "http://yopmail.com"

// YopTimeFormat is the time layout for Yop dates
const YopTimeFormat = "Date: 2006-01-02 15:04"

// SetDebug (de)activates the debug mode
func SetDebug(flag bool) {
	debug = flag
}

// Mailbox defines a mailbox for a given user
type Mailbox struct {
	wd   selenium.WebDriver
	user string
}

// NewMailbox creates a new mail box, and go to the related page
func NewMailbox(user string) (mb *Mailbox) {
	selenium.SetDebug(debug)
	caps := selenium.Capabilities{
		"browserName": "chrome",
		"platform":    "linux"}
	wd, e1 := selenium.NewRemote(caps, HubURL)
	if e1 != nil {
		panic(e1)
	}
	wd.Get(YopURL)
	we, e2 := wd.FindElement(selenium.ByID, "login")
	if e2 != nil {
		panic(e2)
	}
	if e3 := we.SendKeys(user); e3 != nil {
		panic(e3)
	}
	if e4 := we.Submit(); e4 != nil {
		panic(e4)
	}
	return &Mailbox{wd, user}
}

// Close an existing mailbox
func (mb *Mailbox) Close() {
	if mb.wd != nil {
		mb.wd.Quit()
		mb.wd = nil
	}
	mb.user = ""
}

func (mb *Mailbox) readMessage() (mess *Message) {
	if mb.wd == nil {
		panic("Attempt to readMessage from a closed mailbox !")
	}
	// Enter the message frame
	mb.wd.SwitchFrame("ifmail")
	defer mb.wd.SwitchFrame("")

	mess = new(Message)
	mess.downloaded = time.Now()

	we, e := mb.wd.FindElement(selenium.ByTagName, "body")
	if e != nil {
		panic(e)
	}
	var w selenium.WebElement
	w, e = we.FindElement(selenium.ByXPATH, "(.//div[@id='mailhaut']/div)[1]")
	if e != nil {
		log.Print(e)
	} else {
		mess.topic, _ = w.Text()
	}

	w, _ = we.FindElement(selenium.ByXPATH, "(.//div[@id='mailhaut']/div)[2]")
	if e != nil {
		log.Print(e)
	} else {
		tt, _ := w.Text()
		mess.from = strings.Split(tt, "From: ")[1]
	}

	w, _ = we.FindElement(selenium.ByXPATH, "(.//div[@id='mailhaut']/div)[4]")
	if e != nil {
		log.Print(e)
	} else {
		tt, _ := w.Text()
		mess.received, _ = time.ParseInLocation(YopTimeFormat, tt, yopLocation)
	}

	w, _ = we.FindElement(selenium.ByXPATH, ".//div[@id='mailmillieu']")
	if e != nil {
		log.Print(e)
	} else {
		mess.content, _ = w.Text()
	}

	return mess
}
