package unit

import (
	
	"testing"

	"github.com/AradinSeesura/Test/entity"
	
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
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

func TestPhone(t *testing.T){
	g:= NewGomegaWithT(t)

	t.Run(`Phone is not required`,func(t *testing.T) {
		user:=entity.User{
			StudentID: "B6628857",
			Email:"test@gmail.com",
			Phone:"",
			Profile: "",
			LinkedIn: "https://www.linkedin.com/company/ilink/",
			Password: "12345678",
			Age: 18,

		}
		ok,err:= govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Phone is required"))
	})
}

func TestAge(t *testing.T){
	g:= NewGomegaWithT(t)

	t.Run(`Age is required`,func(t *testing.T) {
		user:= entity.User{
			StudentID: "B6628857",
			Email:"test@gmail.com",
			Phone:"0952418469",
			Profile: "",
			LinkedIn: "https://www.linkedin.com/company/ilink/",
			Password: "12345678",
			Age: 10,
		}
		ok,err:= govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Age must be between 18 and 99"))
	})
}