package store

import (
	"context"
	"fmt"

	pb "github.com/authzed/authzed-go/proto/authzed/api/v1"
	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
	authzed "github.com/authzed/authzed-go/v1"
	"github.com/authzed/grpcutil"
	"github.com/manaty226/sample-app-with-spicedb/internal/entity"
	"google.golang.org/grpc"
)

type spiceClient struct {
	client   *authzed.Client
	zedToken *v1.ZedToken
}

func NewSpiceDB(host string, port int) (*spiceClient, error) {
	client, err := authzed.NewClient(
		fmt.Sprintf("%s:%d", host, port),
		grpc.WithInsecure(),
		grpcutil.WithInsecureBearerToken("foobar"),
	)
	if err != nil {
		return nil, err
	}
	return &spiceClient{
		client: client,
	}, nil
}

func (s spiceClient) CreateNameSpace(schema string) error {
	request := &pb.WriteSchemaRequest{Schema: schema}
	_, err := s.client.WriteSchema(context.Background(), request)
	if err != nil {
		return err
	}
	return nil
}

func (s spiceClient) CheckPermission(object, subject entity.Object, action entity.Action) (bool, error) {
	obj := &pb.ObjectReference{ObjectType: string(object.Type), ObjectId: object.ID}
	sub := &pb.SubjectReference{Object: &pb.ObjectReference{ObjectType: string(subject.Type), ObjectId: subject.ID}}
	resp, err := s.client.CheckPermission(context.Background(), &pb.CheckPermissionRequest{
		Resource:   obj,
		Permission: string(action),
		Subject:    sub,
		Consistency: &v1.Consistency{
			Requirement: &v1.Consistency_AtLeastAsFresh{AtLeastAsFresh: s.zedToken},
		},
	})
	if err != nil {
		return false, err
	}
	isPermitted := resp.GetPermissionship() == pb.CheckPermissionResponse_PERMISSIONSHIP_HAS_PERMISSION
	return isPermitted, nil
}

func (s *spiceClient) CreateRelation(object, subject entity.Object, relation string) error {
	obj := &pb.ObjectReference{ObjectType: string(object.Type), ObjectId: object.ID}
	sub := &pb.SubjectReference{Object: &pb.ObjectReference{ObjectType: string(subject.Type), ObjectId: subject.ID}}
	request := &pb.WriteRelationshipsRequest{Updates: []*pb.RelationshipUpdate{
		{
			Operation: pb.RelationshipUpdate_OPERATION_CREATE,
			Relationship: &pb.Relationship{
				Resource: obj,
				Relation: relation,
				Subject:  sub,
			},
		},
	},
	}
	resp, err := s.client.WriteRelationships(context.Background(), request)
	if err != nil {
		return err
	}
	s.zedToken = resp.WrittenAt
	return nil
}
