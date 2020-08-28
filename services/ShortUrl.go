package services

import (
	libs_short "gin_frame/libs/shortUrl"
	"gin_frame/models"
)

type ShortUrl struct {
}

func (short ShortUrl) GetShortUrl(url string) string {
	var surl = libs_short.CreateShortURL(url)
	var modelShort models.ShortUrlMap
	var lurl = modelShort.GetByShortUrl(surl)
	if lurl == url {
		return surl
	}
	//不为空也不等于请求url说明短链重复，重新生成短链
	if lurl != "" {
		url += "!undefind"
		return short.GetShortUrl(url)
	}
	if lurl == "" { //存储记录
		var recordData models.ShortUrlMap
		recordData.Lurl = url
		recordData.Surl = surl
		modelShort.Record(recordData)
	}
	return surl
}
