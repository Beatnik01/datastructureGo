package main

import "fmt"

//Diamond
func main(){
	for i:=0;i<4;i++{
		for j:=0;j<3-i;j++{
			fmt.Print("@")
		}
		for k:=0;k<1+i*2;k++{
			fmt.Print("*")
		}
		fmt.Println()
	}
	for x:=0;x<3;x++{
		for y:=0;y<x+1;y++{
			fmt.Print("@")
		}
		for z:=0;z<5-2*x;z++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}

/*
// Pyramid
func main(){
	for i:=0;i<5;i++{
		for j:=0;j<5-i;j++{
			fmt.Print("@")
		}
		for k:=0;k<1+i*2;k++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}
*/

/*
func main(){
	for i:=0;i<5;i++{
		for j:=0;j<i+1 && j<5-i;j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}
*/