package main

import (
	"fmt"
	"github.com/PlagueCat-Miao/goipfs-lab511/constdef"
	"github.com/PlagueCat-Miao/goipfs-lab511/nodes"
	"log"
)

func main() {
	_, err := nodes.InitEdgeServive()
	nodes.EdgeServerListen(constdef.EdgePort)
	if err != nil {
		log.Printf("[InitEdgeServive-err]: %v", err)
		return
	}
	for {
		printHelp()
		var order string
		fmt.Scanf("%s", &order)
		switch order {
		case "0":
			return
		case "1":
			nodes.EdgeLogin()
		case "2":
			nodes.EdgeAddFile()
		case "3":
			nodes.EdgeReadFile()
		case "4":
			nodes.EdgeGetFile()
		default:
			fmt.Println("unknow order, try again")
		}

	}

}

func printHelp() {
	fmt.Println("1. login")
	fmt.Println("2. Add file")
	fmt.Println("3. Read FileInfo")
	fmt.Println("4. Get File")

	fmt.Println("0. Exit")
}

// +build ignore
