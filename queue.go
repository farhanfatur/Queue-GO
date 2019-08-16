package queue

import (
	"fmt"
)

// Service is interface
type Service interface {
	Push(key interface{}) bool
	Pop() interface{}
	Contains(key interface{}) bool
	Len() int
	Keys() []interface{}
}

type customBridge struct {
	Queue []interface{}
}

func (c *customBridge) Push(key interface{}) bool {
	c.Queue = append(c.Queue, key)
	return true

}

func (c *customBridge) Pop() interface{} {
	if len(c.Queue) <= 0 {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Data is empty")
			}
		}()
	}
	var resEl = c.Queue[0]
	fmt.Println("Take :", resEl)
	c.Queue = c.Queue[0+1:]
	return resEl
}

func (c *customBridge) Keys() []interface{} {
	if len(c.Queue) <= 0 {
		fmt.Println("Data is empty")
	}
	var data = make([]interface{}, 0)
	for _, each := range c.Queue {
		data = append(data, each)
	}
	return data

}

func (c *customBridge) Len() int {
	if len(c.Queue) <= 0 {
		fmt.Println("Data is empty")
	}
	return len(c.Queue)
}

func (c *customBridge) Contains(key interface{}) bool {
	var exist bool
	if len(c.Queue) == 1 {
		exist = true
	} else {
		for _, g := range c.Queue[:len(c.Queue)-1] {
			if key == g {
				c.Queue = c.Queue[:len(c.Queue)-1]
				exist = false
			} else {
				exist = true
			}
		}
	}
	return exist
}

// New is take main func
func New() Service {
	return &customBridge{}
}
