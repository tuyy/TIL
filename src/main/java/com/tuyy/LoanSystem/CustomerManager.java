package com.tuyy.LoanSystem;

import java.util.HashMap;
import java.util.Map;

public class CustomerManager {
	private static CustomerManager instance = new CustomerManager();

	private Map<Integer, Customer> customerMap;
	
	private CustomerManager() {
		customerMap = new HashMap<>();
	}
	
	public static CustomerManager getInstance() {
		return instance;
	}
	
	public Customer getCustomerWithLatestInfo(int id) {
		Customer customer = customerMap.get(id);
		customer.updateLatestLoanInfo();
		return customer;
	}
	
	public void deleteCustomer(int id) {
		customerMap.remove(id);
	}
	
	public int appendCustomer(String name) {
		Customer customer = new Customer(name, new BronzeCustomerGrade());
		int customerId = customer.getId();
		customerMap.put(customerId, customer);
		return customerId; 
	}
}