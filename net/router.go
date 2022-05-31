package net

type HandlerFunc func()

type group struct {
	prefix     string
	handlerMap map[string]HandlerFunc
}

type router struct {
	group []*group
}