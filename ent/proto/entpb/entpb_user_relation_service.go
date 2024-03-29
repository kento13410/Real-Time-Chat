// Code generated by protoc-gen-entgrpc. DO NOT EDIT.
package entpb

import (
	ent "Real-Time-Chat/ent"
	user "Real-Time-Chat/ent/user"
	userrelation "Real-Time-Chat/ent/userrelation"
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

// UserRelationService implements UserRelationServiceServer
type UserRelationService struct {
	client *ent.Client
	UnimplementedUserRelationServiceServer
}

// NewUserRelationService returns a new UserRelationService
func NewUserRelationService(client *ent.Client) *UserRelationService {
	return &UserRelationService{
		client: client,
	}
}

// toProtoUserRelation transforms the ent type to the pb type
func toProtoUserRelation(e *ent.UserRelation) (*UserRelation, error) {
	v := &UserRelation{}
	created_at := timestamppb.New(e.CreatedAt)
	v.CreatedAt = created_at
	id := int64(e.ID)
	v.Id = id
	if edg := e.Edges.User1; edg != nil {
		id := int64(edg.ID)
		v.User1 = &User{
			Id: id,
		}
	}
	if edg := e.Edges.User2; edg != nil {
		id := int64(edg.ID)
		v.User2 = &User{
			Id: id,
		}
	}
	return v, nil
}

// toProtoUserRelationList transforms a list of ent type to a list of pb type
func toProtoUserRelationList(e []*ent.UserRelation) ([]*UserRelation, error) {
	var pbList []*UserRelation
	for _, entEntity := range e {
		pbEntity, err := toProtoUserRelation(entEntity)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		pbList = append(pbList, pbEntity)
	}
	return pbList, nil
}

// Create implements UserRelationServiceServer.Create
func (svc *UserRelationService) Create(ctx context.Context, req *CreateUserRelationRequest) (*UserRelation, error) {
	userrelation := req.GetUserRelation()
	m, err := svc.createBuilder(userrelation)
	if err != nil {
		return nil, err
	}
	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoUserRelation(res)
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

// Get implements UserRelationServiceServer.Get
func (svc *UserRelationService) Get(ctx context.Context, req *GetUserRelationRequest) (*UserRelation, error) {
	var (
		err error
		get *ent.UserRelation
	)
	id := int(req.GetId())
	switch req.GetView() {
	case GetUserRelationRequest_VIEW_UNSPECIFIED, GetUserRelationRequest_BASIC:
		get, err = svc.client.UserRelation.Get(ctx, id)
	case GetUserRelationRequest_WITH_EDGE_IDS:
		get, err = svc.client.UserRelation.Query().
			Where(userrelation.ID(id)).
			WithUser1(func(query *ent.UserQuery) {
				query.Select(user.FieldID)
			}).
			WithUser2(func(query *ent.UserQuery) {
				query.Select(user.FieldID)
			}).
			Only(ctx)
	default:
		return nil, status.Error(codes.InvalidArgument, "invalid argument: unknown view")
	}
	switch {
	case err == nil:
		return toProtoUserRelation(get)
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Update implements UserRelationServiceServer.Update
func (svc *UserRelationService) Update(ctx context.Context, req *UpdateUserRelationRequest) (*UserRelation, error) {
	userrelation := req.GetUserRelation()
	userrelationID := int(userrelation.GetId())
	m := svc.client.UserRelation.UpdateOneID(userrelationID)
	if userrelation.GetUser1() != nil {
		userrelationUser1 := int(userrelation.GetUser1().GetId())
		m.SetUser1ID(userrelationUser1)
	}
	if userrelation.GetUser2() != nil {
		userrelationUser2 := int(userrelation.GetUser2().GetId())
		m.SetUser2ID(userrelationUser2)
	}

	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoUserRelation(res)
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

// Delete implements UserRelationServiceServer.Delete
func (svc *UserRelationService) Delete(ctx context.Context, req *DeleteUserRelationRequest) (*empty.Empty, error) {
	var err error
	id := int(req.GetId())
	err = svc.client.UserRelation.DeleteOneID(id).Exec(ctx)
	switch {
	case err == nil:
		return &emptypb.Empty{}, nil
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// List implements UserRelationServiceServer.List
func (svc *UserRelationService) List(ctx context.Context, req *ListUserRelationRequest) (*ListUserRelationResponse, error) {
	var (
		err      error
		entList  []*ent.UserRelation
		pageSize int
	)
	pageSize = int(req.GetPageSize())
	switch {
	case pageSize < 0:
		return nil, status.Errorf(codes.InvalidArgument, "page size cannot be less than zero")
	case pageSize == 0 || pageSize > entproto.MaxPageSize:
		pageSize = entproto.MaxPageSize
	}
	listQuery := svc.client.UserRelation.Query().
		Order(ent.Desc(userrelation.FieldID)).
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
			Where(userrelation.IDLTE(pageToken))
	}
	switch req.GetView() {
	case ListUserRelationRequest_VIEW_UNSPECIFIED, ListUserRelationRequest_BASIC:
		entList, err = listQuery.All(ctx)
	case ListUserRelationRequest_WITH_EDGE_IDS:
		entList, err = listQuery.
			WithUser1(func(query *ent.UserQuery) {
				query.Select(user.FieldID)
			}).
			WithUser2(func(query *ent.UserQuery) {
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
		protoList, err := toProtoUserRelationList(entList)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return &ListUserRelationResponse{
			UserRelationList: protoList,
			NextPageToken:    nextPageToken,
		}, nil
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// BatchCreate implements UserRelationServiceServer.BatchCreate
func (svc *UserRelationService) BatchCreate(ctx context.Context, req *BatchCreateUserRelationsRequest) (*BatchCreateUserRelationsResponse, error) {
	requests := req.GetRequests()
	if len(requests) > entproto.MaxBatchCreateSize {
		return nil, status.Errorf(codes.InvalidArgument, "batch size cannot be greater than %d", entproto.MaxBatchCreateSize)
	}
	bulk := make([]*ent.UserRelationCreate, len(requests))
	for i, req := range requests {
		userrelation := req.GetUserRelation()
		var err error
		bulk[i], err = svc.createBuilder(userrelation)
		if err != nil {
			return nil, err
		}
	}
	res, err := svc.client.UserRelation.CreateBulk(bulk...).Save(ctx)
	switch {
	case err == nil:
		protoList, err := toProtoUserRelationList(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return &BatchCreateUserRelationsResponse{
			UserRelations: protoList,
		}, nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

func (svc *UserRelationService) createBuilder(userrelation *UserRelation) (*ent.UserRelationCreate, error) {
	m := svc.client.UserRelation.Create()
	userrelationCreatedAt := runtime.ExtractTime(userrelation.GetCreatedAt())
	m.SetCreatedAt(userrelationCreatedAt)
	if userrelation.GetUser1() != nil {
		userrelationUser1 := int(userrelation.GetUser1().GetId())
		m.SetUser1ID(userrelationUser1)
	}
	if userrelation.GetUser2() != nil {
		userrelationUser2 := int(userrelation.GetUser2().GetId())
		m.SetUser2ID(userrelationUser2)
	}
	return m, nil
}
