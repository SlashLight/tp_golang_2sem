package main

import "sync"

func RunPipeline(cmds ...cmd) {
	wg := sync.WaitGroup{}
	in := make(chan interface{})
	for _, com := range cmds {
		out := make(chan interface{})
		wg.Add(1)
		go func(job cmd, in chan interface{}) {
			job(in, out)
			wg.Done()
			close(out)
		}(com, in)
		in = out
	}
	wg.Wait()
	return
}

func SelectUsers(in, out chan interface{}) {
	// 	in - string
	// 	out - User
}

func SelectMessages(in, out chan interface{}) {
	// 	in - User
	// 	out - MsgID
}

func CheckSpam(in, out chan interface{}) {
	// in - MsgID
	// out - MsgData
}

func CombineResults(in, out chan interface{}) {
	// in - MsgData
	// out - string
}
