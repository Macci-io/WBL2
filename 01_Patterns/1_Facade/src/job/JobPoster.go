package job

//Poster работа почтальон
type Poster struct {
	Job
}

//NewPoster конструктор
func NewPoster() *Poster {
	return &Poster{NewJob("Sends mail")}
}
