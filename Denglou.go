package main

import "fmt"

type Hoom struct {//定义一个结构替，存放每个房间是否有楼梯和对应编号
	door int
	num int
}

func main() {
	var n,m int//n层*m个房间
	var hoom Hoom
	fmt.Scan(&n,&m)
	hooms := make([][]Hoom,n)
	for i := 0; i < n; i++ {
		hooms[i] = make([]Hoom,m)
		for j := 0; j < m; j++ {
			fmt.Scan(&hoom.door,&hoom.num)
			hooms[i][j] = hoom
		}
	}

	var enter int//进入的房间

	fmt.Scan(&enter)

	startNum := hooms[0][enter].num//开始进入房间的编号
	password := startNum//密码

	temp := 0//用作标记找到有楼梯房间数
	for i := 0; i < n; i++ {//从第一层开始进入
		for temp < startNum {
			if hooms[i][enter].door == 1{
				temp++
			}
			enter++//寻找下一个房间
			if enter >= m {//如果所进入的房间编号大于m,从0再次开始
				enter = 0
			}
		}
		if i != n-1 {
			if enter == 0 {//这里判断下enter为0，此时应为上一层最后一个房间进来的
				enter = m-1
			}else {
				enter = enter-1
			}
			startNum = hooms[i+1][enter].num//进入下一层，记录看到的编号
			password += startNum
		}
		temp = 0//更新找到房间数
	}

	fmt.Println(password)

}
