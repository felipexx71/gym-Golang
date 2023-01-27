package models

type Person struct {
	Name   string `db:"name"`
	Weight string `db:"weight"`
	Height string `db:"height"`
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

func (p Person) GetWeight() string {
	return p.Weight
}

func (p *Person) SetWeight(weight string) {
	p.Weight = weight
}
func (p Person) GetHeight() string {
	return p.Height
}

func (p *Person) SetHeight(height string) {
	p.Height = height
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
