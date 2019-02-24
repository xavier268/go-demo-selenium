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

// HubURL defines the url to connect to the selenium Hub
var HubURL = "http://127.0.0.1:4444/wd/hub"

var debug = false
var yopLocation, _ = time.LoadLocation("Europe/Berlin")

// YopURL is the URL to the webmail site
const YopURL = "http://yopmail.com"

// YopTimeFormat is the time layout for Yop dates
const YopTimeFormat = "Date: 2006-01-02 15:04"

// Message contains a single message
type Message struct {
	to          string
	from        string
	topic       string
	content     string
	htmlContent string
	url         string
	received    time.Time
	downloaded  time.Time
}

// NewMessage creates a time-stamped new message.
func NewMessage() (mess *Message) {
	mess = new(Message)
	mess.downloaded = time.Now()
	return mess
}

func (m *Message) String() string {
	if debug || len(m.content) < 12 || len(m.htmlContent) < 12 {
		return fmt.Sprintf(
			"\nMessage :"+
				"\n\tto         : %v\n\tfrom       : %v"+
				"\n\ttopic      : %v\n\treceived   : %v"+
				"\n\tdownloaded : %v\n\turl        : %v"+
				"\n\tcontent    : %v\n\thtmlContent: %v\n",
			m.to, m.from,
			m.topic, m.received,
			m.downloaded, m.url,
			m.content, m.htmlContent,
		)
	}
	return fmt.Sprintf(
		"\nMessage :"+
			"\n\tto         : %v\n\tfrom       : %v"+
			"\n\ttopic      : %v\n\treceived   : %v"+
			"\n\tdownloaded : %v\n\turl        : %v"+
			"\n\tcontent    : %v\n\thtmlContent: %v\n",
		m.to, m.from,
		m.topic, m.received,
		m.downloaded, m.url,
		m.content[:10]+"[ ... truncated...]", m.htmlContent[:10]+"[ ... truncated...]",
	)
}

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
	return &Mailbox{wd: wd, user: user}
}

// Close an existing mailbox
func (mb *Mailbox) Close() {
	if mb.wd != nil {
		mb.wd.Quit()
		mb.wd = nil
	}
	mb.user = ""
}

// parse message from the provided mailbox
// Assume we are at the root of the message page/iframe
func (m *Message) parse(mb *Mailbox) {
	we, e := mb.wd.FindElement(selenium.ByTagName, "body")
	if e != nil {
		panic(e)
	}

	var w selenium.WebElement

	w, e = we.FindElement(selenium.ByXPATH, "(.//div[@id='mailhaut']/div)[1]")
	if e != nil {
		log.Print(e)
	} else {
		m.topic, _ = w.Text()
	}

	w, _ = we.FindElement(selenium.ByXPATH, "(.//div[@id='mailhaut']/div)[2]")
	if e != nil {
		log.Print(e)
	} else {
		tt, _ := w.Text()
		m.from = strings.Split(tt, "From: ")[1]
	}

	w, _ = we.FindElement(selenium.ByXPATH, "(.//div[@id='mailhaut']/div)[4]")
	if e != nil {
		log.Print(e)
	} else {
		tt, _ := w.Text()
		m.received, _ = time.Parse(YopTimeFormat, tt)
	}

	w, _ = we.FindElement(selenium.ByXPATH, ".//div[@id='mailmillieu']")
	if e != nil {
		log.Print(e)
	} else {
		m.content, _ = w.Text()
		m.htmlContent, _ = w.GetAttribute("innerHTML")
	}
}

func (mb *Mailbox) CountMessages() int {
	mb.wd.SwitchFrame("")
	mb.wd.SwitchFrame("ifinbox")

	t, e := mb.wd.FindElements(selenium.ByXPATH, ".//body//div[@class='m']")
	if e != nil {
		log.Println("No message found")
		log.Println(e)
		return 0
	}
	return len(t)
}

// ReadMessage read the message specified by its index (1-based)
// If out of bound index, return nil.
func (mb *Mailbox) ReadMessage(n int) (mess *Message) {
	mb.wd.SwitchFrame("")
	mb.wd.SwitchFrame("ifinbox")
	t, e := mb.wd.FindElement(selenium.ByID, fmt.Sprintf("m%d", n))
	if e != nil {
		log.Println("No message found")
		log.Println(e)
		return nil
	}
	t, e = t.FindElement(selenium.ByTagName, "a")
	if e != nil {
		log.Println("Message found, but no message link")
		log.Println(e)
		return nil
	}
	lk, e := t.GetAttribute("href")
	if e != nil {
		log.Println("Message found, but no message link")
		log.Println(e)
		return nil
	}

	mess = NewMessage()
	mess.url = lk
	mess.to = mb.user

	t.Click()

	mb.wd.SwitchFrame("")
	mb.wd.SwitchFrame("ifmail")
	mess.parse(mb)

	return mess
}

// nextPage moves to the next message page
// Do not change iframe - assumes you are still in listing frame
func (mb *Mailbox) nextPage() (e error) {
	w, e := mb.wd.FindElement(selenium.ByXPATH, ".//a[@class='igif next']")
	if e != nil {
		return e
	}
	e = w.Click()
	if e != nil {
		return e
	}
	return nil
}
