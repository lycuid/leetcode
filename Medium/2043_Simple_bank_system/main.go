// https://leetcode.com/problems/simple-bank-system/
package main

type Bank struct{ vault []int64 }

func Constructor(balance []int64) Bank {
	vault := make([]int64, len(balance)+1)
	copy(vault[1:], balance)
	return Bank{vault}
}

func (this Bank) valid(num int) bool {
	return num < len(this.vault)
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if !this.valid(account1) || !this.valid(account2) || this.vault[account1] < money {
		return false
	}
	this.vault[account2] += money
	this.vault[account1] -= money
	return true
}

func (this *Bank) Deposit(account int, money int64) bool {
	if !this.valid(account) {
		return false
	}
	this.vault[account] += money
	return true
}

func (this *Bank) Withdraw(account int, money int64) bool {
	if !this.valid(account) || this.vault[account] < money {
		return false
	}
	this.vault[account] -= money
	return true
}

func main() {}
