/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a client for Greeter service.
package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"flag"
	"log"
	"math"
	"strings"
	"time"

	pb "pingpong/pingpong"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	// addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func mean(values [32]int64) float64 {
	var sum float64 = 0
	for i := 1; i < 32; i++ {
		sum += float64(values[i])
	}
	return sum / 32
}

func standardDeviation(values [32]int64, mean float64) float64 {
	var sum float64 = 0
	for i := 1; i < 32; i++ {
		sum += math.Pow(float64(values[i])-mean, 2)
	}
	return math.Sqrt(sum / 31)
}

func confidenceInterval(mean float64, stdDev float64, size float64, confidenceLevel float64) (time.Duration, time.Duration) {
	c := 0.95 * (stdDev / math.Sqrt(size))
	return time.Duration(mean - c), time.Duration(mean + c)
}

func longString() string {
	// Using strings.Builder for efficient string concatenation
	var builder strings.Builder
	for i := 0; i < 10000; i++ {
		builder.WriteString("Hello Go!")
	}
	result := builder.String()
	return result
}

func generateRandomString(n int) (string, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func main() {
	var server string
	flag.StringVar(&server, "server", "localhost", "Specify the server")
	flag.Parse()
	addr := flag.String("addr", server+":50051", "the address to connect to")
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewPingPongClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	var values [32]int64
	for i := 0; i < 32; i++ {
		log.Println("[client] sending ping")
		start := time.Now()
		_, err = c.Pong(ctx, &pb.PingRequest{})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		values[i] = int64(time.Since(start))
		log.Printf("[client] received pong")
		time_received := time.Since(start)
		log.Printf("[client] time elapsed %v", time_received)
	}

	// // Timing messages of different sizes
	// file, err := os.Create("duration_per_size.txt")
	// if err != nil {
	// 	fmt.Println("Error creating file:", err)
	// 	return
	// }
	// defer file.Close()

	// debit_f, err2 := os.Create("debit.txt")
	// if err2 != nil {
	// 	fmt.Println("Error creating file:", err2)
	// 	return
	// }
	// defer debit_f.Close()

	meanVal := mean(values)

	for i := int(math.Pow(10, 3)); i < 3*int(math.Pow(10, 6))+1; i += 30000 {
		log.Println("message_size:", i)
		randomString, err := generateRandomString(i)
		start := time.Now()
		_, err = c.Pong(ctx, &pb.PingRequest{Message: randomString})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		duration := time.Since(start)
		t := int64(duration)
		// _, err2 := file.WriteString(fmt.Sprintf("%d\n", t))
		// if err2 != nil {
		// 	fmt.Println("Error writing to file:", err2)
		// 	return
		// }
		d := float64(i) / (float64(t) - meanVal)
		log.Println("debit:", d)
		// _, err3 := debit_f.WriteString(fmt.Sprintf("%f\n", d))
		// if err3 != nil {
		// 	fmt.Println("Error writing to file:", err3)
		// 	return
		// }
	}

	// Confidence interval computation
	stdDev := standardDeviation(values, meanVal)
	min_ci, max_ci := confidenceInterval(meanVal, stdDev, 31, 0.95)
	log.Println("Confidence interval:", min_ci, max_ci)

	// Debit computation
	var s string = longString()
	start := time.Now()
	_, err = c.Pong(ctx, &pb.PingRequest{Message: s})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	duration := time.Since(start)
	log.Println("duration :", duration)
	t := int64(duration)
	log.Println("duration t:", t)
	size := float64(len(s))
	log.Println("size :", size)
	min_bw, max_bw := size/float64(t-2*int64(min_ci)), size/float64(t-2*int64(max_ci))
	log.Printf("bandwidth (%v, %v) bytes/second ", min_bw*math.Pow(10, 9), max_bw*math.Pow(10, 9))
}
