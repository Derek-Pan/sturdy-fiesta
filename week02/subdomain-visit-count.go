package week02

import (
	"fmt"
	"strconv"
	"strings"
)


// 子域名访问次数

func subdomainVisits(cpdomains []string) []string {
	result := make(map[string]int)
	ans := make([]string, 0, 0)
	for _, val := range cpdomains {
		// 将字符串按" "进行切分,区分次数与域名
		res := strings.Split(val, " ")
		number, _ := strconv.Atoi(res[0])
		// 将域名“.”切分获取子域名
		subDomains := strings.Split(res[1], ".")
		//遍历子域名
		for i := len(subDomains); i > 0; i-- {
			dm := strings.Join(subDomains[len(subDomains)-i:len(subDomains)], ".")
			result[dm] += number
		}
	}
	for key, value := range result {
		ans = append(ans, fmt.Sprint(value, " ", key))
	}
	return ans
}
