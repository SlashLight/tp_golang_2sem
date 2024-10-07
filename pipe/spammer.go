package main

import (
	"sync"
)

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
	users := make(map[User]int)
	for val := range in {
		usr := GetUser(string(val.([]byte)))
		if _, ok := users[usr]; !ok {
			users[usr] = 1
			out <- usr
		}
	}
}

func SelectMessages(in, out chan interface{}) {
	users := make([]User, 0, 2)
	for val := range in {
		user := val.(User)
		if len(users) == 2 {
			res, err := GetMessages(users...)
			if err != nil {
				panic(err)
			}
			users = make([]User, 0, 2)
			out <- res
		} else {
			users = append(users, user)
		}
	}
	if len(users) != 0 {
		res, err := GetMessages(users...)
		if err != nil {
			panic(err)
		}
		out <- res
	}
}

func CheckSpam(in, out chan interface{}) {
	for val := range in {
		cde3w
	}
}

func CombineResults(in, out chan interface{}) {
	// in - MsgData
	// out - string
}
