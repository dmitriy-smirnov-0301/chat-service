package chat_service

import (
	chatpb "chat-service/pkg/proto/chat/v1"
	"context"
	"log"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"
)

// ChatServiceServer ...
type ChatServiceServer struct {
	chatpb.UnimplementedChatServiceServer
}

// CreateChat ...
func (s *ChatServiceServer) CreateChat(_ context.Context, req *chatpb.CreateChatRequest) (*chatpb.CreateChatResponse, error) {

	log.Printf("[CreateChat] usernames: %v", req.Usernames)

	return &chatpb.CreateChatResponse{ChatId: time.Now().Unix()}, nil

}

// ConnectChat ...
func (s *ChatServiceServer) ConnectChat(_ context.Context, req *chatpb.ConnectChatRequest) (*chatpb.ConnectChatResponse, error) {

	log.Printf("[ConnectChat] chat_id: %d, username: %s", req.ChatId, req.Username)

	return &chatpb.ConnectChatResponse{WelcomeMessage: "Welcome to chat!"}, nil

}

// SendMessage ...
func (s *ChatServiceServer) SendMessage(_ context.Context, req *chatpb.SendMessageRequest) (*emptypb.Empty, error) {

	log.Printf("[SendMessage] chat_id: %d, from: %s, text: %s, timestamp: %v", req.ChatId, req.From, req.Text, req.Timestamp.AsTime())

	return &emptypb.Empty{}, nil

}

// EditMessage ...
func (s *ChatServiceServer) EditMessage(_ context.Context, req *chatpb.EditMessageRequest) (*emptypb.Empty, error) {

	log.Printf("[EditMessage] chat_id: %d, message_id: %d, new_text: %s", req.ChatId, req.MessageId, req.NewText)

	return &emptypb.Empty{}, nil

}

// DeleteMessage ...
func (s *ChatServiceServer) DeleteMessage(_ context.Context, req *chatpb.DeleteMessageRequest) (*emptypb.Empty, error) {

	log.Printf("[DeleteMessage] chat_id: %d, message_id: %d", req.ChatId, req.MessageId)

	return &emptypb.Empty{}, nil

}

// DeleteChat ...
func (s *ChatServiceServer) DeleteChat(_ context.Context, req *chatpb.DeleteChatRequest) (*emptypb.Empty, error) {

	log.Printf("[DeleteChat] chat_id: %d", req.ChatId)

	return &emptypb.Empty{}, nil

}
