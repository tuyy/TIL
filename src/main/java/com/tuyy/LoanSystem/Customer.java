package com.tuyy.LoanSystem;

public class Customer {
	
	private static int latestCustomerId = 0;
	
	private int id;
	private String name;
	private Loan loan;
	private CustomerGrade grade;

	public Customer(String name, CustomerGrade customerGrade) {
		this.id = latestCustomerId++;
		this.name = name;
		this.loan = new Loan();
		this.grade = customerGrade;
	}
	
	public int getId() {
		return id;
	}
	
	public String getName() {
		return name;
	}
	
	public boolean applyForLoan(int money) {
		if (grade.isPossibleLoan(money) == true) {
			loan.addLoan(money);
			return true;
		}else {
			System.out.println("대출불가");
			return false;
		}
	}
	
	public Loan repayLoan(int money) {
		loan.removePrincipalAndInterest(money);
		if (loan.isEmpty()) {
			grade = new GoldCustomerGrade();
		}
		return loan;
	}
	
    public void updateLatestLoanInfo() {
    	grade = loan.updateLoanInfoAndGetGrade(grade);
	}
    
    @Override
    public String toString() {
    	return "id: " + id + ", name: " + name + ", loan: " + loan;
    }
}