package handlers

import (
	"blog/gitmodel"
	genproto "blog/protogen"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

type Server struct {
	Todos []*genproto.TodoObject
	Repos *genproto.LanguageResponse
	genproto.UnimplementedTodoServiceServer
}

func (s *Server) GetTodos(_ context.Context, _ *genproto.GetTodos) (*genproto.TodoResponse, error) {
	slog.Info("Get Info")
	return &genproto.TodoResponse{Todos: s.Todos}, nil
}

func (s *Server) AddTodo(_ context.Context, newTodo *genproto.AddTodos) (*genproto.TodoObject, error) {
	slog.Info(fmt.Sprintf("Received new task %s ", newTodo.Task))
	todoobject := &genproto.TodoObject{Id: uuid.New().String(), Task: newTodo.Task}
	s.Todos = append(s.Todos, todoobject)
	return todoobject, nil
}

func (s *Server) DeleteTodo(_ context.Context, deleteTodo *genproto.DeleteTodos) (*genproto.DeleteResponse, error) {
	var updatedTodos []*genproto.TodoObject
	for index, todos := range s.Todos {
		if todos.Id == deleteTodo.Id {
			updatedTodos = append(s.Todos[:index], s.Todos[index+1:]...)
			break
		}
	}
	s.Todos = updatedTodos
	return &genproto.DeleteResponse{Message: "Success"}, nil
}

func (s *Server) GetRepos(_ context.Context, _ *genproto.GetRepos) (*genproto.LanguageResponse, error) {
	slog.Info("Get Repos")
	return s.Repos, nil
}

func (s *Server) LibResponse(_ context.Context, language *genproto.Language) (*genproto.LanguageResponse, error) {

	queryString := strings.ToLower(language.GetName())

	switch {
	case queryString == "c#":
		queryString = "csharp"
	}

	data, err := Queries(queryString)
	if err != nil {
		slog.Error(err.Error())
	}

	var answer = &genproto.LanguageResponse{}

	for _, v := range data.Items {
		var protoLib = &genproto.LanguageLib{Name: v.Name, Stars: v.Stars}
		answer.Libraries = append(answer.Libraries, protoLib)
	}

	s.Repos = answer

	return answer, nil

}

func Queries(language string) (gitmodel.Data, error) {
	query := fmt.Sprintf("https://api.github.com/search/repositories?q=language:%s&sort=stars", language)

	req, err := http.Get(query)
	if err != nil {
		slog.Error(err.Error())
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
