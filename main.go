package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/94peter/sterna/util"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

var (
	v         = flag.Bool("v", false, "version")
	isDev     = flag.Bool("dev", false, "dev mode")
	Version   = "1.0.0"
	BuildTime = time.Now().Local().GoString()
)

func main() {
	flag.Parse()

	if *v {
		fmt.Println(Version)
		fmt.Println("Build Time: " + BuildTime)
		return
	}

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	envFile := path + "/.env"
	if util.FileExists(envFile) {
		err = godotenv.Load(envFile)
		if err != nil {
			panic(errors.Wrap(err, "load .env file fail"))
		}
	}

	fmt.Println("just test")
}
