package main

import (
	"context"
	"log"
	"net"

	pb_ptypes "github.com/golang/protobuf/ptypes"
	pb_empty "github.com/golang/protobuf/ptypes/empty"
	bazel_pb "github.com/smukherj1/bazel_bes/proto"
	pb "google.golang.org/genproto/googleapis/devtools/build/v1"
	"google.golang.org/grpc"
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
			be := &bazel_pb.BuildEvent{}
			if err := pb_ptypes.UnmarshalAny(e.BazelEvent, be); err != nil {
				log.Printf("ERROR: Unable to unmarshall Bazel BuildEvent from BazelEvent: %v\n", err)
			} else {
				log.Println("Got Bazel BuildEvent", be.GetId())
			}
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
	log.Printf("Launching Bazel BES service on endpoint 0.0.0.0%s\n", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
