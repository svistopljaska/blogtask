package main

import (
	"context"
	pb "github.com/svistopljaska/blogtask/blogrpc"
	"github.com/svistopljaska/blogtask/data"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net"
)

const (
	port = ":4004"
)

type server struct {
	pb.UnimplementedBlogServer
}

func makePbPost(post *data.BlogPost) *pb.Post {
	var res pb.Post

	i := post.Id
	res = pb.Post{
		Id:        &i,
		Title:     post.Title,
		Content:   post.Content,
		Createdon: timestamppb.New(post.CreatedOn),
		Author:    post.Author,
		Tags:      make([]*pb.Tag, len(post.Tags)),
	}

	for _, t := range post.Tags {
		tagid := t.Id
		res.Tags = append(res.Tags, &pb.Tag{
			Id:   &tagid,
			Name: t.Name,
		})
	}

	return &res
}

func makePost(post *pb.Post) data.BlogPost {
	blogpost := data.BlogPost{
		Title:     post.Title,
		Content:   post.Content,
		Tags:      make([]*data.Tag, len(post.Tags)),
		CreatedOn: post.Createdon.AsTime(),
		Author:    post.Author,
	}

	for _, t := range post.Tags {
		blogpost.Tags = append(blogpost.Tags, &data.Tag{Name: t.Name})
	}

	return blogpost
}

func (s server) GetPosts(ctx context.Context, empty *emptypb.Empty) (*pb.Posts, error) {
	posts, err := data.GetBlogPosts()
	if err != nil {
		return nil, err
	}
	responsePosts := pb.Posts{Posts: make([]*pb.Post, len(posts))}

	for _, p := range posts {
		pbpost := makePbPost(&p)
		responsePosts.Posts = append(responsePosts.Posts, pbpost)
	}

	return &responsePosts, nil
}

func (s server) CreatePost(ctx context.Context, post *pb.Post) (*pb.PostId, error) {
	blogpost := makePost(post)
	err := blogpost.Save()
	if err != nil {
		return nil, err
	}
	return &pb.PostId{Id: blogpost.Id}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	grpcserver := grpc.NewServer()
	pb.RegisterBlogServer(grpcserver, &server{})
	if err := grpcserver.Serve(lis); err != nil {
		panic(err)
	}
}
