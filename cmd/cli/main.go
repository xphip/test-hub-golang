package main

import "test-hub-golang/pkg/cli"

func main() {

	app := cli.NewCli()
	if errRun := app.Run(); errRun != nil {
		panic(errRun)
	}
}
