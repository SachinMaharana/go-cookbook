package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("usage: siri <command> [<args>]")
		fmt.Println("The most commonly used git commands are: ")
		fmt.Println(" ask   Ask questions")
		fmt.Println(" send  Send messages to your contacts")
		return
	}

	askCommand := flag.NewFlagSet("ask", flag.ExitOnError)
	questionFlag := askCommand.String("question", "", "Question you are asking..")

	sendCommand := flag.NewFlagSet("send", flag.ExitOnError)
	recipientFlag := sendCommand.String("recipient", "", "Recipient of your message")
	messageFlag := sendCommand.String("message", "", "Text message")

	switch os.Args[1] {
	case "ask":
		askCommand.Parse(os.Args[2:])
	case "send":
		sendCommand.Parse(os.Args[2:])
	default:
		fmt.Printf("%q is not valid command.\n", os.Args[1])
		os.Exit(2)
	}
	if askCommand.Parsed() {
		if *questionFlag == "" {
			fmt.Println("Please supply the question using -question option.")
			return
		}
		fmt.Printf("You asked: %q\n", *questionFlag)
	}

	if sendCommand.Parsed() {
		if *recipientFlag == "" {
			fmt.Println("Please supply the recipient using -recipient option.")
			return
		}

		if *messageFlag == "" {
			fmt.Println("Please supply the message using -message option.")
			return
		}

		fmt.Printf("Your message is sent to %q.\n", *recipientFlag)
		fmt.Printf("Message: %q.\n", *messageFlag)
	}
}

//siri send -recipient=john@example.com -message="Call me?"
//siri ask -question="What is the whether in London?"
