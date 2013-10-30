package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
  "math/rand"
	"os"
  "sort"
  "time"
)

func main() {
	width, height := 1024, 1024
	canvas := NewCanvas(image.Rect(0, 0, width, height))
	rand.Seed(time.Now().UTC().UnixNano())

  // Create and populate the slice of Nodes
  n := 1000
  peers := 4
  nodes := make([]*Node, n)
  for i := 0; i < n; i++ {
    nodes[i] = NewNode(peers, canvas)
  }

  // Randomly point Nodes at each other
  // for _, node := range nodes {
  //   for _, j := range rand.Perm(n)[:peers] {
  //     node.Peers = append(node.Peers, nodes[j])
  //   }
  // }

  nodesCopy := make([]*Node, n)
  copy(nodesCopy, nodes)
  for _, node := range nodes {
    // Sort the nodes by distance
    sorter := NodeSorter{nodesCopy, node}
    sort.Sort(sorter)
    node.Peers = append(node.Peers, nodesCopy[1:peers+1]...)
  }

  // Draw connections between nodes
  for _, node := range nodes {
    for _, peer := range node.Peers {
      canvas.DrawLine(color.RGBA{0, 0, 0, 60}, node.Position, peer.Position)
    }
	}

  // Start sending messages between nodes
  for i := 0; i < 10; i++ {
    nodes[i].Power = 255
    nodes[i].Ch <- nodes[i].Peers[0]
  }
  time.Sleep(time.Second)

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
  Power uint8
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
  node.Power = 0
  go node.Listen()
  return node
}

func (n *Node) Listen() {
  // Listen for incoming connection on node's channel
  for {
    peer := <-n.Ch
    peer.Power -= 5
    n.Power = peer.Power
    n.Canvas.DrawLine(color.RGBA{255, n.Power, 0, 255}, n.Position, peer.Position)
    // Retransmit to random node
    if n.Power > 0 {
      go n.Send()
    }
  }
}

func (n *Node) Send() {
  for _, target := range n.Peers {
    if target.Power == 0 {
      target.Ch <- n
      break
    }
  }
}

type NodeSorter struct {
  data []*Node
  target *Node
}

func (sorter NodeSorter) Len() int { return len(sorter.data) }
func (sorter NodeSorter) Less(i, j int) bool {
  iDelta := sorter.data[i].Position.Sub(sorter.target.Position)
  jDelta := sorter.data[j].Position.Sub(sorter.target.Position)
  return iDelta.Length() < jDelta.Length()
}
func (sorter NodeSorter) Swap(i, j int) {
  sorter.data[i], sorter.data[j] = sorter.data[j], sorter.data[i]
}
