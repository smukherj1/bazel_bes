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

func (s *server) processCompleted(id *bazel_pb.BuildEventId_TargetCompletedId, c *bazel_pb.TargetComplete) error {
	return nil
}

func (s *server) processAction(id *bazel_pb.BuildEventId_ActionCompletedId, a *bazel_pb.ActionExecuted) error {
	return nil
}

func (s *server) processTestResult(id *bazel_pb.BuildEventId_TestResultId, r *bazel_pb.TestResult) error {
	return nil
}

func (s *server) processTestSummary(id *bazel_pb.BuildEventId_TestSummaryId, ts *bazel_pb.TestSummary) error {
	return nil
}

func (s *server) processBuildFinished(f *bazel_pb.BuildFinished) error {
	return nil
}

func (s *server) processBuildAborted(a *bazel_pb.Aborted) error {
	return nil
}

func (s *server) processBuildEvent(be *bazel_pb.BuildEvent) error {
	if c := be.GetCompleted(); c != nil {
		return s.processCompleted(be.GetId().GetTargetCompleted(), c)
	}
	if a := be.GetAction(); a != nil {
		return s.processAction(be.GetId().GetActionCompleted(), a)
	}
	if r := be.GetTestResult(); r != nil {
		return s.processTestResult(be.GetId().GetTestResult(), r)
	}
	if ts := be.GetTestSummary(); ts != nil {
		return s.processTestSummary(be.GetId().GetTestSummary(), ts)
	}
	if f := be.GetFinished(); f != nil {
		return s.processBuildFinished(f)
	}
	if a := be.GetAborted(); a != nil {
		return s.processBuildAborted(a)
	}
	return nil
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
				log.Printf("ERROR: Unable to unmarshall Bazel BuildEvent from BazelEvent: %v", err)
				break
			}
			if err := s.processBuildEvent(be); err != nil {
				log.Printf("ERROR: Unable to process Bazel BuildEvent: %v", err)
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
