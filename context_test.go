package golang_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()
	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextB, "c", "C")
	contextD := context.WithValue(contextC, "d", "D")
	contextE := context.WithValue(contextD, "e", "E")
	contextF := context.WithValue(contextE, "f", "F")
	contextG := context.WithValue(contextB, "g", "G")
	contextH := context.WithValue(contextC, "h", "H")
	contextI := context.WithValue(contextD, "i", "I")
	contextJ := context.WithValue(contextE, "j", "J")
	fmt.Println("context a:", contextA)
	fmt.Println("context b:", contextB)
	fmt.Println("context c:", contextC)
	fmt.Println("context d:", contextD)
	fmt.Println("context e:", contextE)
	fmt.Println("context f:", contextF)
	fmt.Println("context g:", contextG)
	fmt.Println("context h:", contextH)
	fmt.Println("context i:", contextI)
	fmt.Println("context j:", contextJ)

	fmt.Println(contextF.Value("b"))
	fmt.Println(contextG.Value("c"))
	fmt.Println(contextI.Value("e"))
	fmt.Println(contextJ.Value("j"))
}

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second)
			}
		}
	}()
	return destination
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounter(ctx)
	fmt.Println("Total Goroutine", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println("counter :", n)
		if n == 10 {
			break
		}
	}
	cancel()
	time.Sleep(2 * time.Second)
	fmt.Println("Total Goroutine", runtime.NumGoroutine())
}

func TestContextTimeout(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	destination := CreateCounter(ctx)
	fmt.Println("Total Goroutine", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println("counter :", n)
	}
	time.Sleep(2 * time.Second)
	fmt.Println("Total Goroutine", runtime.NumGoroutine())
}

func TestContextWithDeadline(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
	defer cancel()

	destination := CreateCounter(ctx)
	fmt.Println("Total Goroutine", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println("counter :", n)
	}
	time.Sleep(2 * time.Second)
	fmt.Println("Total Goroutine", runtime.NumGoroutine())
}
