package services

import (
	"users-backend/models"

	"github.com/jinzhu/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		db: db,
	}
}

// CreateUser - ฟังก์ชันในการเพิ่มผู้ใช้ใหม่
func (s *UserService) CreateUser(user *models.User) (*models.User, error) {
	tx := s.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	if err := tx.Create(user).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return user, nil
}

// GetUserByID - ฟังก์ชันในการดึงข้อมูลผู้ใช้จาก ID
func (s *UserService) GetUserByID(id string) (*models.User, error) {
	var user models.User
	if err := s.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetAllUsers - ฟังก์ชันในการดึงข้อมูลผู้ใช้ทั้งหมด
func (s *UserService) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := s.db.Order("id ASC").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// UpdateUser - ฟังก์ชันในการอัปเดตข้อมูลผู้ใช้
func (s *UserService) UpdateUser(id string, user *models.User) (*models.User, error) {
	var ex_user models.User
	if err := s.db.Where("id = ?", id).First(&ex_user).Error; err != nil {
		return nil, err
	}

	ex_user.FirstName = user.FirstName
	ex_user.LastName = user.LastName
	ex_user.Username = user.Username
	ex_user.Image = user.Image

	if err := s.db.Save(&ex_user).Error; err != nil {
		return nil, err
	}
	return &ex_user, nil
}

// DeleteUser - ฟังก์ชันในการลบผู้ใช้
func (s *UserService) DeleteUser(id string) error {
	var user models.User
	// ค้นหาผู้ใช้ในฐานข้อมูล
	if err := s.db.Where("id = ?", id).First(&user).Error; err != nil {
		return err
	}

	// ลบผู้ใช้จากฐานข้อมูล
	if err := s.db.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
