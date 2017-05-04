package com.tuyy.LoanSystem;

public class LoanApplication {
	
	private CustomerManager customerManager;
	
	public LoanApplication() {
		this.customerManager = CustomerManager.getInstance();
	}
	
	public Customer searchCustomer(int id) {
		return customerManager.getCustomerWithLatestInfo(id);
	}
	
	public int addCustomer(String name) {
		return customerManager.appendCustomer(name);
	}
	
	public void removeCustomer(int id) {
		customerManager.deleteCustomer(id);
	}
	
	public boolean applyForLoan(int id, int money) {
		Customer customer = customerManager.getCustomerWithLatestInfo(id);
		return customer.applyForLoan(money);
	}
	
	public Loan repayLoan(int id, int money) {
		Customer customer = customerManager.getCustomerWithLatestInfo(id);
		return customer.repayLoan(money);
	}
}