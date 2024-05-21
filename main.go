package magic

import (
	"fmt"
	"math/rand"
	"time"
)

func Magic() {
	// 初始化4张扑克牌
	pokers := []int{2, 7, 6, 5}
	// 1. 洗牌
	rand.NewSource(time.Now().UnixNano())
	rand.Shuffle(len(pokers), func(i, j int) {
		pokers[i], pokers[j] = pokers[j], pokers[i]
	})
	fmt.Printf("1. 洗牌后的牌：%v\n", pokers)
	// 2. 撕牌
	pokers = append(pokers, pokers...)
	fmt.Printf("2. 撕牌后：%v\n", pokers)
	// 3. 你的名字有多少个字 就依次拿多少牌到牌底，比如：刘谦 2张
	pokers = append(pokers[2:], pokers[:2]...)
	fmt.Printf("3. 根据名字个数（刘谦）移动牌到牌底：%v\n", pokers)
	// 4. 将顶部3张牌随机放到牌中间（不能是最后！）
	pokers = append(pokers[3:7], pokers[0], pokers[1], pokers[2], pokers[7])
	fmt.Printf("4. 将顶部3张牌随机放到牌中间：%v\n", pokers)
	// 5. 将最上面的牌藏起来
	top := pokers[0]
	pokers = pokers[1:]
	fmt.Printf("5. 拿出最上面的 1 张牌： %d, %v\n", top, pokers)
	// 6. 再次插牌（南方：1张 北方：2张 不知道哪一方：3张）放在牌堆中间（不能是最后！），比如：北方人 2张
	// 6. 尼格买提....
	pokers = append(pokers[2:6], pokers[0], pokers[6], pokers[1])

	//pokers = append(pokers[2:6],pokers[0],pokers[1],pokers[6])
	fmt.Printf("6. 从上面拿出 2 张牌: %v\n", pokers)
	// 7. 男生将顶部一张牌扔掉，女生2张，比如：男生 1张
	pokers = pokers[1:]
	fmt.Printf("7. 男生或者女生取走对应数量的牌: %v\n", pokers)
	// 8. 见证奇迹的时刻，每念一个字，一张牌放在牌底
	for i := 0; i < 7; i++ {
		pokers = append(pokers[1:], pokers[0])
	}
	fmt.Printf("8. 见证奇迹的时刻：%v\n", pokers)
	// 9. 好运留下来（放到牌底） 烦恼丢出去（丢走）
	for i := 0; i < 5; i++ {
		pokers = append(pokers[1:], pokers[0])
		pokers = pokers[1:]
	}
	fmt.Printf("9. 好运留下来和烦恼丢出去：%v\n", pokers)
	if top == pokers[0] {
		fmt.Printf("10. %v == %v,新年好运", top, pokers[0])
	} else {
		fmt.Printf("10. %v == %v,其实呢对不上也会有好运，新年快乐", top, pokers[0])
	}

}
