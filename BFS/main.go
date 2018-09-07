package main

import (
	"studygo/BFS/drive"
	"studygo/BFS/elasticsearch"
	"studygo/BFS/scheduler"
	"studygo/BFS/zhilian/parser"
)

func main() {
	e := drive.ConcurentDrive{
		Scheduler: &scheduler.SimpleScheduler{},
		WorkCount: 10,
		ItemChan:  elasticsearch.ItemSaver(),
	}
	e.Run(drive.Request{
		Url:        "https://www.zhaopin.com/",
		ParserFunc: parser.PrintCityList,
	})
}
