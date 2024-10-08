package main

import (
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
	usr := GetUser(in.(string))
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
	/*	outId := make(chan []MsgID)*/
	for val := range in {
		user := val.(User)
		if fuser.ID != 0 {
			wg.Add(1)
			go GetMessagesWorker([]User{fuser, user}, out, &wg)
			fuser.ID = 0
			/*inUsers <- users
			out <- <-outId*/
		} else {
			fuser = user
		}
	}
	if fuser.ID != 0 {
		res, err := GetMessages(fuser)
		if err != nil {
			panic(err)
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
		panic(err)
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
		msg := val.(MsgID)
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
	res := make([]string, 0, 10)
	strData := ""
	for val := range in {
		data := val.(MsgData)
		strData = strconv.FormatBool(data.HasSpam) + " " + strconv.FormatUint(uint64(data.ID), 10)
		res = append(res, strData)
	}
	sort.Slice(res, func(i, j int) bool {
		word1, word2 := res[i], res[j]
		if word1[0] == 't' && word2[0] != 't' {
			return true
		}
		if word2[0] == 't' && word1[0] != 't' {
			return false
		}
		id1, id2 := strings.Fields(word1)[1], strings.Fields(word2)[1]
		uid1, _ := strconv.ParseUint(id1, 10, 64)
		uid2, _ := strconv.ParseUint(id2, 10, 64)
		return uid1 < uid2
	})
	out <- res
}
