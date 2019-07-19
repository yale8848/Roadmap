// Create by Yale 2019/7/19 10:04
package main

import (
	"fmt"
)

type requestType int

const (
	add requestType = iota
	sub
)

type message struct {
	rt requestType
}

type Service struct {
	queue chan *message
	v int
}

func NewService(buffer int) *Service  {
	sv:=  &Service{queue:make(chan *message,buffer)}
	go sv.schedule()
	return sv
}
func (sv *Service)Value()int{
	return sv.v
}
func (sv *Service)schedule()  {

	for v:=range sv.queue{
		if v.rt==add {
			sv.v++
		}else if v.rt == sub {
				sv.v--
		}
	}
}
func (sv *Service)Sub() {
	ms:=&message{rt:sub}
	sv.queue<-ms
}
func (sv *Service)Add(){

	ms:=&message{rt:add}
	sv.queue<-ms
}

func main()  {
	sv:=NewService(1)

	sv.Add()
	sv.Add()
	sv.Sub()

	fmt.Printf("%d",sv.Value())

}