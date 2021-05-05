package utils

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func ParseBool(s string) bool {
	if s == "1" || s == "true" {
		return true
	}
	return false
}

func ParseUint32(s string) uint32 {
	value, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0
	}
	return uint32(value)
}
func ParseUint64(s string) uint64 {
	value, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0
	}
	return uint64(value)
}

func ParseInt32(s string) int32 {
	value, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0
	}
	return int32(value)
}
func ParseInt64(s string) int64 {
	value, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return int64(value)
}

//固定形式  x&y&z
func Split1(s string, retSlice *[]uint32) {
	slice := strings.Split(s, "&")
	*retSlice = make([]uint32, 0, len(slice))
	for _, value := range slice {
		*retSlice = append(*retSlice, ParseUint32(value))
	}
}

//固定形式   x&y&z;a&b&c;l_m_n
func Split2(s string, retSlice *[][]uint32) {
	slice1 := strings.Split(s, ";")
	*retSlice = make([][]uint32, 0, len(slice1))
	for _, value := range slice1 {
		var sl1 []uint32
		Split1(value, &sl1)
		*retSlice = append(*retSlice, sl1)
	}
}

//固定形式 x&y&z;a&b&c:x&y&z;a&b&c
func Split3(s string, retSlice *[][][]uint32) {
	slice1 := strings.Split(s, ":")
	*retSlice = make([][][]uint32, 0, len(slice1))
	for _, value := range slice1 {
		var sl2 [][]uint32
		Split2(value, &sl2)
		*retSlice = append(*retSlice, sl2)
	}
}

//固定形式  x&y&z
func SplitString1(s string, retSlice *[]string) {
	*retSlice = strings.Split(s, "&")
}

//固定形式  x&y&z;a&b&c;l_m_n
func SplitString2(s string, retSlice *[][]string) {
	slice1 := strings.Split(s, ";")
	*retSlice = make([][]string, 0, len(slice1))
	for _, value := range slice1 {
		*retSlice = append(*retSlice, strings.Split(value, "&"))
	}
}

//固定形式  x&y&z;a&b&c:x&y&z;a&b&c:
func SplitString3(s string, retSlice *[][][]string) {
	slice1 := strings.Split(s, ":")
	*retSlice = make([][][]string, 0, len(slice1))
	for _, value := range slice1 {
		var sl2 [][]string
		SplitString2(value, &sl2)
		*retSlice = append(*retSlice, sl2)
	}
}

//随机数返回[min,max)
func RandBetween(min, max int) int {
	if min >= max || max == 0 {
		return max
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	random := r.Intn(max-min) + min
	return random
}

func RandString(count int) string {
	var randomstr string
	for r := 0; r < count; r++ {
		i := RandBetween(65, 90)
		a := rune(i)
		randomstr += string(a)
	}
	return randomstr
}

//随机数返回[min,max)中的count个不重复数值
//一般用来从数组中随机一部分数据的下标
//2个随机数种子保证参数相同，返回值不一定相同，达到伪随机目的
func RandSliceBetween(min, max, count int) []int {
	if min > max {
		min, max = max, min
	}
	if min == max || max == 0 || count <= 0 {
		return []int{max}
	}
	randomRange := max - min
	retSlice := make([]int, 0, count)
	if count >= randomRange {
		for i := min; i < max; i++ {
			retSlice = append(retSlice, i)
		}
		return retSlice
	}
	r := rand.New(rand.NewSource(time.Now().Unix()))
	random := r.Intn(randomRange) + min
	baseRand := RandBetween(random*min, random*max)
	retSlice = append(retSlice, random)
	for i := 1; i < count; i++ {
		random = (i+baseRand*random)%randomRange + min
		isReapeated := false
		for j := 0; j < count; j++ {
			for _, v := range retSlice {
				if random == v {
					isReapeated = true
					break
				}
			}
			if isReapeated {
				random = (random-min+1)%randomRange + min
			} else {
				break
			}
		}
		retSlice = append(retSlice, random)
	}

	return retSlice
}

type valueWeightItem struct {
	weight uint32
	value  uint64
}

// 权值对，根据权重随机一个值出来
type GBValueWeightPair struct {
	allweight uint32
	valuelist []*valueWeightItem
}

func NewValueWeightPair() *GBValueWeightPair {
	return &GBValueWeightPair{}
}

func (this *GBValueWeightPair) Add(weight uint32, value uint64) {
	if weight == 0 {
		return
	}
	valueinfo := &valueWeightItem{weight, value}
	this.valuelist = append(this.valuelist, valueinfo)
	this.allweight += weight
}

func (this *GBValueWeightPair) Random() uint64 {
	if 1 == len(this.valuelist) {
		return this.valuelist[0].value
	}
	if this.allweight > 0 {
		randvalue := uint32(rand.Intn(int(this.allweight))) + 1 //[1,allweight]
		addweight := uint32(0)
		for i := 0; i < len(this.valuelist); i++ {
			addweight += this.valuelist[i].weight
			if randvalue <= addweight {
				return this.valuelist[i].value
			}
		}
	}
	return 0
}
func (this *GBValueWeightPair) GetValueList() []*valueWeightItem {
	return this.valuelist
}
func SafeSubInt32(a, b int32) int32 {
	if a > b {
		return a - b
	}
	return 0
}

func SafeSub(a, b uint32) uint32 {
	if a > b {
		return a - b
	}
	return 0
}
func SafeSub64(a, b uint64) uint64 {
	if a > b {
		return a - b
	}
	return 0
}

func SafeSubInt64(a, b int64) int64 {
	if a > b {
		return a - b
	}
	return 0
}

//三元运算符
func Ternary(val1 bool, ret1, ret2 interface{}) interface{} {
	if val1 {
		return ret1
	}
	return ret2
}
