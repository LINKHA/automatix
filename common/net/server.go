package net

type Server struct {
	Name string
	IP   string
	Port int

	ConnMgr ConnManager
}
