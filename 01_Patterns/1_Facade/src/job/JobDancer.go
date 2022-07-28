package job

//Dancer работа где танцуют
type Dancer struct {
	Job
}

//NewDancer конструктор
func NewDancer() *Dancer {
	return &Dancer{NewJob("Dancing")}
}
