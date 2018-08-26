package communication

import (
	"io/ioutil"
	"net/http"

	"github.com/damienfamed75/engo-xaro/src/common"
	pb "github.com/damienfamed75/engo-xaro/src/proto"
)

// Client is used to store all client data to run the game
type Client struct {
	Conn   pb.XaroClient
	Player *pb.Player
}

// NewClient returns a new client to send and receive data from server
func NewClient(conn pb.XaroClient, username string) *Client {
	return &Client{
		Conn: conn,
		Player: &pb.Player{
			Username: username,
			ID:       0,
			IP:       getPublicIP(),
			Position: &pb.Point{},
		},
	}
}

// UpdatePosition updates the client's player position
func (c *Client) UpdatePosition(x, y float32) {
	c.Player.Position.X, c.Player.Position.Y = x, y
}

// SetID Updates the client's player ID
func (c *Client) SetID(id uint32) {
	c.Player.ID = id
}

// GetConn returns reference to client's connection
func (c *Client) GetConn() *pb.XaroClient {
	return &c.Conn
}

// GetPlayer returns reference to client's player
func (c *Client) GetPlayer() *pb.Player {
	return c.Player
}

func getPublicIP() string {
	resp, err := http.Get("http://ipv4.myexternalip.com/raw")
	common.ErrorCheck("failed to fetch ip:", err)

	defer resp.Body.Close()

	htmlData, err := ioutil.ReadAll(resp.Body)
	common.ErrorCheck("failed to read html page:", err)

	return string(htmlData)
}
