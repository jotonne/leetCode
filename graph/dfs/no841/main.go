package main

import "fmt"

// 自分で考えて最初にAcceptしたやつ
// func main() {
// 	rooms := [][]int{
// 		{1},
// 		{2},
// 		{3},
// 		{},
// 	}

// 	hasEntered := map[int]struct{}{
// 		0: {},
// 	}
// 	hasEntered = dfs(rooms, 0, hasEntered)
// 	fmt.Println((hasEntered))
// 	fmt.Println(len(rooms))
// }

// func dfs(rooms [][]int, key int, hasEntered map[int]struct{}) map[int]struct{} {
// 	if len(rooms[key]) == 0 {
// 		return hasEntered
// 	}

// 	for i := 0; i < len(rooms[key]); i++ {
// 		if _, ok := hasEntered[rooms[key][i]]; ok {
// 			continue
// 		}
// 		hasEntered[rooms[key][i]] = struct{}{}
// 		hasEntered = dfs(rooms, rooms[key][i], hasEntered)
// 	}
// 	return hasEntered
// }

func main() {
	rooms := [][]int{
		{1},
		{2},
		{3},
		{},
	}
	fmt.Println(canVisitAllRooms(rooms))
}

func canVisitAllRooms(rooms [][]int) bool {
	visit := map[int]bool{
		0: true,
	}

	var dfs func(rooms [][]int, key int)
	dfs = func(rooms [][]int, key int) {
		if len(rooms[key]) == 0 {
			return
		}
		for i := 0; i < len(rooms[key]); i++ {
			if ok := visit[rooms[key][i]]; ok {
				continue
			}
			visit[rooms[key][i]] = true
			dfs(rooms, rooms[key][i])
		}
	}

	dfs(rooms, 0)

	return len(visit) == len(rooms)
}
