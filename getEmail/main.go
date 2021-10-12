package main

import "getEmail/snov"

func main () {
	jsonValue := snov.GetEmail("boc.cn","0")
	snov.JsonToSlice(jsonValue)
}
