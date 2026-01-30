package main

type Worker interface {
	Save()
	Run()
	Get()
	Delete()
}
