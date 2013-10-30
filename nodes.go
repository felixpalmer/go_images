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
  n := 100
  peers := 3
  nodes := make([]*Node, n)
  for i := 0; i < n; i++ {
    nodes[i] = new(Node)
    nodes[i].Peers = make([]*Node, 0, peers)
		x := float64(width) * rand.Float64()
		y := float64(height) * rand.Float64()
    nodes[i].Position = Vector{x, y}
  }

  // Randomly point Nodes at each other
  for _, node := range nodes {
    for _, j := range rand.Perm(n)[:peers] {
      node.Peers = append(node.Peers, nodes[j])
    }
  }

  // Draw connections between nodes
  for _, node := range nodes {
    for _, peer := range node.Peers {
      canvas.DrawLine(color.RGBA{0, 0, 0, 255}, node.Position, peer.Position)
    }
	}

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
}
