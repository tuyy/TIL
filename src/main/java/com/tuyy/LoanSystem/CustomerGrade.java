package com.tuyy.LoanSystem;

public interface CustomerGrade {
	public boolean isPossibleLoan(int money);
	public CustomerGrade getPreviousGrade();
}
