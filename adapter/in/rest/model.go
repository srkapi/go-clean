package controller

import(
	domains "firstgo_app/application/domain"
)

type CreatedUserDto struct {
	Name string `json:"name"`
	Lastname string `json:"lastname"`
}

type ResponseCreateUser struct {
	id int32 `json:"id"`
}



func (user CreatedUserDto) ToDomain() domains.UserDomain  {
	return domains.UserDomain {
		0,
	user.Name,
 user.Lastname,
	}

}
