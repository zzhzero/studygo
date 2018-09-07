package parser

import (
	"net/url"
	"regexp"
	"studygo/BFS/drive"
)

func PrintCityList(contents []byte, userurl string) drive.ParseRequest {

	re := regexp.MustCompile(`<a href="//(www.zhaopin.com/[0-9a-z]+/)">([^<]+)</a>`)
	bytes := re.FindAllSubmatch(contents, -1)

	resule := drive.ParseRequest{}

	i := 0
	for _, byt := range bytes {
		i++

		u := url.Values{}
		cityname := string(byt[2])
		u.Set("cityId", cityname)
		u.Set("kw", "java")
		u.Set("kt", "3")
		temper := `https://fe-api.zhaopin.com/c/i/sou?` + u.Encode()

		resule.Requests = append(resule.Requests, drive.Request{
			Url:        temper,
			ParserFunc: PrintCity,
		})
	}
	return resule

}
func NiuFunc(contents []byte, userurl string) drive.ParseRequest {
	return drive.ParseRequest{}
}
