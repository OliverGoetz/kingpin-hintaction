package main

import (
	"log"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app    = kingpin.New("kingpin-hintaction", "demo how hintaction breaks arg order")
	argSet = app.Flag("argset", "set of arguments to use").Default("x").String()
)

var (
	correctOrder     *kingpin.CmdClause
	correctOrderArg1 *string
	correctOrderArg2 *string
)

func initCorrectOrderCommand() {
	correctOrder = app.Command("correctorder", "Command with correct arg order")
	correctOrderArg1 = correctOrder.Arg("arg1", "Argument 1").Required().HintAction(giveHintStatic).String()
	correctOrderArg2 = correctOrder.Arg("arg2", "Argument 2").Required().String()

	correctOderCommand := &correctOrderCommandStruct{}
	correctOrder.Action(correctOderCommand.run)
}

var (
	fixedOrder     *kingpin.CmdClause
	fixedOrderArg1 *string
	fixedOrderArg2 *string
)

func initFixedOrderCommand() {
	fixedOrder = app.Command("fixedorder", "Command with fixed arg order")
	fixedOrderArg1 = fixedOrder.Arg("barg1", "Argument 1").Required().HintAction(giveHintFromVar).String()
	fixedOrderArg2 = fixedOrder.Arg("barg2", "Argument 2").Required().String()

	fixedOderCommand := &fixedOrderCommandStruct{}
	fixedOrder.Action(fixedOderCommand.run)
}

type correctOrderCommandStruct struct{}

type fixedOrderCommandStruct struct{}

func giveHintStatic() []string {
	log.Printf("Executing giveHintStatic")
	return []string{"foo", "bar"}
}

var args = []string{"foo", "bar"}

func giveHintFromVar() []string {
	log.Printf("Executing giveHintFromVar")
	return args
}
func (n *correctOrderCommandStruct) run(c *kingpin.ParseContext) error {
	log.Printf("correct order command")
	log.Printf("arg1: %v, arg2: %v", *correctOrderArg1, *correctOrderArg2)
	log.Printf("Hints: %v", giveHintStatic())
	return nil
}

func (n *fixedOrderCommandStruct) run(c *kingpin.ParseContext) error {
	log.Printf("fixed order command")
	log.Printf("barg1: %v, barg2: %v", *fixedOrderArg1, *fixedOrderArg2)
	log.Printf("Hints: %v", giveHintFromVar())
	return nil
}

func main() {
	initCorrectOrderCommand()
	initFixedOrderCommand()

	app.Version("0.1")
	app.HelpFlag.Short('h')

	_, err := app.Parse(os.Args[1:])
	if err != nil {
		log.Printf("Error: %v", err)
		os.Exit(1)
	}
}
