//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer scenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"time"
)

func producer(stream Stream, tweetsCh chan<- *Tweet) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			close(tweetsCh)
			return
		}

		tweetsCh <- tweet
	}
}

func consumer(tweetsCh <-chan *Tweet) {
	for t := range tweetsCh {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()

	tweetsCh := make(chan *Tweet)

	// Producer
	go producer(stream, tweetsCh)

	// Consumer
	consumer(tweetsCh)

	fmt.Printf("Process took %s\n", time.Since(start))
}
