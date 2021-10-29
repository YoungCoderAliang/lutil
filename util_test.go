package lutil

import (
	"fmt"
	"testing"
)

func TestRegrex(t *testing.T) {
	fmt.Println(IsMobile("18600000000"))
	fmt.Println(IsMobile("186000000001"))
	fmt.Println(ContainsMobile("1860000000000"))
	fmt.Println()

	fmt.Println(IsName("李三四"))
	fmt.Println(IsName("天下第一"))
	fmt.Println(ContainsName("辣妹子啦啦妹子啦"))
	fmt.Println()

	fmt.Println(IsBankCard("6222020000011111111"))
	fmt.Println(IsBankCard("62220200000111111111"))
	fmt.Println(ContainsBankCard("62220200000111111111"))
	fmt.Println()

	fmt.Println(IsIdCard("430881199808080088"))
	fmt.Println(IsIdCard("930881199808080088"))
	fmt.Println(IsIdCard("1430881199808080088"))
	fmt.Println(ContainsIdCard("1430881199808080088"))
	fmt.Println()

	fmt.Println(IsEmail("abc@bca.com"))
	fmt.Println(IsEmail("1abc@bca.com1"))
	fmt.Println(IsEmail("1abc@bcacom1"))
	fmt.Println(ContainsEmail("l@123.com"))
	fmt.Println(ContainsEmail("<dfexxcf>//123abc@bca.com急急急"))
	fmt.Println()

	fmt.Println(IsAddress("湖北省武汉市江汉路132号"))
	fmt.Println(IsAddress("上海市沪宁路abc单元"))
	fmt.Println(ContainsAddress("辣妹子啦啦妹子啦四川省成都市辣妹子啦辣妹子辣路啦啦啦"))
	fmt.Println()
}
