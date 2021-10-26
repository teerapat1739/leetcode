package leetcode

/**
 * Definition for Employee.
 * type Employee struct {
 *     Id int
 *     Importance int
 *     Subordinates []int
 * }
 */

/*
Runtime: 8 ms, faster than 98.96% of Go online submissions for Employee Importance.
Memory Usage: 6.8 MB, less than 87.50% of Go online submissions for Employee Importance.
*/

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {
	dict := make(map[int]*Employee, len(employees))
	for _, employee := range employees {
		dict[employee.Id] = employee
	}

	return sum690(id, dict)
}

func sum690(id int, employeesMap map[int]*Employee) int {
	importance := employeesMap[id].Importance
	for _, emp := range employeesMap[id].Subordinates {
		importance += sum690(emp, employeesMap)
	}
	return importance
}
