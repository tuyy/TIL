package com.tuyy.LoanSystem;

public class main {

	public static void main(String[] args) {
		LoanApplication loanApp = new LoanApplication();
		int id = loanApp.addCustomer("테스터");
		
		System.out.println("1. 고객 id 검색");
		Customer customer = loanApp.searchCustomer(id);
		System.out.println(customer);

		System.out.println("2. 대출불가 금액 빌리기 요청");
		loanApp.applyForLoan(id, 400);
		System.out.println(customer);
		
		System.out.println("3. 100원 빌리기");
		loanApp.applyForLoan(id, 100);
		System.out.println(customer);
		
		System.out.println("4. 일부 원금 갚기");
		loanApp.repayLoan(id, 50);
		System.out.println(customer);
		
		System.out.println("5. 남은 금액 갚기");
		loanApp.repayLoan(id, 50);
		System.out.println(customer);
	}
}