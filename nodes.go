package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
  "math/rand"
	"os"
  "time"
)

func main() {
	width, height := 1024, 1024
	canvas := NewCanvas(image.Rect(0, 0, width, height))
	rand.Seed(time.Now().UTC().UnixNano())

  // Create and populate the slice of Nodes
  n := 1000
  peers := 8
  nodes := make([]*Node, n)
  for i := 0; i < n; i++ {
    nodes[i] = NewNode(peers, canvas)
  }

  // Randomly point Nodes at each other
  for _, node := range nodes {
    for _, j := range rand.Perm(n)[:peers] {
      node.Peers = append(node.Peers, nodes[j])
    }
  }

  // Draw connections between nodes
  // for _, node := range nodes {
  //   for _, peer := range node.Peers {
  //     canvas.DrawLine(color.RGBA{0, 0, 0, 255}, node.Position, peer.Position)
  //   }
	// }

  // Start sending messages between nodes
  nodes[0].Ch <- nodes[1]
  //time.Sleep(time.Second / 10)

  // Write out image
	outFilename := "nodes.png"
	outFile, err := os.Create(outFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()
	log.Print("Saving image to: ", outFilename)
	png.Encode(outFile, canvas)
}

type Node struct {
  Position Vector
	Ch chan *Node
  Peers []*Node
  Canvas *Canvas
}

func NewNode(peers int, canvas *Canvas) *Node {
  node := new(Node)
  node.Peers = make([]*Node, 0, peers)
	size := canvas.Bounds().Size()
  x := float64(size.X) * rand.Float64()
  y := float64(size.Y) * rand.Float64()
  node.Position = Vector{x, y}
  node.Canvas = canvas
  node.Ch = make(chan *Node)
  go node.Listen()
  return node
}

func (n *Node) Listen() {
  // Listen for incoming connection on node's channel
  for {
    peer := <-n.Ch

    n.Canvas.DrawLine(color.RGBA{255, 0, 0, 255}, n.Position, peer.Position)
    // Retransmit to random node
    // n.Peers[rand.Intn(len(n.Peers))].Ch <- n
    for _, p := range n.Peers {
      p.Ch <- n
    }
  }
}
