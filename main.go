package main
import "fmt"

func main() {
	var tasks []string
	var task  string
	var choice int

	{
		
		fmt.Println("MY TODO LIST")
		fmt.Println("1. Create Task")
		fmt.Println("2. View Tasks")
		fmt.Println("3. Exit")
		
	}

	switch choice{
	case 1: 
		fmt.Println("Input task:")
		fmt.Scan(&task)
		tasks = append(tasks, task)
		fmt.Println("Your task has been added")
	case 2:
		fmt.Println("Your Tasks")
		if len(tasks) == 0 {
			fmt.Println("No current tasks")
		} else {
			for i, t := range tasks {
				fmt.Println(i+1, t)
			}
		}
	case 3:
		fmt.Println("See you later")
		return
	default:
		fmt.Println("Try again")

		
	}
}