package repository_user


func (db userRepository) Create(user User) error {
	err := db.db.Table("users").Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (db userRepository) Update(user User) error {
	err := db.db.Table("users").Model(&User{}).Where("id=?", user.ID).Updates(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (db userRepository) Delete(id string) error {
	err := db.db.Table("users").Where("id=?", id).Delete(&User{}).Error
	if err != nil {
		return err
	}
	return nil

}
func (db userRepository) FindAll() (users []User, err error) {
	err = db.db.Table("users").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
func (db userRepository) FindbyId(id string) (user User, err error) {
	err = db.db.Table("users").Where("id=?", id).First(&user).Error
	if err != nil {
		return
	}
	return user, nil
} 