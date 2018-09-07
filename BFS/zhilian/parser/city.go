package parser

import (
	"encoding/json"
	"net/url"
	"strconv"
	"strings"
	"studygo/BFS/drive"
)

func PrintCity(contents []byte, userurl string) drive.ParseRequest {

	resule := drive.ParseRequest{}
	var f interface{}
	json.Unmarshal(contents, &f)
	m := f.(map[string]interface{})

	p := m["data"].(map[string]interface{})["results"].([]interface{})
	i := 0
	numFound := int(m["data"].(map[string]interface{})["numFound"].(float64))

	index := strings.Index(userurl, "&p=")
	nowp, _ := strconv.Atoi(userurl[index+3:])
	citycode := ""

	for _, vvv := range p {
		//一个vvv对应一份工作
		i++
		companyem := drive.Company{}
		job := drive.Job{
			C: companyem,
		}
		q := vvv.(map[string]interface{})

		job.JobName = q["jobName"].(string)         //工作
		job.Salary = q["jobName"].(string)          //工资
		job.PositionUrl = q["positionURL"].(string) //url
		job.CreateDate = q["createDate"].(string)
		job.EnDate = q["endDate"].(string)
		job.EmplType = q["emplType"].(string)
		job.UpdateDate = q["updateDate"].(string)
		citycode = q["city"].(map[string]interface{})["display"].(string)

		r := q["company"].(map[string]interface{})
		job.City = citycode
		job.C.CompanyName = r["name"].(string)
		job.C.CompanyUrl = r["url"].(string)
		companysizes := r["size"].(map[string]interface{})
		job.C.CompanySize = companysizes["name"].(string)

		eduLevel := q["eduLevel"].(map[string]interface{})
		job.EduLevel = eduLevel["name"].(string) //学历

		workingExp := q["workingExp"].(map[string]interface{})
		job.WorkExp = workingExp["name"].(string) //经验

		resule.Requests = append(resule.Requests, drive.Request{
			Url:        "",
			ParserFunc: NiuFunc,
		})
		resule.Item = append(resule.Item, job)
	}
	if nowp == numFound/60 {
		nowp++
		u := url.Values{}
		cityname := string(citycode)
		u.Set("cityId", cityname)
		u.Set("kw", "java")
		u.Set("kt", "3")
		u.Set("p", strconv.Itoa(nowp))
		temper := `https://fe-api.zhaopin.com/c/i/sou?` + u.Encode()
		//fmt.Printf("%s",userurl)
		resule.Requests = append(resule.Requests, drive.Request{
			Url:        temper,
			ParserFunc: PrintCity,
		})
	}
	return resule

}
