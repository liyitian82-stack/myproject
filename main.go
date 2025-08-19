package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ch := make(chan int)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("程序挂起中，按 Ctrl+C 退出...")

	go func() {
		i := 0
		for {
			i++
			ch <- i
			time.Sleep(time.Second * 1)
		}
	}()

	for {
		select {
		case data := <-ch:
			fmt.Println(fmt.Sprintf("data:%d", data))
		case sig := <-sigs:
			fmt.Printf("收到信号 %v，退出程序\n", sig)
			os.Exit(0)
		}
	}
}
