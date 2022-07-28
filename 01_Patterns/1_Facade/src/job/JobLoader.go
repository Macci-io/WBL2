package job

//Loader грузчик
type Loader struct {
	Job
}

//NewLoader конструктор
func NewLoader() *Loader {
	return &Loader{NewJob("mowing Garbage")}
}
