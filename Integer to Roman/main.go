package main

func intToRoman(num int) string{
	count:=[]int{1000,900,500,400,100,90,50,40,10,9,5,4,1}
	Rims:=[]string{"M","CM","D","CD","C","XC","L","XL","X","IX","V","IV","I"}
	res:=""
	for i,val:=range count{
		j:=0
		for j=0;j<num/val;j++{
			res+=Rims[i]
		}
		num%=val
	}
return res
}



func main()  {
print(intToRoman(58))

}