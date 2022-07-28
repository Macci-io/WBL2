package address

import (
	"fmt"
	"strings"
)

//Address структура с данными об адресате
type Address struct {
	country, region, city, post, home, userName, phone string
}

//PrintInfo печатает информацию
func (a Address) PrintInfo() {
	fmt.Println(strings.Join([]string{a.country, a.region, a.city, a.post, a.home, a.userName, a.phone}, "\n"))
}
