package unit

import (
	 "testing"

 
	 "github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	 "github.com/AradinSeesura/Test/entity"
)

func TestValidate(t *testing.T) {
	 g:= NewGomegaWithT(t)

	 t.Run(`user is valid`,func(t *testing.T){
	 		user:= entity.User{
			StudentID: "B6628857",
			Email:"test@gmail.com",
			Phone:"0952418469",
			Profile: "",
			LinkedIn: "https://www.linkedin.com/company/ilink/",
			Password: "12345678",
			Age: 18,


 }

		ok,err:= govalidator.ValidateStruct(user)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
}