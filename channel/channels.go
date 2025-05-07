package main

// sending

// func processNum(numChan chan int) {

// 	for num := range numChan {
// 		fmt.Println("Processing number", num)
// 		time.Sleep(time.Second)
// 	}

// }

func sum(result chan int, num1 int, num2 int) {
	numResult := num1 + num2
	result <- numResult
}

func main() {

	result := make(chan int)

	go sum(result, 4, 5)

	res := <-result

	println(res)

	// numChan := make(chan int)

	// go processNum(numChan)

	// for {
	// 	numChan <- rand.Intn(100)
	// }

	// // craete a channel
	// messageChannel := make(chan string)

	// // Send data in channel

	// messageChannel <- "ping"

	// // recieve data

	// message := <-messageChannel

	// fmt.Println(message)
}
