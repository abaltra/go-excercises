package main

/*
	https://leetcode.com/problems/course-schedule/

	There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course bi first if you want to take course ai.

	For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
	Return true if you can finish all courses. Otherwise, return false.



	Example 1:

	Input: numCourses = 2, prerequisites = [[1,0]]
	Output: true
	Explanation: There are a total of 2 courses to take.
	To take course 1 you should have finished course 0. So it is possible.
*/

func canFinish(numCourses int, prerequisites [][]int) bool {
	/*
		Runtime: O(n) we traverse all elements one
		Memory: O(n) since we need to create a map of the nodes and an array of seen values
	*/
	courseToBlockers := make(map[int][]int, numCourses)

	for idx := 0; idx < numCourses; idx++ {
		courseToBlockers[idx] = []int{}
	}

	for _, pair := range prerequisites {
		course := pair[0]
		blockedBy := pair[1]
		courseToBlockers[course] = append(courseToBlockers[course], blockedBy)
	}

	seen := make([]int, numCourses)
	for course := 0; course < numCourses; course++ {
		if seen[course] == 0 {
			// this course has no blockers, we could start here
			if hasLoop(course, courseToBlockers, seen) {
				return false
			}
		}
	}

	return true
}

func hasLoop(course int, graph map[int][]int, seen []int) bool {
	if seen[course] == 2 {
		return true
	}

	seen[course] = 2

	for _, node := range graph[course] {
		if seen[node] != 1 {
			if hasLoop(node, graph, seen) {
				return true
			}
		}
	}
	seen[course] = 1
	return false
}
