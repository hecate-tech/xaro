package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/damienfamed75/engo-xaro/src/communication"
	pb "github.com/damienfamed75/engo-xaro/src/proto"

	"google.golang.org/grpc"
)

const (
	address     = "98.144.164.154:8081"
	defaultName = "Imposter"
)

var (
	// Conn is the connection
	Conn *grpc.ClientConn
)

func main() {
	Conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	//defer closeConn()
	c := pb.NewXaroClient(Conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	client := communication.NewClient(c, name)

	r, _ := c.UserJoined(ctx, client.Player)
	log.Printf("Join Message: %s", r.Message)

	cancel()
	for {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		client.Player.Position.X++

		r, _ := c.SendPlayerData(ctx, client.Player)
		for _, p := range r.Players {
			log.Printf("Player %v: {%v, %v}", p.Username, p.Position.GetX(), p.Position.GetY())
		}

		time.Sleep(1 * time.Second)
		cancel()
		if client.Player.Position.X > 20 {
			break
		}
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	c.UserLeft(ctx, client.Player)

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
