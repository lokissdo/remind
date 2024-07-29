package service

import (
	"context"
	"remind/todo/model"
	pb "remind/todo/pb"
	"remind/todo/repository"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// TodoServiceServer is the server API for TodoService service.
type TodoServiceServer struct {
	pb.UnimplementedTodoServiceServer
}

// CreateTodo creates a new todo
func (s *TodoServiceServer) CreateTodo(ctx context.Context, req *pb.CreateTodoRequest) (*pb.CreateTodoResponse, error) {
	var todo model.Todo
	todo = fromProto(req.GetTodo())
	todo.UserID = getUsernameFromContext(ctx)
	FCM_token := req.GetFCMToken()
	err := repository.CreateOrUpdateFCMToken(ctx, todo.UserID, FCM_token)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create todo: %v", err)
	}
	for i := range todo.Reminders {
		todo.Reminders[i].UserID = getUsernameFromContext(ctx)
        todo.Reminders[i].Sent = false
	}

	createdTodo, err := repository.CreateTodo(todo)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create todo: %v", err)
	}

	return &pb.CreateTodoResponse{Todo: toProto(createdTodo)}, nil
}

// GetTodoByID retrieves a todo by its ID
func (s *TodoServiceServer) GetTodoByID(ctx context.Context, req *pb.GetTodoByIDRequest) (*pb.GetTodoByIDResponse, error) {
	todo, err := repository.GetTodoByID(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Todo not found: %v", err)
	}
	return &pb.GetTodoByIDResponse{Todo: toProto(todo)}, nil
}

// ListTodos lists all todos
func (s *TodoServiceServer) ListTodos(ctx context.Context, req *pb.ListTodosRequest) (*pb.ListTodosResponse, error) {
	todos, err := repository.ListAllTodosByUserID(getUsernameFromContext(ctx))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to fetch todos: %v", err)
	}

	var protoTodos []*pb.Todo
	for _, todo := range todos {
		protoTodos = append(protoTodos, toProto(todo))
	}

	return &pb.ListTodosResponse{Todos: protoTodos}, nil
}

// UpdateTodo updates a todo
func (s *TodoServiceServer) UpdateTodo(ctx context.Context, req *pb.UpdateTodoRequest) (*pb.UpdateTodoResponse, error) {
	updatedTodo := fromProto(req.GetTodo())
	updatedTodo.UserID = getUsernameFromContext(ctx)

	todo, err := repository.UpdateTodoByID(updatedTodo.ID, updatedTodo)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to update todo: %v", err)
	}

	return &pb.UpdateTodoResponse{Todo: toProto(todo)}, nil
}

// DeleteTodo deletes a todo by its ID
func (s *TodoServiceServer) DeleteTodo(ctx context.Context, req *pb.DeleteTodoRequest) (*pb.DeleteTodoResponse, error) {
	err := repository.DeleteTodoByID(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to delete todo: %v", err)
	}
	return &pb.DeleteTodoResponse{Message: "Todo deleted successfully"}, nil
}

// Helper functions to convert between proto and model

func fromProto(protoTodo *pb.Todo) model.Todo {
	return model.Todo{
		ID:          protoTodo.GetId(),
		Title:       protoTodo.GetTitle(),
		Description: protoTodo.GetDescription(),
		StartTime:   protoTodo.GetStartTime().AsTime(),
		EndTime:     protoTodo.GetEndTime().AsTime(),
		UserID:      protoTodo.GetUserId(),
		Reminders:   fromProtoReminders(protoTodo.GetReminders()),
	}
}

func toProto(todo model.Todo) *pb.Todo {
	return &pb.Todo{
		Id:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		StartTime:   timestamppb.New(todo.StartTime),
		EndTime:     timestamppb.New(todo.EndTime),
		UserId:      todo.UserID,
		Reminders:   toProtoReminders(todo.Reminders),
	}
}

func fromProtoReminders(protoReminders []*pb.Reminder) []model.Reminder {
	var reminders []model.Reminder
	for _, protoReminder := range protoReminders {
		reminders = append(reminders, model.Reminder{
			ID:      protoReminder.GetId(),
			Start:   protoReminder.GetStart().AsTime(),
			Message: protoReminder.GetMessage(),
			TodoID:  protoReminder.GetTodoId(),
			UserID:  protoReminder.GetUserId(),
			Sent:    protoReminder.GetSent(),
		})
	}
	return reminders
}

func toProtoReminders(reminders []model.Reminder) []*pb.Reminder {
	var protoReminders []*pb.Reminder
	for _, reminder := range reminders {
		protoReminders = append(protoReminders, &pb.Reminder{
			Id:      reminder.ID,
			Start:   timestamppb.New(reminder.Start),
			Message: reminder.Message,
			TodoId:  reminder.TodoID,
			UserId:  reminder.UserID,
			Sent:    reminder.Sent,
		})
	}
	return protoReminders
}
func getUsernameFromContext(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		println("No metadata found in context")
		return ""
	}
	usernames := md.Get("username")
	if len(usernames) > 0 {
		return usernames[0]
	}
	return ""
}
