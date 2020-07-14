package models

func GetAdmin(maps interface{}) (admin Admin) {
	db.Select("username").Where(maps).First(&admin)
	return
}
