package ises

import (
	"log"
	"net/url"
	"sync"
)

type IHandlerGenerator interface {
	GetHandler(values url.Values) HandlerFunc
}


type HandlerManager struct {
	sync.RWMutex
	handlers map[string]IHandlerGenerator
}

//
//func init()  {
//	RegisterFunctionalGenerator(gin.Logger, gin.Recovery)
//}


//panic if name the same
func (h *HandlerManager)Register(k string, v IHandlerGenerator){
	h.Lock()
	defer h.Unlock()
	if old, ok := h.handlers[k]; ok{
		log.Println("Warning: SKIP, Register with the same name:", k, old)
		return
	}
	h.handlers[k] = v
}

func (h *HandlerManager)UnRegister(k string)IHandlerGenerator{
	h.Lock()
	defer h.Unlock()
	if old, ok := h.handlers[k]; ok{
		delete(h.handlers, k)
		return old
	}
	return nil
}


func (h *HandlerManager)Get(k string)IHandlerGenerator{
	h.RLock()
	defer h.RUnlock()
	if old, ok := h.handlers[k]; ok{
		return old
	}
	return nil
}

//-----------------------------------------------------------------------



/*
var(
	gManager = &HandlerManager{
		handlers: map[string]IHandlerGenerator{},
	}
)
//apis
func Register(k string, v IHandlerGenerator){
	log.Println("register:", k)
	gManager.Register(k, v)
}

func UnRegister(k string)IHandlerGenerator{
	return gManager.UnRegister(k)
}


func GetHandlerGenerator(k string)IHandlerGenerator{
	return gManager.Get(k)
}
*/


//-----------------------------------------------------------------------