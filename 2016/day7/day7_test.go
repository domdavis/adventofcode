package day7

import "fmt"

func ExampleSolution_1() {
	fmt.Println(supportsTLS("abba[mnop]qrst"))
	// Output: true
}

func ExampleSolution_2() {
	fmt.Println(supportsTLS("abcd[bddb]xyyx"))
	// Output: false
}

func ExampleSolution_3() {
	fmt.Println(supportsTLS("aaaa[qwer]tyui"))
	// Output: false
}

func ExampleSolution_4() {
	fmt.Println(supportsTLS("ioxxoj[asdfgh]zxcvbn"))
	// Output: true
}

func ExampleSolution_5() {
	input := `abba[mnop]qrst
	abcd[bddb]xyyx
	aaaa[qwer]tyui
	ioxxoj[asdfgh]zxcvbn`
	fmt.Println(tlsSupportCount(input))
	// Output: 2
}

func ExampleSolution_6() {
	fmt.Println(supportsTLS("uvdgeatxkofgoyoi"))
	// Output: false
}

func ExampleSolution_7() {
	fmt.Println(supportsTLS("bxrcozuxifgevmog"))
	// Output: false
}

func ExampleSolution_8() {
	fmt.Println(supportsTLS("lknaffpzamlkufgt[uvdgeatxkofgoyoi]ajtqcsfdarjrddrzo[bxrcozuxifgevmog]rlyfschtnrklzufjzm"))
	// Output: true
}

func ExampleSolution_9() {
	fmt.Println(supportsTLS("oxpnarespssxlvbhe[mjactxdxwrjxjoa]aodrhgqmztemdxtbelo[vuslqwnngueagez]olwrulgbcmflenua"))
	// Output: false
}

func ExampleSolution_10() {
	fmt.Println(supportsTLS("tnntncwjwkepyjjj[rhiklabfhxebqoxjjd]nccutnmjduptofslfw[ztvcsayffkslzawquzf]bzicdywymajrjihcc[eaxrssfvhfvbswpqley]itikrgohakoqnmbxv"))
	// Output: true
}

func ExampleSolution_11() {
	fmt.Println(supportsTLS("xtxspkgfoepqquwnf[viivstflpbvqrhmyt]gormtajywcijwfbpmo[jnkgkcuhodivxggiw]fdpkuzlipozqwtiwiq"))
	// Output: false
}

func ExampleSolution_12() {
	fmt.Println(supportsTLS("nnevzrqtilosoamp[korgdgnaogoonln]ojjmrvbhjjylrnc[dzpncsqmuzsykyyxlru]ruvcsmwpqvsgkrd[ivjfkyskzxjlodhrcf]gaohcofquvhuyyu"))
	// Output: true
}

func ExampleSolution_13() {
	fmt.Println(supportsSSL("aba[bab]xyz"))
	// Output: true
}

func ExampleSolution_14() {
	fmt.Println(supportsSSL("xyx[xyx]xyx"))
	// Output: false
}

func ExampleSolution_15() {
	fmt.Println(supportsSSL("aaa[kek]eke"))
	// Output: true
}

func ExampleSolution_16() {
	fmt.Println(supportsSSL("zazbz[bzb]cdb"))
	// Output: true
}
