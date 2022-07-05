package main

import (
	"fmt"
	"strings"
)

func wordfrequency(sentence string) (string,int){
	m := make(map[string]int)

	splitedwords := strings.Split(sentence, " ")

	for i:=0;i<len(splitedwords);i++{
		if _, found := m[splitedwords[i]]; found {
			m[splitedwords[i]]=m[splitedwords[i]]+1

		}else{
			m[splitedwords[i]]=1
		}
	}
    
	max:=0
    word:=""
	for key, value := range m {
     if value>max{
		max=value
		word=key
	 }
	}
return word,max
}
func main() {
   
    var sentence ="Geeks for Geeks is Best Geeks"
	key ,value:=wordfrequency(sentence)
	fmt.Println(key,"word has highest frequency of",value)
}