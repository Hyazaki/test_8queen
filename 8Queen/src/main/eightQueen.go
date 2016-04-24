package main

import (
 "fmt";
 //matrix"github.com/skelterjohn/go.matrix"
)


const N int =8
//var  count int =99

var flg bool=true

func main(){

	fmt.Println("8queen -start-")
	queen :=make([]int,N)
	var count int=0
	var board [N][N]int

//	nBoard:=matrix.Zero(N,N)
//	fmt.Println(nBoard)
	
	for i :=0;i< N ;i++{//0行目のi列目にＱを置く

		count++
		//board[0][i]=1
		//fmt.Println(board)
		queen[0]=i		//0行目のi列目にＱを置く
		//置いたＱの筋を埋めに行く
		for  n :=0;n<N;n++{
			board[0][n]=1//横を埋める
			board[n][i]=1//縦を埋める
		//	fmt.Println(len(board[0]));
		//	fmt.Println(board[7][0])
		}
		for m:=0;m<N-i;m++{
			board[m][i+m]=1//右斜め下
		}
		for m:=0;m<i+1;m++{
			board[m][i-m]=1//左斜め下
		}
		
	//	fmt.Println(board)
		
		
		
		//fmt.Println(board)
		//1行目以降に置くＱを求める

		for j:=1;j<N;j++{//1行目からN行目まで
			for k:=0;k<N;k++{//0列目からN列目まで
				//fmt.Println(j,"行",k,"列 ： ",board[j][k])
				if board[j][k]==0/*checkBoard(j,k,board)*/{//置けたとき

					queen[j]=k
					
					if j==N-1{
						fmt.Println("Last ",j,"行",k,"列 :",board[j][k])
						break
					}
					
					for l:=0;l<N;l++{
						board[j][l]=1//横を埋める
						//fmt.Println("jl : ",board[j][l])
						board[l][k]=1//縦を埋める
						fmt.Println(l,"行",k,"列 :",board[l][k])
					}
					
					for m:=j;m<N-k;m++{
						board[m][k+m]=1//右斜め下
					}
					for m:=j;m<k+1;m++{
						board[m][k-m]=1//左斜め下
					}
					
					
					
					break
				}else{//置けなかった時
					//fmt.Println("確認",board[j][k])
				}
			}
		}
		flg = true
		fmt.Println(queen)
	//	fmt.Println(board)
		queen = initArray(queen)
		board = initBoard(board)
	}
	fmt.Println("8queen -finish-")
}



/*
func checkBoard(i int, j int,board [N][N]int) bool{
	var check bool

	if board[i][j]==0{
		check = true
	}else{
		check = false
	}
	return check
}
*/





func initArray(queen []int) []int{
	for i:=0;i<N;i++{
		queen[i]=0
	}
	return queen
}

func initBoard(board [N][N]int) [N][N]int{

	for i:=0;i<N;i++{
		for j:=0;j<N;j++{
			board[i][j]=0
		}
	}
	//fmt.Println(board)
	fmt.Println()
	return board
}