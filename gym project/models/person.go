package models

type Person struct {
	name   string
	weight string
	height string
}
type Records struct {
	email    string
	password string
}

func (p Person) Name() string {
	return p.name
}

func (p *Person) SetName(name string) {
	p.name = name
}

func (p Person) Weight() string {
	return p.weight
}

func (p *Person) SetWeight(weight string) {
	p.weight = weight
}
func (p Person) Height() string {
	return p.height
}

func (p *Person) SetHeight(height string) {
	p.height = height
}

func (r Records) Email() string {
	return r.email
}

func (r *Records) SetEmail(email string) {
	r.email = email
}

func (r Records) Password() string {
	return r.password
}

func (r *Records) SetPassword(password string) {
	r.password = password
}
