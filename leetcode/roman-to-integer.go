package main
//罗马数字的组数规则：
//  I=1，V=5，X=10，L=50，C=100，D=500，M=1000
//1、相同的数字连写，所表示的数等于这些数字相加得到的数，如，Ⅲ = 3；
//2、小的数字在大的数字的右边，所表示的数等于这些数字相加得到的数，如，Ⅷ = 8,Ⅻ = 12；
//3、小的数字，（限于Ⅰ、X 和 C）在大的数字的左边，所表示的数等于大数减小数得到的数，如，Ⅳ = 4，Ⅸ = 9；
//4、在一个数的上面画一条横线，表示这个数增值1 000倍，如
func romanToInt(s string) int {
	sum := 0
	/* 创建集合 */
	maps := make(map[int]int)
	/* map 插入 key-value 对，各个key对应的数字 */
	maps[73] = 1
	maps[86] = 5
	maps[88] = 10
	maps[76] = 50
	maps[67] = 100
	maps[68] = 500
	maps[77] = 1000

	var lens = len(s)
	for key, value := range s {
		captial, ok := maps[int(value)]

		fmt.Println(value)
		if ok {
			if key == lens-1 {
				sum += captial

				break;
			} else {
				next, ok := maps[int(s[key+1])]

				if ok {
					if captial < next {
						sum -= captial
					} else {
						sum += captial
					}
				}
			}

		} else {
			//不在该集合中~有问题
		}

	}
	return sum

}
