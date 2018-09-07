package drive

import (
	"log"
	"studygo/BFS/tools/fetcher"
)

type SimpleDrive struct {
}

func Run(seeds ...Request) {

	var requests []Request

	itemcount := 0

	for _, seed := range seeds {
		requests = append(requests, seed)
	}
	for len(requests) > 0 {
		request := requests[0]
		requests = requests[1:]

		parseRequest, err := Worker(request)
		if err != nil {
			continue
		}

		requests = append(requests, parseRequest.Requests...)
		for _, ite := range parseRequest.Item {
			itemcount++
			log.Printf("item%d:%s", itemcount, ite)
		}
	}
}
func Worker(request Request) (ParseRequest, error) {
	all2utf8, err := fetcher.Fetcher(request.Url)

	if err != nil {
		//log.Printf("all2utf8 err %s:%v", request.Url, err)
		return ParseRequest{}, err
	}

	return request.ParserFunc(all2utf8, request.Url), nil
}
