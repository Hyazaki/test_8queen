package main

import(
	"fmt";
)

const N int=8
var board[N][N]int
var queen [N]int
var count int =1

func main(){
	fmt.Println(N,"queen -start-")
	nQueen(0)
	fmt.Println(N,"queen -finish-")
}

func nQueen(m int){
	if N==m{
		fmt.Println("No.",count,":",queen)
		count++
	}else{
		for x:=0;x<N;x++{
			if checkBoard(m,x){
				continue
			}
			queen[m]=x
			fillBoard(m,x,1)
			nQueen(m+1)
			queen[m]=0
			fillBoard(m,x,-1)
		}
	}
}

func checkBoard(m int, x int)bool{
	var check bool=false
	if board[m][x]>0{
		check = true
	}
	return check
}

func fillBoard(m int, x int,num int){
	for l:=0;l<N;l++{
		board[m][l]+=num//横を埋める
		board[l][x]+=num//縦を埋める
	}
	var y int =x
	for i:=m;i<N;i++{
		if y<N{
			board[i][y]+=num//右斜め下
			//fmt.Println("y :",y)
		}else{
			break
		}
		y++
	}
	var z int =x
	for i:=m;i<N;i++{
		if z>=0{
			board[i][z]+=num//左斜め下
		}else{
			break
		}
		z--
	}
}