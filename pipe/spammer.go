package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func RunPipeline(cmds ...cmd) {
	wg := sync.WaitGroup{}
	in := make(chan interface{})
	for _, com := range cmds {
		out := make(chan interface{})
		wg.Add(1)
		go func(job cmd, in chan interface{}) {
			defer wg.Done()
			defer close(out)
			job(in, out)
		}(com, in)
		in = out
	}
	wg.Wait()
	return
}

func SelectUsers(in, out chan interface{}) {
	wg := sync.WaitGroup{}
	users := make(map[User]int)
	mu := &sync.RWMutex{}
	for val := range in {
		wg.Add(1)
		go GetUserWorker(val, out, mu, users, &wg)
	}
	wg.Wait()
}

func GetUserWorker(in interface{}, out chan interface{}, mu *sync.RWMutex, users map[User]int, wg *sync.WaitGroup) {
	defer wg.Done()
	usrString, err := in.(string)
	if err != true {
		fmt.Println(err)
	}
	usr := GetUser(usrString)
	mu.RLock()
	_, ok := users[usr]
	mu.RUnlock()
	if !ok {
		mu.Lock()
		users[usr] = 1
		mu.Unlock()
		out <- usr
	}
}

func SelectMessages(in, out chan interface{}) {
	fuser := User{}
	wg := sync.WaitGroup{}
	for val := range in {
		user, err := val.(User)
		if err != true {
			fmt.Println(err)
		}
		if fuser.ID != 0 {
			wg.Add(1)
			go GetMessagesWorker([]User{fuser, user}, out, &wg)
			fuser.ID = 0
		} else {
			fuser = user
		}
	}
	if fuser.ID != 0 {
		res, err := GetMessages(fuser)
		if err != nil {
			fmt.Println(err)
		}
		for _, id := range res {
			out <- id
		}
	}
	wg.Wait()
}

func GetMessagesWorker(in []User, out chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	res, err := GetMessages(in...)
	if err != nil {
		fmt.Println(err)
	}
	for _, id := range res {
		out <- id
	}
}

func CheckSpam(in, out chan interface{}) {
	wg := sync.WaitGroup{}
	counterChan := make(chan interface{}, 5)
	for val := range in {
		wg.Add(1)
		msg, err := val.(MsgID)
		if err != true {
			fmt.Println(err)
		}
		counterChan <- struct{}{}
		go CheckSpamWorker(msg, out, counterChan, &wg)
	}
	wg.Wait()
}

func CheckSpamWorker(msg MsgID, out chan interface{}, counter chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	res, err := HasSpam(msg)
	<-counter
	if err != nil {
		panic(err)
	}
	data := MsgData{
		ID:      msg,
		HasSpam: res,
	}
	out <- data
}

func CombineResults(in, out chan interface{}) {
	outData := make([]MsgData, 0, 100)
	for val := range in {
		data, err := val.(MsgData)
		if err != true {
			fmt.Println(err)
		}
		outData = append(outData, data)
	}
	sort.Slice(outData, func(i, j int) bool {
		msg1, msg2 := outData[i], outData[j]
		if msg1.HasSpam && !msg2.HasSpam {
			return true
		}
		if msg2.HasSpam && !msg1.HasSpam {
			return false
		}
		return msg1.ID < msg2.ID
	})
	for _, str := range outData {
		spamString := strconv.FormatBool(str.HasSpam)
		uidString := strconv.FormatUint(uint64(str.ID), 10)
		var builder strings.Builder
		builder.Grow(len(spamString) + 1 + len(uidString))
		builder.WriteString(spamString)
		builder.WriteString(" ")
		builder.WriteString(uidString)
		out <- builder.String()
	}
}
