package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

type Item struct {
	ID      int
	Name    string
	InStock bool
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n======Menu======")
		fmt.Println("1. Read file and print every X-th line")
		fmt.Println("2. Count bytes and characters in a string")
		fmt.Println("3. Demonstrate Pass-by-Value with a Single Struct")
		fmt.Println("4. Demonstrate Pass-by-Value with a Slice of Structs")
		fmt.Println("5. Quit")
		fmt.Print("================\n")
		fmt.Print("Enter your choice (1-5): ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		choice := strings.TrimSpace(input)

		switch choice {
		case "1":
			{
				readEveryXthLine(reader)
			}
		case "2":
			{
				analyzeStringLength(reader)
			}
		case "3":
			{
				demonstrateSingleStructPassByValue()
			}
		case "4":
			{
				demonstrateSlicePassByValue()
			}
		case "5":
			{
				fmt.Println("Exiting program...\nGoodbye!")
				os.Exit(0)
			}
		default:
			{
				fmt.Println("Invalid choice...\n" +
					"Please choose a number between 1 and 5...")
			}
		}
	}
}

func readEveryXthLine(reader *bufio.Reader) {
	fmt.Print("Enter file name: ")
	filePath, err := reader.ReadString('\n')
	filePath = strings.TrimSpace(filePath)

	allLines := getFile(filePath)
	if allLines == nil {
		return
	}

	fmt.Print("Enter interval number X to print X lines: ")
	xStr, err := reader.ReadString('\n')
	x, err := strconv.Atoi(strings.TrimSpace(xStr))
	if err != nil || x <= 0 {
		fmt.Println("Invalid interval number X...\nMust be X>0...")
		return
	}

	fmt.Printf("Printing Every %d Lines of '%s'...\n ", x, filePath)

	printedCount := 0
	for i := x - 1; i < len(allLines); i += x {
		cleanedLine := clearString(allLines[i])

		fmt.Println(cleanedLine)
		printedCount++
	}

	if printedCount == 0 {
		fmt.Println("No lines found...\nFile must be empty...")
	}
}

func analyzeStringLength(reader *bufio.Reader) {
	fmt.Print("\nEnter a string (English, emojis, or non-English characters): ")
	input, _ := reader.ReadString('\n') //would not let me just use 'err' to pass value, had to use '_'
	input = strings.TrimSpace(input)

	byteCount := len(input)
	charCount := utf8.RuneCountInString(input)

	fmt.Println("\n---String Length Report---")
	fmt.Printf("Input text: 		\"%s\"\n", input)
	fmt.Printf("Byte count: 		%d bytes\n", byteCount)
	fmt.Printf("Character count: %d characters (runes)\n", charCount)
}

func modifyStruct(item Item) {
	fmt.Println("\n[Inside modifyStruct Function]")
	fmt.Printf("Received original struct value: %+v\n", item) //%+v adds struct name when printing

	item.Name = "Updated Widget"
	item.InStock = false

	fmt.Printf("Modified local struct value: %+v\n", item)
	fmt.Println("\n[Exiting modifyStruct Function]")
}

func demonstrateSingleStructPassByValue() {
	myItem := Item{
		ID:      101,
		Name:    "Standard Widget",
		InStock: true,
	}

	fmt.Printf("Initial struct in caller function: %+v\n", myItem)

	modifyStruct(myItem)

	fmt.Printf("Struct in caller function AFTER function call: %+v\n", myItem)
}

func modifySlice(items []Item) {
	fmt.Println("\n[Inside modifySlice Function]")
	fmt.Printf("Received list of structs: %+v\n", items)

	if len(items) >= 2 {
		items[1].Name = "Modified Widget"
		items[1].InStock = true
		fmt.Printf("Modified second element in list: %+v\n", items)
	} else {
		fmt.Println("List has fewer than two elements; cannot modify second struct.")
	}

	fmt.Println("\n[Exiting modifySlice Function]")
}

func demonstrateSlicePassByValue() {
	itemList := []Item{
		{ID: 201, Name: "Gadget Alpha", InStock: false},
		{ID: 202, Name: "Gadget Beta", InStock: false},
	}

	fmt.Printf("Initial list in caller function: %+v\n", itemList)

	modifySlice(itemList)

	fmt.Printf("List in caller function AFTER function call: %+v\n", itemList)
}
