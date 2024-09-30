package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Status string

const (
	NotDone        Status = "not done"
	InProgress     Status = "In Progress"
	Done           Status = "Done"
	Add            string = "add"
	Update         string = "update"
	List           string = "list"
	Delete         string = "delete"
	MarkInProgress string = "mark-in-progress"
	MarkDone       string = "mark-done"
)

type Task struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status Status `json:"status"`
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var tasks []Task

	fileName := "data.json"

	// checking if json file already exists or not
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		createNewJsonFile(fileName)
	} else {
		readJsonFile(fileName, &tasks)
	}

	for {
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Something went wrong!!", err)
			return
		}

		// removing the extra space when the user hits "Enter"
		input = strings.TrimSpace(input)

		// getting the user's command and converting it to lower case
		command := strings.ToLower(strings.Split(input, " ")[0])

		if command == "exit" {
			fmt.Println("Exiting...")
			break
		}

		if !contains(command) {
			fmt.Println("Please eneter a valid command!!")
			continue
		}

		switch command {
		case Add:
			addTask(fileName, input, &tasks)
		case List:
			fmt.Println(tasks)
		case Update, MarkInProgress, MarkDone:
			fmt.Println("Update task")
		case Delete:
			fmt.Println("delete tasks")
		}
	}
}

// function for creating json file if it doesn't exists
func createNewJsonFile(fileName string) {
	fmt.Println("Json file doesn't exists")
	file, err := os.Create(fileName)

	if err != nil {
		fmt.Print("Error opening the newly created json file!")
		return
	}

	defer file.Close()

	emptyArray := []any{}
	encoder := json.NewEncoder(file)
	err = encoder.Encode(emptyArray)

	if err != nil {
		fmt.Println("Error writing empty array to json file!", err)
		return
	}

	fmt.Println("Successfully created a new json file with empty array!")
}

// function for reading json file if already exists
func readJsonFile(fileName string, tasks *[]Task) {
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error reading the JSON file", err)
		return
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)

	if err != nil {
		fmt.Println("Error DECODING the json file", err)
		return
	}
}

// function to check if the user entered a valid command
func contains(command string) bool {
	commands := []string{Add, List, Update, Delete, MarkInProgress, MarkDone}
	for _, option := range commands {
		if option == command {
			return true
		}
	}
	return false
}

// function for adding task
func addTask(fileName string, input string, tasks *[]Task) {
	inputSlice := strings.Split(input, `"`)
	inputSlice = inputSlice[:len(inputSlice)-1]

	if len(inputSlice) < 2 {
		fmt.Println("Please enter the task name followed by the add command")
		return
	}

	// making sure user doesn't do this ("        ")
	taskName := strings.TrimSpace(inputSlice[1])

	if len(taskName) < 1 {
		fmt.Println("Please enter a valid task name!")
		return
	}

	id := len(*tasks)

	newTask := Task{
		ID:     id,
		Name:   taskName,
		Status: NotDone,
	}

	fmt.Println("new tasks", newTask)

	*tasks = append(*tasks, newTask)

	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)

	if err != nil {
		fmt.Println("Error opening the file for inserting data!!", err)
		return
	}

	defer file.Close()

	// data := *tasks
	encoder := json.NewEncoder(file)
	err2 := encoder.Encode(*tasks)

	if err2 != nil {
		fmt.Println("Error enccoding tasks to write to json file!", err)
		return
	}

	fmt.Println(`Task added successfully witn ID: `, id)
}
