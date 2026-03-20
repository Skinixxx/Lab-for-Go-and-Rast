package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"

	"google.golang.org/grpc"

	pb "github.com/hard/go-service/proto"
)

type IntegratorServer struct {
	pb.UnimplementedIntegratorServer
}

func (s *IntegratorServer) Integrate(ctx context.Context, req *pb.IntegrationRequest) (*pb.IntegrationResponse, error) {
	result := TrapezoidRule(req.Function, req.LowerBound, req.UpperBound, req.Partitions)
	return &pb.IntegrationResponse{
		Result:         result.Result,
		ErrorEstimate:  result.ErrorEstimate,
		PartitionsUsed: result.Partitions,
	}, nil
}

func (s *IntegratorServer) IntegrateStream(req *pb.IntegrationRequest, stream pb.Integrator_IntegrateStreamServer) error {
	partitions := req.Partitions
	step := int64(1000)
	if step > partitions {
		step = partitions
	}

	for i := int64(0); i < partitions; i += step {
		currentPartitions := step
		if i+step > partitions {
			currentPartitions = partitions - i
		}
		result := TrapezoidRule(req.Function, req.LowerBound, req.UpperBound, currentPartitions)
		if err := stream.Send(&pb.IntegrationResponse{
			Result:         result.Result,
			ErrorEstimate:  result.ErrorEstimate,
			PartitionsUsed: result.Partitions,
		}); err != nil {
			return err
		}
		time.Sleep(10 * time.Millisecond)
	}
	return nil
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	funcStr := r.URL.Query().Get("func")
	lowerStr := r.URL.Query().Get("a")
	upperStr := r.URL.Query().Get("b")
	partsStr := r.URL.Query().Get("n")

	lower, _ := strconv.ParseFloat(lowerStr, 64)
	upper, _ := strconv.ParseFloat(upperStr, 64)
	partitions, _ := strconv.ParseInt(partsStr, 10, 64)

	result := TrapezoidRule(funcStr, lower, upper, partitions)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"result": %f, "error_estimate": %e, "partitions": %d}`, result.Result, result.ErrorEstimate, result.Partitions)
}

func StartGRPCServer(port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	pb.RegisterIntegratorServer(s, &IntegratorServer{})
	log.Printf("gRPC server starting on port %d", port)
	return s.Serve(lis)
}

func StartHTTPServer(port int) error {
	http.HandleFunc("/integrate", httpHandler)
	log.Printf("HTTP server starting on port %d", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
