package main

import cw "github.com/sidav/goLibRL/console"

func main() {
	cw.Init_console("Wololo!", cw.SDLRenderer)
	defer cw.Close_console()

	cw.PutString("zomg", 0, 0)
	cw.Flush_console()
	cw.ReadKey()

}
