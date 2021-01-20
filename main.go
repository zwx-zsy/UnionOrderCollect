package main

import (
	"UnionOrderCollect/lib"

	"github.com/jasonlvhit/gocron"
)

func main() {
	s := gocron.NewScheduler()
	//err := s.Every(1).Minute().Do(lib.JdOrderTask)
	err := s.Every(10).Second().Do(lib.UserTask)
	if err != nil {
		return
	}
	sc := s.Start() // keep the channel
	<-sc
}
