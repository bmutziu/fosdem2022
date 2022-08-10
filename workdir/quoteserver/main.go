package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os/exec"

	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/load"
	"google.golang.org/grpc"

	pb "acme.com/x/quote"
)

var port = flag.Int("port", 3000, "The server port")

func main() {
	log.SetFlags(0)
	log.Println("in quoteserver")
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterQuoterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type server struct {
	pb.UnimplementedQuoterServer
}

func (s *server) Quote(ctx context.Context, in *pb.QuoteRequest) (*pb.QuoteResponse, error) {
	log.Printf("Received: %#v", in)

	num := in.GetNum()
	log.Printf("Num: %v", num)

	numPart := fmt.Sprintf("%d", num)

	cmd := exec.Command("./validate.sh", numPart)

	out, err := cmd.Output()
	if err != nil {
		println(err.Error())
	}
	print(string(out))

	instances := load.Instances([]string{}, &load.Config{Dir: "."})
	for _, i := range instances {
		ctx := cuecontext.New()
		v := ctx.BuildInstance(i)
		if v.Err() != nil {
			// log.Fatalf("%s", v.Err())
			log.Println(v.Err())
		}
		fmt.Printf("%s\n", v)
	}

	pick := []string{}
	if num > 2 {
		pick = append(pick, "Only two quotes as maximum")
	} else {

		qts := []string{"Concurrency is not parallelism.", "Errors are values.", "Reflection is never clear."}

		m := map[string]bool{}
		for ok := true; ok; ok = !(len(pick) == int(num)) {
			randomIndex := rand.Intn(len(qts))
			if !m[qts[randomIndex]] {
				m[qts[randomIndex]] = true
				pick = append(pick, qts[randomIndex])
			}
		}
	}
	log.Printf("Pick: %+v", pick)

	return &pb.QuoteResponse{Quotes: pick}, nil
}
