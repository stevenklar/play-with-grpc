package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"todo/todo"
)

func main() {
	setupClient()
}

func setupClient() {
	conn, e := grpc.Dial(":8080", grpc.WithInsecure())
	if e != nil {
		log.Fatalln(e)
	}
	client := todo.NewTasksClient(conn)
	taskList, e := client.List(context.Background(), &todo.Void{})
	if e != nil {
		log.Fatalln(e)
	}

	for i, task := range taskList.Tasks {
		fmt.Println(i, task.Text)
		fmt.Println("Done?", task.Done)
	}
}
