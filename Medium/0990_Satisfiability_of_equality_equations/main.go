// https://leetcode.com/problems/satisfiability-of-equality-equations/
package main

type VM struct {
	value_of [26]int
}

func MakeVM() (vm VM) {
	for i := range vm.value_of {
		vm.value_of[i] = i
	}
	return vm
}

func (this *VM) Assign(x, y int) {
	if vx, vy := this.ValueOf(x), this.ValueOf(y); vx != vy {
		this.value_of[vy] = vx
	}
}

func (this *VM) ValueOf(x int) int {
	if x != this.value_of[x] {
		this.value_of[x] = this.ValueOf(this.value_of[x])
	}
	return this.value_of[x]
}

func equationsPossible(equations []string) bool {
	vm := MakeVM()
	for _, eq := range equations {
		lhs, assign, rhs := int(eq[0]-'a'), eq[1] == '=', int(eq[3]-'a')
		if assign {
			vm.Assign(lhs, rhs)
		}
	}
	for _, eq := range equations {
		lhs, cmp, rhs := int(eq[0]-'a'), eq[1] == '=', int(eq[3]-'a')
		if cmp != (vm.ValueOf(lhs) == vm.ValueOf(rhs)) {
			return false
		}
	}
	return true
}

func main() {}
