package main
import (
	"flag"
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"

)

func turn(x []int,time time.Time, id int, size int){
	sort.Ints(x)
	fmt.Println( id,time,x[0],x[size/2],x[size -1] )
	return
}

func array_init(size int) []int{
	x := make([]int, size)
	for i := 0; i < size; i++ {
		rand.Seed(time.Now().UnixNano())
		x[i] = rand.Intn(1000)
	}
	return x
}

func rut_init(size int, iter int, id int){
	var wg sync.WaitGroup
	wg.Add(iter)
	for i := 0; i < iter; i++{
		//fmt.Println("iter",i, "id", id)
		time := time.Now()
		go turn(array_init(size),time,id,size)
	}
	wg.Wait()
	return
}


func main(){
	writersPtr := flag.Int("writers", 1, "an int")
	arrPtr := flag.Int("arr-size", 1, "an int")
	iterPtr := flag.Int("iter-count", 1, "an int")
	flag.Parse()
	for i:= 0; i< *writersPtr; i++ {
		go rut_init(*arrPtr, *iterPtr,i)
	}
	fmt.Scanln()
}
