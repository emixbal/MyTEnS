package main

import "github.com/thatisuday/commando"

func main() {
	commando.
		SetExecutableName("MyTens app").
		SetDescription(
			"Tool untuk mengambil log file pada OS Linux di dalam directory /var/log",
		)
	commando.
		SetExecutableName("reactor").
		SetVersion("v1.0.0").
		SetDescription("This CLI tool helps you create and manage React projects. /n ok")

	commando.Parse(nil)
}
