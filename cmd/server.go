package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"todo/todo"
)

func main() {
	setupServer()
}

func setupServer() {
	var err error
	log.Println("Register tasks server")
	server := grpc.NewServer()
	tasks := MyTasksServer{}
	todo.RegisterTasksServer(server, tasks)
	log.Println("Start listen to :8080")
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("could not listen to :8080: %v", err)
	}
	log.Println("Start serve server")
	err = server.Serve(listener)
	if err != nil {
		log.Fatalf("could not serve: %v", err)
	}
}

type MyTasksServer struct {
}

func (s MyTasksServer) List(context.Context, *todo.Void) (*todo.TaskList, error) {
	var tasks []*todo.Task
	tasks = append(tasks, &todo.Task{Text: "wash dishes", Done: false})
	tasks = append(tasks, &todo.Task{Text: "kadusch", Done: true})

	log.Println("List has been called by any client")

	return &todo.TaskList{Tasks: tasks}, nil
}
