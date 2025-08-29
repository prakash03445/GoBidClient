package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/prakash03445/GoBidClient/client"
)

func main() {
	serverAddr := "localhost:50051"
	c := client.NewGoBidClient(serverAddr)
	defer c.Close()

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("GoBid Client CLI")
	fmt.Println("----------------")

	for {
		fmt.Print("> ")
		cmdLine, _ := reader.ReadString('\n')
		cmdLine = strings.TrimSpace(cmdLine)
		parts := strings.Split(cmdLine, " ")

		switch parts[0] {
		case "exit":
			fmt.Println("Exiting...")
			return

		case "add":
			if len(parts) < 5 {
				fmt.Println("Usage: add <id> <name> <description> <startingPrice>")
				continue
			}
			price, err := strconv.ParseFloat(parts[4], 64)
			if err != nil {
				fmt.Println("Invalid price")
				continue
			}
			c.AddProduct(parts[1], parts[2], parts[3], price)

		case "list":
			c.GetProducts()
		
		case "help":
			fmt.Println("Available commands:")
			fmt.Println("  add <id> <name> <description> <startingPrice>  - Add a new product")
			fmt.Println("  list                                           - List all products")
			fmt.Println("  help                                           - Show this help message")
			fmt.Println("  exit                                           - Exit the CLI")

		default:
			fmt.Println("Unknown command. Type 'help' to see the list of available commands.")
		}
	}
}
