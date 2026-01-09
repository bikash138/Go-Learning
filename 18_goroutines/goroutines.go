package main

import (
	"fmt"
	"sync"
)

func task(id int, w *sync.WaitGroup) {
	//Defer runs after the function ends
	defer w.Done()
	fmt.Println("Doing task", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		//Wait group will increament by 1
		wg.Add(1)
		go task(i, &wg)
	}

	//It will until the waitgroup is done
	wg.Wait()

}