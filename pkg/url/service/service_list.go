package service

import "fmt"

type Ele struct {
	Char byte
	Next *Ele
}

type SingleList struct {
	Head *Ele
	P    [10]*Ele
}

func InitList() *SingleList {
	list := SingleList{}
	list.Fill()
	return &list
}

func (s *SingleList) Add(char byte) {
	ele := &Ele{
		Char: char,
	}
	if s.Head == nil {
		s.Head = ele
	} else {
		ele.Next = s.Head
		s.Head = ele
	}
}

func (s *SingleList) Fill() {
	for i := 48; i < 58; i++ {
		s.Add(byte(i))
	}
	for i := 65; i < 91; i++ {
		s.Add(byte(i))
	}
	s.Add('_')
	for i := 97; i < 123; i++ {
		s.Add(byte(i))
	}
	for i := range s.P {
		s.P[i] = s.Head
	}
}

func (s *SingleList) GetUrl() (url string) {
	for _, j := range s.P {
		url += string(j.Char)
	}
	var flag_next_nil bool
	if s.P[9].Next == nil {
		flag_next_nil = true
		for i := 9; i >= 0; i-- {
			if flag_next_nil == true {
				if s.P[i].Next == nil {
					s.P[i] = s.Head
					flag_next_nil = true
				} else {
					s.P[i] = s.P[i].Next
					flag_next_nil = false
				}
			}
		}
	} else {
		s.P[9] = s.P[9].Next
	}
	return
}

func (s *SingleList) PrintList() {
	if s.Head == nil {
		fmt.Println("List is empty")
	} else {
		current := s.Head
		for current.Next != nil {
			fmt.Println(string(current.Char))
			current = current.Next
		}
		fmt.Println(string(current.Char))
	}
}
