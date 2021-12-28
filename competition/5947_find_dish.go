package main

import (
	"fmt"
)

// 1. 自己写，哈希表，Failed
//func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
//	recipeIdxList := []int{}
//	suppliedMap := make(map[string]int)
//	for _, supplie := range supplies {
//		if ok := suppliedMap[supplie]; ok > 0 {
//			suppliedMap[supplie]++
//		} else {
//			suppliedMap[supplie] = 1
//		}
//	}
//	for i, ingredient := range ingredients {
//		isSatisfied := true
//		for _, require := range ingredient {
//			if ok := suppliedMap[require]; ok > 0 {
//			} else {
//				isSatisfied = false
//			}
//		}
//		if isSatisfied {
//			recipeIdxList = append(recipeIdxList, i)
//			suppliedMap[recipes[i]] = 1
//		}
//	}
//	fmt.Println(suppliedMap)
//	recipeIdxList = []int{}
//	for i, ingredient := range ingredients {
//		isSatisfied := true
//		for _, require := range ingredient {
//			if ok := suppliedMap[require]; ok > 0 {
//			} else {
//				isSatisfied = false
//			}
//		}
//		if isSatisfied {
//			recipeIdxList = append(recipeIdxList, i)
//		}
//	}
//	res := []string{}
//	for _, i := range recipeIdxList {
//		res = append(res, recipes[i])
//	}
//	return res
//}

// 2. 自己写，哈希表，Failed！
//func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
//	recipeMap := make(map[string]int)
//	for i, recipe := range recipes {
//		recipeMap[recipe] = i+1
//	}
//	fmt.Println(recipeMap)
//	res := []string{}
//	for i, ingredient := range ingredients {
//		if deepSearch(ingredient, ingredients, recipeMap, supplies) {
//			res = append(res, recipes[i])
//		}
//	}
//	return res
//}
//
//func deepSearch(ingredient []string, ingredients [][]string, recipeMap map[string]int, supplies []string) bool {
//	for _, one := range ingredient {
//		fmt.Println(one)
//		fmt.Println(recipeMap[one])
//		if ok := recipeMap[one]; ok > 0 { // 判断是否为待做菜
//			if deepSearch(ingredients[recipeMap[one]-1], ingredients, recipeMap, supplies) {
//			} else {
//				return false
//			}
//		} else { // 不是待做菜，可以直接从原材料中找到
//			isSupplied := false
//			for _ , v := range supplies {
//				if one == v {
//					isSupplied = true
//				}
//			}
//			if !isSupplied {
//				return false
//			}
//		}
//	}
//	return true
//}


func findAllRecipes(recipes []string, ingredients [][]string, q []string) (ans []string) {
	g := map[string][]string{}
	deg := map[string]int{}
	for i, r := range recipes {
		for _, s := range ingredients[i] {
			g[s] = append(g[s], r) // 从这道菜的原料想这道菜连边
			deg[r]++
		}
	}
	for len(q) > 0 { // 拓扑排序（这里直接用初始原料当队列）
		s := q[0]
		q = q[1:]
		for _, r := range g[s] {
			if deg[r]--; deg[r] == 0 { // 表示原料都具备了
				q = append(q, r)
				ans = append(ans, r)
			}
		}

	}
	return
}


func main() {
	//recipes := []string{"bread"}
	//ingredients := [][]string{{"yeast","flour"}}
	//supplies := []string{"yeast"}

	recipes := []string{"bread","sandwich"}
	ingredients := [][]string{{"yeast","flour"},{"bread","meat"}}
	supplies := []string{"yeast","flour","meat"}

	//recipes := []string{"bread","sandwich","burger"}
	//ingredients := [][]string{{"yeast","flour"},{"bread","meat"},{"sandwich","meat","bread"}}
	//supplies := []string{"yeast","flour","meat"}
	fmt.Println(findAllRecipes(recipes, ingredients, supplies))
}

