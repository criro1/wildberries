package main

import (
	"fmt"
	"strconv"

	"github.com/criro1/wildberries/command/pkg/command_copy"
	"github.com/criro1/wildberries/command/pkg/command_paste"
	"github.com/criro1/wildberries/command/pkg/invoker"
	"github.com/criro1/wildberries/command/pkg/receiver"
)

const (
	amCom    = "Amount of commands = "
	errEx    = "Error the Execute method"
	errFinal = "Error, expected != result"
	everyOk  = "Everything is OK. The result is:\n\n"
	expected = "Amount of commands = 2\nButton `copy` is success\nButton `paste` is success\n"
)

func main() {
	rec := receiver.NewReceiver()
	
	copy := ctrlc.NewCtrlC(rec)
	
	paste := ctrlv.NewCtrlV(rec)
	
	inv := invoker.NewIvoker()
	inv.Add(copy)
	inv.Add(paste)

	result := amCom + strconv.Itoa(inv.Len()) + "\n"
	str, err := inv.Execute()
	if err != nil {
		fmt.Println(errEx)
	}
	result += str
	
	if expected != result {
		fmt.Println(errFinal)
	} else {
		fmt.Print(everyOk, result)
	}
}
