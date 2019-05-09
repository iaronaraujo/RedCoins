package main

import (
	r "github.com/iaronaraujo/RedCoins/routers"
)

func main() {
	e := r.App
	e.Start(":3000")
}
