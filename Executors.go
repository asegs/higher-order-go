package main

import (
	"fmt"
	"math/rand"
	"time"
)

type PooledExecutor struct {
	NodeCount  int
	ReadyQueue chan *ExecutorNode
	AllNodes   []*ExecutorNode
}

type ExecutorNode struct {
	Work   chan func()
	Size   int
	Parent *PooledExecutor
}

func Accept[T any](node *ExecutorNode, operation func() T) CompletableFuture[T] {
	future := CompletableFuture[T]{channel: make(chan bool, 1), Finished: false}

	node.Work <- func() {
		future.Complete(operation())
	}

	return future
}

func NewNode(size int, parent *PooledExecutor) *ExecutorNode {
	node := ExecutorNode{
		Work:   make(chan func(), size),
		Size:   size,
		Parent: parent,
	}
	go func() {
		for {
			select {
			case work := <-node.Work:
				work()
			default:
				parent.ReadyQueue <- &node
			}
		}
	}()

	return &node
}

func NewPooledExecutor(nodeCount int, nodeSize int) *PooledExecutor {
	readyQueue := make(chan *ExecutorNode, nodeCount)
	allNodes := make([]*ExecutorNode, nodeCount)
	executor := PooledExecutor{
		NodeCount:  nodeCount,
		ReadyQueue: readyQueue,
	}
	for i := 0; i < nodeCount; i++ {
		newNode := NewNode(nodeSize, &executor)
		allNodes = append(allNodes, newNode)
		readyQueue <- NewNode(nodeCount, &executor)
	}

	executor.AllNodes = allNodes

	return &executor
}

func Submit[T any](executor *PooledExecutor, operation func() T) CompletableFuture[T] {
	select {
	case node := <-executor.ReadyQueue:
		return Accept(node, operation)
	default:
		node := executor.AllNodes[rand.Intn(executor.NodeCount)]
		return Accept(node, operation)
	}
}

func main() {
	executor := NewPooledExecutor(3, 10)
	f1 := Submit(executor, func() int {
		time.Sleep(time.Millisecond * 1500)
		fmt.Println("Doing future 1")
		return 11
	})

	f2 := Submit(executor, func() int {
		time.Sleep(time.Millisecond * 2500)
		fmt.Println("Doing future 2")
		return 22
	})

	f3 := Submit(executor, func() int {
		time.Sleep(time.Millisecond * 1000)
		fmt.Println("Doing future 3")
		return 31
	})
	fmt.Println("Began execution!")

	fmt.Println(f1.Get())
	fmt.Println(f2.Get())
	fmt.Println(f3.Get())
}
