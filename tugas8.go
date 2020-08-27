package main

import "fmt"
import "runtime"
import "math/rand"
import "time"

func main() {
  runtime.GOMAXPROCS(2)
  var time = time.Now().Unix()
  fmt.Println(time)
  rand.Seed(time)

  var pesan = make(chan int)
  go kirimPesan(pesan)
  terimaPesan(pesan)
}

func kirimPesan(ch chan int){
  for i := 0 ; true ; i++ {
    var duration = time.Duration(rand.Int()%10+1) * time.Second
    fmt.Println("Waktu yang diberikan:",duration)
    ch <- i
    time.Sleep(duration)
  }
}

func terimaPesan(ch chan int){
  loop:
  for {
    select {
    case <- ch :
      fmt.Println("Apa Kabar Teman Teman")
    case <- time.After(time.Second * 5) :
      fmt.Println("Timeout, tidak ada aktifitas dalam 5 detik")
      break loop
    }
  }
}
