package libs

//常量
const (
	BIGM = 0xc6a4a7935bd1e995
	BIGR = 47
	SEED = 0x1234ABCD
)

var tenToAny map[int64]string = map[int64]string{
	0:  "0",
	1:  "1",
	2:  "2",
	3:  "3",
	4:  "4",
	5:  "5",
	6:  "6",
	7:  "7",
	8:  "8",
	9:  "9",
	10: "a",
	11: "b",
	12: "c",
	13: "d",
	14: "e",
	15: "f",
	16: "g",
	17: "h",
	18: "i",
	19: "j",
	20: "k",
	21: "l",
	22: "m",
	23: "n",
	24: "o",
	25: "p",
	26: "q",
	27: "r",
	28: "s",
	29: "t",
	30: "u",
	31: "v",
	32: "w",
	33: "x",
	34: "y",
	35: "z",
	36: "A",
	37: "B",
	38: "C",
	39: "D",
	40: "E",
	41: "F",
	42: "G",
	43: "H",
	44: "I",
	45: "J",
	46: "K",
	47: "L",
	48: "M",
	49: "N",
	50: "O",
	51: "P",
	52: "Q",
	53: "R",
	54: "S",
	55: "T",
	56: "U",
	57: "V",
	58: "W",
	59: "X",
	60: "Y",
	61: "Z"}

//murmurHash64A 哈希算法
func murmurHash64A(data []byte) (h int64) {
	var k int64
	h = SEED ^ int64(uint64(len(data))*BIGM)

	var ubigm uint64 = BIGM
	var ibigm = int64(ubigm)
	for l := len(data); l >= 8; l -= 8 {
		k = int64(int64(data[0]) | int64(data[1])<<8 | int64(data[2])<<16 | int64(data[3])<<24 |
			int64(data[4])<<32 | int64(data[5])<<40 | int64(data[6])<<48 | int64(data[7])<<56)

		k := k * ibigm
		k ^= int64(uint64(k) >> BIGR)
		k = k * ibigm

		h = h ^ k
		h = h * ibigm
		data = data[8:]
	}

	switch len(data) {
	case 7:
		h ^= int64(data[6]) << 48
		fallthrough
	case 6:
		h ^= int64(data[5]) << 40
		fallthrough
	case 5:
		h ^= int64(data[4]) << 32
		fallthrough
	case 4:
		h ^= int64(data[3]) << 24
		fallthrough
	case 3:
		h ^= int64(data[2]) << 16
		fallthrough
	case 2:
		h ^= int64(data[1]) << 8
		fallthrough
	case 1:
		h ^= int64(data[0])
		h *= ibigm
	}

	h ^= int64(uint64(h) >> BIGR)
	h *= ibigm
	h ^= int64(uint64(h) >> BIGR)
	return
}

//TenToSixTwo 十进制转62进制缩短字符串
func TenToSixTwo(num int64) string {
	var str62 []byte
	for {
		var result byte
		var tmp []byte

		number := num % 62                   // 100%62 = 38
		result = []byte(tenToAny[number])[0] // C

		// 临时变量，为了追加到头部
		tmp = append(tmp, result)

		str62 = append(tmp, str62...)
		num = num / 62

		if num == 0 {
			break
		}
	}
	return string(str62)
}

//初始化短链配置
func init() {

}

//CreateShortURL 生成短链接
func CreateShortURL(lurl string) (surl string) {

	return
}

//GetURL 根据短链接获取长链接
func GetURL(surl string) (lurl string) {
	return
}
