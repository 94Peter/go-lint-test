package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/94peter/sterna/util"
	"github.com/joho/godotenv"
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
	_ = isDev
	path, err := os.Getwd()
	handler(err)
	envFile := path + "/.env"
	if util.FileExists(envFile) {
		err = godotenv.Load(envFile)
		handler(err)
	}

	fmt.Println("just test")
}

func handler(err error) {
	if err != nil {
		panic(err)
	}
}
