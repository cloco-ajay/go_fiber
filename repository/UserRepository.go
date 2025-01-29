package repository

import (
	"sales-api/models"
	"sales-api/utils"
	"gorm.io/gorm"
	"sales-api/service/emailProvider"
	
)

type UserRepository interface {
	GetAllUsers() ([]models.User, error)
	GetUserById(id uint) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	DeleteUser(id uint) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) GetUserById(id uint) (models.User, error) {
	var user models.User
	err := r.db.Find(&user, id).Error
	return user, err
}

func (r *userRepository) CreateUser(user models.User) (models.User, error) {

	hashedPassword, hashError := utils.HashPassword(user.Password)
	if hashError != nil {
		return user, hashError
	}
	user.Password = hashedPassword
	err := r.db.Create(&user).Error

	// m := gomail.NewMessage()
	// m.SetHeader("From", "alex@example.com")
	// m.SetHeader("To", "bob@example.com", "cora@example.com")
	// m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	// m.SetHeader("Subject", "Hello!")
	// m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
	// // m.Attach("/home/Alex/lolcat.jpg")

	// d := gomail.NewDialer("sandbox.smtp.mailtrap.io", 2525, "ba711fbbe0c962", "561740d0aaff95")

	// // Send the email to Bob, Cora and Dan.
	// if err := d.DialAndSend(m); err != nil {
	// 	panic(err)
	// }

	if err == nil {
		// randomString, gErr := utils.GenerateRandomString(64)
		// if gErr != nil {
		// 	return user, gErr
		// }
		// toEncode := fmt.Sprint(string(*user.Email), "-", randomString, "-", user.ID)
		// base64Encode := base64.StdEncoding.EncodeToString([]byte(toEncode))
		// verificationUrl := fmt.Sprint(constant.GetBaseURL(), "/", "verify-email/", base64Encode)

		// data := map[string]interface{}{
		// 	"Name": user.Name,
		// 	"Url":  verificationUrl,
		// }
		err := emailProvider.SendEmail("User Registration Successful", []string{*user.Email}, "<p>shgdhf</p>")
		// err := service.sendEmail("User Registration Successful", []string{*user.Email}, "UserRegistration.html", data)
		if err != nil {
			r.DeleteUser(user.ID)
			return user, err
		}
	}
	return user, err
}

func (r *userRepository) DeleteUser(id uint) (models.User, error) {
	var user models.User

	err := r.db.Delete(&user, id).Error

	return user, err

}

func (r *userRepository) UpdateUser(user models.User) (models.User, error) {

	err := r.db.Model(&user).Updates(map[string]interface{}{"name": user.Name}).Error

	return user, err
}
