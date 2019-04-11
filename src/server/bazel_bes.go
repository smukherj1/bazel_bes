package main

import (
	"context"
	pb_empty "github.com/golang/protobuf/ptypes/empty"
	pb "google.golang.org/genproto/googleapis/devtools/build/v1"
	//bazel_pb "bazel_bes/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":8080"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

func (s *server) PublishLifecycleEvent(ctx context.Context, req *pb.PublishLifecycleEventRequest) (*pb_empty.Empty, error) {
	log.Println("Got life cycle event for build", req.BuildEvent.StreamId.BuildId, req.BuildEvent)
	return &pb_empty.Empty{}, nil
}

func (s *server) PublishBuildToolEventStream(stream pb.PublishBuildEvent_PublishBuildToolEventStreamServer) error {
	log.Println("Reading build tool event stream")
	for {
		req, _ := stream.Recv()
		if req == nil || req.OrderedBuildEvent == nil || req.OrderedBuildEvent.Event == nil {
			break
		}
		switch e := req.OrderedBuildEvent.Event.Event.(type) {
		case *pb.BuildEvent_InvocationAttemptStarted_:
			log.Println("Got InvocationAttemptStarted event, attempt", e.InvocationAttemptStarted.AttemptNumber)
		case *pb.BuildEvent_InvocationAttemptFinished_:
			log.Println("Got InvocationAttemptFinished event, result", e.InvocationAttemptFinished.GetInvocationStatus().GetResult().String())
		case *pb.BuildEvent_BazelEvent:
			log.Println("Got BazelEvent", e.BazelEvent)
		}
		resp := &pb.PublishBuildToolEventStreamResponse{StreamId: req.OrderedBuildEvent.StreamId, SequenceNumber: req.OrderedBuildEvent.SequenceNumber}
		stream.Send(resp)
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPublishBuildEventServer(s, &server{})
	log.Println("Launching Bazel BES service on", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
