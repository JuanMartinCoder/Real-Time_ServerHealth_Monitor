package main

type Routes struct {
	MAIN    string
	MONITOR string
}

var RoutesInstance = Routes{
	MAIN:    "/",
	MONITOR: "/monitor",
}
