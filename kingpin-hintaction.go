package main

import (
	"log"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New("kingpin-hintaction", "demo how hintaction breaks arg order")

	correctOrder     = app.Command("correctorder", "Command with correct arg order")
	correctOrderArg1 = correctOrder.Arg("arg1", "Argument 1").Required().HintAction(giveHintOk).String()
	correctOrderArg2 = correctOrder.Arg("arg2", "Argument 2").Required().String()

	brokenOrder     = app.Command("brokenorder", "Command with broken arg order")
	brokenOrderArg1 = brokenOrder.Arg("barg1", "Argument 1").Required().HintAction(giveHint).String()
	brokenOrderArg2 = brokenOrder.Arg("barg2", "Argument 2").Required().String()

	argSet = app.Flag("argset", "set of arguments to use").Default("x").String()
)

type correctOrderCommandStruct struct{}

type brokenOrderCommandStruct struct{}

var args = []string{"foo", "bar"}

func giveHint() []string {
	return args
}
func giveHintOk() []string {
	return []string{"foo", "bar"}
}

func (n *correctOrderCommandStruct) run(c *kingpin.ParseContext) error {
	log.Printf("correct order command")
	return nil
}

func (n *brokenOrderCommandStruct) run(c *kingpin.ParseContext) error {
	log.Printf("broken order command")
	return nil
}

func main() {
	app.Version("0.1")
	app.HelpFlag.Short('h')

	correctOderCommand := &correctOrderCommandStruct{}
	correctOrder.Action(correctOderCommand.run)
	brokenOderCommand := &brokenOrderCommandStruct{}
	brokenOrder.Action(brokenOderCommand.run)

	_, err := app.Parse(os.Args[1:])
	if err != nil {
		log.Printf("Error: %v", err)
		os.Exit(1)
	}
}
