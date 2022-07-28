package address

//Builder наш билдер address
type Builder struct {
	address *Address
}

//NewBuilder Создает новый объект билдер
func NewBuilder() *Builder {
	return &Builder{&Address{}}
}

//Build билдит структуру address
func (a *Builder) Build() *Address {
	return a.address
}

//SetCountry сеттер для country
func (a *Builder) SetCountry(country string) *Builder {
	a.address.country = "Country: " + country
	return a
}

//SetRegion сеттер для региона
func (a *Builder) SetRegion(Region string) *Builder {
	a.address.region = "Region: " + Region
	return a
}

//SetCity сеттер страны
func (a *Builder) SetCity(City string) *Builder {
	a.address.city = "City: " + City
	return a
}

//SetPost сеттер для почтового индекса
func (a *Builder) SetPost(Post string) *Builder {
	a.address.post = "Post: " + Post
	return a
}

//SetHome сеттер для номера дома
func (a *Builder) SetHome(Home string) *Builder {
	a.address.home = "Home: " + Home
	return a
}

//SetUserName сеттер для имени пользователя
func (a *Builder) SetUserName(UserName string) *Builder {
	a.address.userName = "UserName: " + UserName
	return a
}

//SetPhone сеттер для номера телефона
func (a *Builder) SetPhone(Phone string) *Builder {
	a.address.phone = "Phone: " + Phone
	return a
}
