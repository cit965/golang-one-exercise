package main

import (
	"fmt"
	"log"
	"time"
)

type Restful interface {
	Get(id string) (interface{}, error)
}

type PodRest struct {
	Action string
	Url    string
	D      Dao
}

func (p *PodRest) Get(id string) (interface{}, error) {
	return p.D.Get(id)
}

type Dao interface {
	Get(id string) (interface{}, error)
}

type Registry struct {
	etcdClient etcdClient
	config     string
}

func (c *Registry) Get(id string) (interface{}, error) {
	return c.etcdClient.Get(id), nil
}

type etcdClient struct {
	Pods map[string]string
}

func (e *etcdClient) Get(id string) string {
	result, ok := e.Pods[id]
	if !ok {
		return "empty"
	}
	return result
}

func main() {
	//fmt.Println("hello cit")
	//url := "api/pods/3"
	//podRestful := PodRest{
	//	Action: "get",
	//	Url:    url,
	//	D:      &Registry{
	//		etcdClient: etcdClient{},
	//		config:     "",
	//	},
	//}
	//podRestful.Get("3")
	//Forever(func() {
	//	fmt.Println("hello")
	//}, 1)

	//const nihongo = "日本語dd"
	//
	//for index, runeValue := range nihongo {
	//	fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	//}
	const nihongo2 = "日本語d"
	//for index, runeValue := range nihongo2 {
	//	fmt.Printf("%v starts at byte position %d\n", runeValue, index)
	//}

	for i := 0; i < len(nihongo2); i++ {
		fmt.Println(nihongo2[i])
	}
}

func HandleCrash() {
	r := recover()
	if r != nil {
		log.Printf("Recovered from panic: %#v", r)
	}
}

// Loops forever running f every d.  Catches any panics, and keeps going.
func Forever(f func(), period time.Duration) {
	for {
		func() {
			defer HandleCrash()
			f()
		}()
		time.Sleep(period)
	}
}
