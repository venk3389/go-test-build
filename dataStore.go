package main

import (
	"fmt"
	"sync"
)

type data struct {
	store map[string]string
	mux   sync.Mutex
}

func (d *data) New() {
	if d.store == nil {
		fmt.Println("Instantiating the datastore once")
		d.store = make(map[string]string)
	}
}

func (d *data) Add(key, value string) {
	d.mux.Lock()
	d.store[key] = value
	d.mux.Unlock()
}

func (d *data) Get(key string) string {
	temp := ""
	d.mux.Lock()
	temp = d.store[key]
	d.mux.Unlock()
	return temp
}

func (d *data) Delete(key string) {
	//lock
	d.mux.Lock()
	delete(d.store, key)
	//unlock
	d.mux.Unlock()
}
