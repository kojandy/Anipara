package main

type Setting struct {
	Subscribe []Subscribe
	Directory Directory
}

type Directory struct {
	Source string
	Target string
}

type Subscribe struct {
	Source   string
	Subtitle Subtitle
	Title    string
}
