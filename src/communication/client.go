package communication

import (
	"io/ioutil"
	"net/http"
	"strings"

	pb "github.com/damienfamed75/engo-xaro/src/proto"
	"github.com/damienfamed75/engo-xaro/src/report"
)

// Client is used to store all client data to run the game
type Client struct {
	Conn   pb.XaroClient
	Player *pb.Player
}

// NewClient returns a new client that is ready to
// send and receive information to and from a server.
func NewClient(conn pb.XaroClient, username string) *Client {
	return &Client{
		Conn: conn, // the connection information to the server.
		Player: &pb.Player{
			Username: username,
			ID:       0,
			IP:       getPublicIP(),
			Position: &pb.Point{},
		},
	}
}

// UpdatePosition updates the client's position with the
// fed in parameters. (typically a Player game object)
func (c *Client) UpdatePosition(x, y float32) {
	c.Player.Position.X, c.Player.Position.Y = x, y
}

// SetID updates the client's ID using the fed in
// parameter. (typically from a Player game object)
func (c *Client) SetID(id uint32) {
	c.Player.ID = id
}

// GetConn returns a reference pointer to
// the client's connection information.
func (c *Client) GetConn() *pb.XaroClient {
	return &c.Conn
}

// GetPlayer returns a reference pointer to
// the connection client's player information.
func (c *Client) GetPlayer() *pb.Player {
	return c.Player
}

func getPublicIP() string {
	// queries the website for a signal to get their IP address.
	resp, err := http.Get("http://ipv4.myexternalip.com/raw")
	report.Error("failed to query site:", err)

	// defer to close the website's response.
	defer resp.Body.Close()

	// reads the text from the response's body.
	htmlData, err := ioutil.ReadAll(resp.Body)
	report.Error("failed to read html page:", err)

	// converts the body's text to string
	stringData := string(htmlData)

	// trims the new line from the string'd data.
	stringData = strings.TrimSuffix(stringData, "\n")

	return string(htmlData)
}
