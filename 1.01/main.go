package main

// 1
import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/google/uuid"
)

// 2
func randomPrint(c string) {
	fmt.Println(time.Now().UTC().Format(time.RFC3339) + ": " + c)
}

func runCronJobs(c string) {
	// 3
	s := gocron.NewScheduler(time.UTC)

	// 4
	s.Every(5).Seconds().Do(func() {
		randomPrint(c)
	})

	// 5
	s.StartBlocking()
}

// 6
func main() {

	id := uuid.New()

	// Getting random character
	c := id.String()

	runCronJobs(string(c))
}
