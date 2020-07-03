package observer

import "fmt"

type Observer interface {
	Update()
}

type Subject struct {
	observers []Observer
	ontext    string
}

func NewSubject() *Subject {
	return &Subject{
		observers: make([]Observer, 0),
	}
}

func (s *Subject) Attach(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *Subject) Notifycation() {
	for _, v := range s.observers {
		v.Update()
	}
}

type Reader struct {
	reader string
}

func NewReader(reader string) *Reader {
	return &Reader{
		reader: reader,
	}
}

func (r *Reader) Update() {
	fmt.Printf("===== reader: %s\n", r.reader)
}
