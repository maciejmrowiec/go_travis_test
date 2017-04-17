package main

import (
	cli "github.com/maciejmrowiec/go_travis_test/Godeps/_workspace/src/github.com/ivpusic/go-clicolor/clicolor"
	"github.com/maciejmrowiec/go_travis_test/sub"
)

func PrintMeSomething() {
	cli.Print("some text").In("green")
	cli.Print("{red}Some text in red. {white}Some text in white. {default}Some text in default color").InFormat()
}

func PrintForThePeople() {
	cli.Print("some text").In("blue")
}

func PrintForThePeople2() {
	cli.Print("some text").In("blue")
}

func main() {
	PrintMeSomething()
	sub.PrintMeSomething()
}
