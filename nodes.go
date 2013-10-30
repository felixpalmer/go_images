package main

import (
	"image"
	"image/png"
	"log"
  "math/rand"
	"os"
)

func main() {
	width, height := 512, 512
	canvas := NewCanvas(image.Rect(0, 0, width, height))

  // Create and populate the slice of Nodes
  n := 10
  peers := 5
  nodes := make([]*Node, n)
  for i := 0; i < n; i++ {
    nodes[i] = new(Node)
    nodes[i].Peers = make([]*Node, 0, peers)
  }

  // Randomly point Nodes at each other
  for _, node := range nodes {
    for _, j := range rand.Perm(n)[:peers] {
      node.Peers = append(node.Peers, nodes[j])
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
	Ch chan *Node
  Peers []*Node
}
