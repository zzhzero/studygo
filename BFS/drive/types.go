package drive

type Request struct {
	Url        string
	ParserFunc func([]byte, string) ParseRequest
}
type ParseRequest struct {
	Requests []Request
	Item     []interface{}
}
type Company struct {
	CompanyName string
	CompanyUrl  string
	CompanySize string
}
type Job struct {
	C           Company
	City        string
	PositionUrl string
	EmplType    string
	JobName     string
	CreateDate  string
	EnDate      string
	UpdateDate  string
	WorkExp     string
	EduLevel    string
	Salary      string
}
