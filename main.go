package main

import "github.com/dynastymasra/privy/config"

func init() {
	config.Load()
	config.Logger().Setup()
}

func main() {

}
