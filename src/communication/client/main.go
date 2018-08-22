package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/damienfamed75/engo-xaro/src/proto"

	"google.golang.org/grpc"
)

const (
	address     = "98.144.164.154:8081"
	defaultName = "Damien"
)

var (
	// Conn is the connection
	Conn       *grpc.ClientConn
	testPlayer = &pb.Player{
		ID:       4321,
		IP:       "153.164.23.673",
		Position: &pb.Point{X: 1, Y: 1},
	}
)

func main() {
	Conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	//defer closeConn()
	c := pb.NewGreeterClient(Conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	testPlayer.Username = name
	r, _ := c.UserJoined(ctx, testPlayer)
	log.Printf("Join Message: %s", r.Message)

	for {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		testPlayer.Position.X++
		log.Println("tp:", testPlayer.Position.X)

		r, _ := c.SendPositions(ctx, testPlayer)
		log.Printf("%v", r.Positions)
		time.Sleep(2 * time.Second)
		cancel()
	}
	/* Pseudo code
	for {
		send the server my player's global position (x,y)

		receive my signal back along with any other client's
		messages to the server.

		Draw my received messages that's not from me. (Other clients)
		or
		print out my received messages from server.

		if I stop receiving messages from the server
		then break from this for loop and quit the scene.
	}
	*/
}
