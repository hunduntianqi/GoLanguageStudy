/*
	字符串操作:
		strings包提供了一些对于字符串的操作函数
		字符串常用操作函数:
			1. Contains:
				func Contains(s, substr, string) bool:判断字符串s中是否包含字符串substr, 返回bool值
			2. Join:
				func Join(a []string, sep string) string:将切片a中的字符串通过字符sep链接为一个新的字符串并返回
			3. Index:
				func Index(s, sep string) int: 在字符串s中查找字符sep对应索引, 返回索引值, 如果没有返回-1
			4. Repeat:
				func Repeat(s string, count int) string:将字符串s复制count次后返回新的字符串
			5. Replace:
				func Replace(s, old, new string, n int) string: 在s字符串中, 用new字符串替换old字符串, n表示
																替换次数, 小于0表示全部替换
			6. Split:
				func Split(s, sep string) []string:将字符串s通过sep字符进行切割, 生成一个字符串切片返回
			7. Trim:
				func Trim(s string, cutset string) string: 在字符串s中头部和尾部去除cutset指定的字符串
			8. Fields:
				func Fields(s string) []string: 去除字符串s中的空格符, 并按照空格分割返回字符串类型切片
*/

package main

func main() {

}
