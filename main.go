package main

import (
	"fmt"
	cnfg "tskscheduler/task-scheduler/config"
	manag "tskscheduler/task-scheduler/scheduler"
	storage "tskscheduler/task-scheduler/storage"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("error load env %v\n", err)
	}

	done := make(chan int)
	consumer, err := storage.NewConsumer(done)
	if err != nil {
		fmt.Printf("amq connection error %v\n", err)
	} else {
		consumer.SetupCloseHandler()
	}
	producer, err := storage.NewProducer(done)
	if err != nil {
		fmt.Printf("amq connection error %v\n", err)
	} else {
		producer.SetupCloseHandler()
	}
	supa, error := storage.NewSupabaseClient()
	if error != nil {
		fmt.Printf("supabase cloient failed %v\n", error)
	}
	taskM := manag.InitManager(consumer, producer, supa, done, cnfg.LoadConfig())
	taskM.StartManager()
	<-done

}
