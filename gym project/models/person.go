package models

type Person struct {
	Id     int     `db:"id"`
	Name   string  `db:"name"`
	Weight float64 `db:"weight"`
	Height float64 `db:"height"`
}
type Records struct {
	Email    string `db:"email"`
	Password string `db:"password"`
}

func (p Person) GetName() string {
	return p.Name
}

func (p *Person) SetName(name string) {
	p.Name = name
}

func (p Person) GetWeight() float64 {
	return p.Weight
}

func (p *Person) SetWeight(weight float64) {
	p.Weight = weight
}
func (p Person) GetHeight() float64 {
	return p.Height
}

func (p *Person) SetHeight(height float64) {
	p.Height = height
}

func (p Person) GetId() int {
	return p.Id
}

func (r Records) GetEmail() string {
	return r.Email
}

func (r *Records) SetEmail(email string) {
	r.Email = email
}

func (r Records) GetPassword() string {
	return r.Password
}

func (r *Records) SetPassword(password string) {
	r.Password = password
}
