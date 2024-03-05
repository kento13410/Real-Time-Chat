// Code generated by protoc-gen-entgrpc. DO NOT EDIT.
package entpb

import (
	ent "Real-Time-Chat/ent"
	chat "Real-Time-Chat/ent/chat"
	user "Real-Time-Chat/ent/user"
	context "context"
	base64 "encoding/base64"
	entproto "entgo.io/contrib/entproto"
	runtime "entgo.io/contrib/entproto/runtime"
	sqlgraph "entgo.io/ent/dialect/sql/sqlgraph"
	fmt "fmt"
	empty "github.com/golang/protobuf/ptypes/empty"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	strconv "strconv"
)

// ChatService implements ChatServiceServer
type ChatService struct {
	client *ent.Client
	UnimplementedChatServiceServer
}

// NewChatService returns a new ChatService
func NewChatService(client *ent.Client) *ChatService {
	return &ChatService{
		client: client,
	}
}

// toProtoChat transforms the ent type to the pb type
func toProtoChat(e *ent.Chat) (*Chat, error) {
	v := &Chat{}
	id := int64(e.ID)
	v.Id = id
	message := e.Message
	v.Message = message
	sent_at := timestamppb.New(e.SentAt)
	v.SentAt = sent_at
	if edg := e.Edges.ReceivedUser; edg != nil {
		id := int64(edg.ID)
		v.ReceivedUser = &User{
			Id: id,
		}
	}
	if edg := e.Edges.SentUser; edg != nil {
		id := int64(edg.ID)
		v.SentUser = &User{
			Id: id,
		}
	}
	return v, nil
}

// toProtoChatList transforms a list of ent type to a list of pb type
func toProtoChatList(e []*ent.Chat) ([]*Chat, error) {
	var pbList []*Chat
	for _, entEntity := range e {
		pbEntity, err := toProtoChat(entEntity)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		pbList = append(pbList, pbEntity)
	}
	return pbList, nil
}

// Create implements ChatServiceServer.Create
func (svc *ChatService) Create(ctx context.Context, req *CreateChatRequest) (*Chat, error) {
	chat := req.GetChat()
	m, err := svc.createBuilder(chat)
	if err != nil {
		return nil, err
	}
	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoChat(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return proto, nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Get implements ChatServiceServer.Get
func (svc *ChatService) Get(ctx context.Context, req *GetChatRequest) (*Chat, error) {
	var (
		err error
		get *ent.Chat
	)
	id := int(req.GetId())
	switch req.GetView() {
	case GetChatRequest_VIEW_UNSPECIFIED, GetChatRequest_BASIC:
		get, err = svc.client.Chat.Get(ctx, id)
	case GetChatRequest_WITH_EDGE_IDS:
		get, err = svc.client.Chat.Query().
			Where(chat.ID(id)).
			WithReceivedUser(func(query *ent.UserQuery) {
				query.Select(user.FieldID)
			}).
			WithSentUser(func(query *ent.UserQuery) {
				query.Select(user.FieldID)
			}).
			Only(ctx)
	default:
		return nil, status.Error(codes.InvalidArgument, "invalid argument: unknown view")
	}
	switch {
	case err == nil:
		return toProtoChat(get)
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Update implements ChatServiceServer.Update
func (svc *ChatService) Update(ctx context.Context, req *UpdateChatRequest) (*Chat, error) {
	chat := req.GetChat()
	chatID := int(chat.GetId())
	m := svc.client.Chat.UpdateOneID(chatID)
	chatMessage := chat.GetMessage()
	m.SetMessage(chatMessage)
	if chat.GetReceivedUser() != nil {
		chatReceivedUser := int(chat.GetReceivedUser().GetId())
		m.SetReceivedUserID(chatReceivedUser)
	}
	if chat.GetSentUser() != nil {
		chatSentUser := int(chat.GetSentUser().GetId())
		m.SetSentUserID(chatSentUser)
	}

	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoChat(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return proto, nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Delete implements ChatServiceServer.Delete
func (svc *ChatService) Delete(ctx context.Context, req *DeleteChatRequest) (*empty.Empty, error) {
	var err error
	id := int(req.GetId())
	err = svc.client.Chat.DeleteOneID(id).Exec(ctx)
	switch {
	case err == nil:
		return &emptypb.Empty{}, nil
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// List implements ChatServiceServer.List
func (svc *ChatService) List(ctx context.Context, req *ListChatRequest) (*ListChatResponse, error) {
	var (
		err      error
		entList  []*ent.Chat
		pageSize int
	)
	pageSize = int(req.GetPageSize())
	switch {
	case pageSize < 0:
		return nil, status.Errorf(codes.InvalidArgument, "page size cannot be less than zero")
	case pageSize == 0 || pageSize > entproto.MaxPageSize:
		pageSize = entproto.MaxPageSize
	}
	listQuery := svc.client.Chat.Query().
		Order(ent.Desc(chat.FieldID)).
		Limit(pageSize + 1)
	if req.GetPageToken() != "" {
		bytes, err := base64.StdEncoding.DecodeString(req.PageToken)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "page token is invalid")
		}
		token, err := strconv.ParseInt(string(bytes), 10, 32)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "page token is invalid")
		}
		pageToken := int(token)
		listQuery = listQuery.
			Where(chat.IDLTE(pageToken))
	}
	switch req.GetView() {
	case ListChatRequest_VIEW_UNSPECIFIED, ListChatRequest_BASIC:
		entList, err = listQuery.All(ctx)
	case ListChatRequest_WITH_EDGE_IDS:
		entList, err = listQuery.
			WithReceivedUser(func(query *ent.UserQuery) {
				query.Select(user.FieldID)
			}).
			WithSentUser(func(query *ent.UserQuery) {
				query.Select(user.FieldID)
			}).
			All(ctx)
	}
	switch {
	case err == nil:
		var nextPageToken string
		if len(entList) == pageSize+1 {
			nextPageToken = base64.StdEncoding.EncodeToString(
				[]byte(fmt.Sprintf("%v", entList[len(entList)-1].ID)))
			entList = entList[:len(entList)-1]
		}
		protoList, err := toProtoChatList(entList)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return &ListChatResponse{
			ChatList:      protoList,
			NextPageToken: nextPageToken,
		}, nil
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// BatchCreate implements ChatServiceServer.BatchCreate
func (svc *ChatService) BatchCreate(ctx context.Context, req *BatchCreateChatsRequest) (*BatchCreateChatsResponse, error) {
	requests := req.GetRequests()
	if len(requests) > entproto.MaxBatchCreateSize {
		return nil, status.Errorf(codes.InvalidArgument, "batch size cannot be greater than %d", entproto.MaxBatchCreateSize)
	}
	bulk := make([]*ent.ChatCreate, len(requests))
	for i, req := range requests {
		chat := req.GetChat()
		var err error
		bulk[i], err = svc.createBuilder(chat)
		if err != nil {
			return nil, err
		}
	}
	res, err := svc.client.Chat.CreateBulk(bulk...).Save(ctx)
	switch {
	case err == nil:
		protoList, err := toProtoChatList(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return &BatchCreateChatsResponse{
			Chats: protoList,
		}, nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

func (svc *ChatService) createBuilder(chat *Chat) (*ent.ChatCreate, error) {
	m := svc.client.Chat.Create()
	chatMessage := chat.GetMessage()
	m.SetMessage(chatMessage)
	chatSentAt := runtime.ExtractTime(chat.GetSentAt())
	m.SetSentAt(chatSentAt)
	if chat.GetReceivedUser() != nil {
		chatReceivedUser := int(chat.GetReceivedUser().GetId())
		m.SetReceivedUserID(chatReceivedUser)
	}
	if chat.GetSentUser() != nil {
		chatSentUser := int(chat.GetSentUser().GetId())
		m.SetSentUserID(chatSentUser)
	}
	return m, nil
}
