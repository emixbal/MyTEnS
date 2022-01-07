package main

import (
	"MyTensApp/jobs"
	"fmt"

	"github.com/thatisuday/commando"
)

func main() {
	commando.
		SetExecutableName("MyTensApp").
		SetVersion("1.0.0").
		SetDescription(
			"Tool untuk mengambil log file pada OS Linux di dalam directory /var/log",
		)

	commando.
		Register(nil).
		AddArgument("fileName", "File Location", "").                                            // get file name
		AddFlag("type, t", "Convert log file into plaintext or json ", commando.String, "text"). // get opsi output format
		AddFlag("output, o", "Set new output location", commando.String, "same").                // get opsi output location
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {

			var logPath = "/var/log/"
			filePathTxt := args["fileName"].Value
			filePath := logPath + filePathTxt

			fileTypeFinal := fmt.Sprintf("%s", flags["type"].Value)

			fileLocFinal := fmt.Sprintf("%s", flags["output"].Value)

			// menjalankan import log file
			err := jobs.Convert(filePath, fileTypeFinal, fileLocFinal)
			if err != nil {
				fmt.Println(err)
			}
		})

	commando.Parse(nil)
}
