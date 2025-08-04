package main

type Subject interface {
	RegsiterObsercer(o Observer)
	RemoveObserver(o Observer)
	NotifiyObservers()
}
