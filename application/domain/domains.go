package domains




type UserDomain struct {
	Id int32
	Name string
	LastName string
}

func (e UserDomain) ToResponse() string  {
	return ""
}