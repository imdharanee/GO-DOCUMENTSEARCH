package utils

import (
	"strings"
	"unicode"
	snowballeng "github.com/kljensen/snowball/english",
)
func lowercasefilter(token []string) []string{
	r:=make([]string,len(tokens))
	for i,token:=range tokens{
		r[i]=strings.ToLower(token)
	}
	return r
}

func stopwordfilter(tokens []string) []string {

	var stopwords=map[string]struct{
		"a":{}, "and":{},"be":{},"have":{},"i":{},
		"in":{}, "of":{}, "that":{},"the":{},"to":{}
	}
	r:=make([]string,0,len(tokens))
	for _,token:=range tokens {
		if _,ok:=stopwords[token]; !ok {
			r=append(r,token)

		}
	}
	return r

}

func stemmerfilter(tokens []string) []string{
	r:=make([]string,len(tokens))
	for i,token:=range tokens {

		r[i]=sbsnowballeng.Stem(token,false)
	}
	return r
}