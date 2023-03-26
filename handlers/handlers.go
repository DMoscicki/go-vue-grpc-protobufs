package handlers

import (
	"blog/gitmodel"
	"blog/protogen"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"
	"strings"
)

type Server struct {
	Todos []*protogen.TodoObject
	Repos *protogen.LanguageResponse
	protogen.UnimplementedTodoServiceServer
}

func (s *Server) GetTodos(_ context.Context, _ *protogen.GetTodos) (*protogen.TodoResponse, error) {
	log.Println("Get task")
	return &protogen.TodoResponse{Todos: s.Todos}, nil
}

func (s *Server) AddTodo(_ context.Context, newTodo *protogen.AddTodos) (*protogen.TodoObject, error) {
	log.Printf("Received new task %s ", newTodo.Task)
	todoobject := &protogen.TodoObject{Id: uuid.New().String(), Task: newTodo.Task}
	s.Todos = append(s.Todos, todoobject)
	return todoobject, nil
}

func (s *Server) DeleteTodo(_ context.Context, deleteTodo *protogen.DeleteTodos) (*protogen.DeleteResponse, error) {
	var updatedTodos []*protogen.TodoObject
	for index, todos := range s.Todos {
		if todos.Id == deleteTodo.Id {
			updatedTodos = append(s.Todos[:index], s.Todos[index+1:]...)
			break
		}
	}
	s.Todos = updatedTodos
	return &protogen.DeleteResponse{Message: "Success"}, nil
}

func (s *Server) GetRepos(_ context.Context, _ *protogen.GetRepos) (*protogen.LanguageResponse, error) {
	log.Println("Get Repos")
	return s.Repos, nil
}

func (s *Server) LibResponse(_ context.Context, language *protogen.Language) (*protogen.LanguageResponse, error) {

	queryString := strings.ToLower(language.GetName())

	switch {
	case queryString == "c#":
		queryString = "csharp"
	}

	data, err := Queries(queryString)
	if err != nil {
		log.Println(err)
	}

	var answer = &protogen.LanguageResponse{}

	for _, v := range data.Items {
		var protoLib = &protogen.LanguageLib{Name: v.Name, Stars: v.Stars}
		answer.Libraries = append(answer.Libraries, protoLib)
	}

	s.Repos = answer

	return answer, nil

}

func Queries(language string) (gitmodel.Data, error) {
	query := fmt.Sprintf("https://api.github.com/search/repositories?q=language:%s&sort=stars", language)

	req, err := http.Get(query)
	if err != nil {
		log.Println(err)
	}

	req.Header.Add("Accept", "application/vnd.github+json")

	defer req.Body.Close()

	var libs gitmodel.Data

	body := json.NewDecoder(req.Body)

	err = body.Decode(&libs)
	if err != nil {
		return gitmodel.Data{}, err
	}

	return libs, nil
}
