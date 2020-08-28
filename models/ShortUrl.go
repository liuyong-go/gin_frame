package models

//根据短链获取数据
func (sum ShortUrlMap) GetByShortUrl(surl string) string {
	var maps = make(map[string]string)
	maps["surl"] = surl
	db.Select("*").Where(maps).First(&sum)
	return sum.Lurl
}

//存储记录
func (sum ShortUrlMap) Record(data ShortUrlMap) bool {
	db.Create(&data)
	if sum.ID > 0 {
		return true
	} else {
		return false
	}
}
