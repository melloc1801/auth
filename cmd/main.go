package main

import (
	desc "auth/pkg/user_v1"
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"
	"time"
)

const grpcPort = 50051
const dbDSN = "host=localhost port=54321 dbname=auth-service user=dev-course password=1801 sslmode=disable"

type server struct {
	desc.UnimplementedUserV1Server
}

func (s *server) Create(ctx context.Context, req *desc.CreateUserRequest) (*desc.CreateUserResponse, error) {
	if req.User.Password != req.User.PasswordConfirm {
		return nil, errors.New("Passwords are not equal")
	}

	pool, err := pgx.Connect(ctx, dbDSN)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to connect to database %s", err)
	}

	insertBuilder := squirrel.Insert("\"user\"").
		PlaceholderFormat(squirrel.Dollar).
		Columns("name", "email", "role", "password").
		Values(req.User.Name, req.User.Email, req.User.Role.String(), req.User.Password).
		Suffix("RETURNING id")

	query, args, err := insertBuilder.ToSql()
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to build query %s", err)
	}

	var id int64
	err = pool.QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to make query: %s", err)
	}

	return &desc.CreateUserResponse{
		Id: id,
	}, nil
}

func (s *server) Get(ctx context.Context, req *desc.GetUserRequest) (*desc.GetUserResponse, error) {
	pool, err := pgx.Connect(ctx, dbDSN)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to connect to database %s", err)
	}

	selectBuilder := squirrel.Select("id", "name", "email", "role", "created_at", "updated_at").
		PlaceholderFormat(squirrel.Dollar).
		From("\"user\"").
		Where(squirrel.Eq{"id": req.Id})

	query, args, err := selectBuilder.ToSql()
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to build query %s", err)
	}

	var userId int64
	var userName string
	var email string
	var role string
	var createdAt time.Time
	var updatedAt *time.Time
	err = pool.QueryRow(ctx, query, args...).Scan(&userId, &userName, &email, &role, &createdAt, &updatedAt)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to make query: %s", err)
	}

	user := &desc.GetUserResponse{
		Id:        userId,
		Name:      userName,
		Email:     email,
		Role:      desc.Role(desc.Role_value[role]),
		CreatedAt: timestamppb.New(createdAt),
	}
	if updatedAt != nil {
		user.UpdatedAt = timestamppb.New(*updatedAt)
	}

	return user, nil
}

func (s *server) Update(ctx context.Context, req *desc.UpdateUserRequest) (*empty.Empty, error) {
	if req.Email == nil && req.Name == nil && req.Role.Number() == 0 {
		return &empty.Empty{}, nil
	}

	pool, err := pgx.Connect(ctx, dbDSN)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to connect to database %s", err)
	}

	builderUpdate := squirrel.Update("\"user\"").
		PlaceholderFormat(squirrel.Dollar).
		Set("updated_at", time.Now())
	if req.Email != nil {
		builderUpdate = builderUpdate.Set("email", req.Email.Value)
	}
	if req.Name != nil {
		builderUpdate = builderUpdate.Set("name", req.Name.Value)
	}
	if req.Role.Number() != 0 {
		builderUpdate = builderUpdate.Set("role", req.Role.String())
	}
	builderUpdate = builderUpdate.Where(squirrel.Eq{"id": req.Id})

	query, args, err := builderUpdate.ToSql()
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to build query: %s", err)
	}

	_, err = pool.Exec(ctx, query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to executed query: %s", err)
	}

	return &empty.Empty{}, nil
}

func (s *server) Delete(ctx context.Context, req *desc.DeleteUserRequest) (*empty.Empty, error) {
	pool, err := pgx.Connect(ctx, dbDSN)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to connect to database %s", err)
	}

	deleteBuilder := squirrel.Delete("\"user\"").
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{"id": req.Id})

	query, args, err := deleteBuilder.ToSql()
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to build query %s", err)
	}
	_, err = pool.Exec(ctx, query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to execute %s", err)
	}

	return &empty.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserV1Server(s, &server{})

	fmt.Println("Server has been started")
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
