package com.tuyy.LoanSystem;

import java.util.Calendar;
import java.util.Date;

public class Loan {

	public static final int ONE_DAY = 86400;
	public static final int INTEREST_INCREASE_PERCENT = 3;
	
	private int principal;
	private int interest;
	private boolean isAleadyDownGrade;
	
	private Date createdDate;
	private Date dueDate;
	private Date paymentDate;
	private Date latestUpdatedDate;
	
	public Loan() {
		initLoanInfo();
	}
	
	public Loan(int borrowMoney) {
		addLoan(borrowMoney);
	}

	public int getPrincipal() {
		return principal;
	}
	
	public int getInterest() {
		return interest;
	}
	
	public Date getCreatedDate() {
		return createdDate;
	}
	
	public Date getDueDate() {
		return dueDate;
	}
	
	public Date getPaymentDate() {
		return paymentDate;
	}
	
	public Date getLatestUpdatedDate() {
		return latestUpdatedDate;
	}

	public void addLoan(int borrowMoney) {
		if (shouldPayBackLoan()) {
			// Exception already have a loan
			return;
		} else {
			principal = borrowMoney;
			createdDate = new Date();
			paymentDate = new Date();
			latestUpdatedDate = new Date();
			
			dueDate = createDueDateWithAddedDay(30);
		}
	}
	
	public void removePrincipalAndInterest(int moneyToPayOff) {
		int balance = (principal + interest - moneyToPayOff);
		if (balance < 0) {
			// 원금 + 이자 보다 많은 돈을 갚을 수 없다.
			return;
		}
		
		if (isAllPrincipalAndInterest(balance)) {
			initLoanInfo();
		} else if (isInterestOnly(moneyToPayOff)) {
			interest -= moneyToPayOff;
		} else {
			principal = balance;
			interest = 0;
		}
	}
	
	public void addInterest(int money) {
		interest += money;
	}
	
	public boolean isEmpty() {
		return principal == 0;
	}
	
	public void toggleisAleadyDownGrade() {
		isAleadyDownGrade = !isAleadyDownGrade;
	}
	
	private boolean isPassedDueDate() {
		if (dueDate == null) {
			return false;
		}
		
		Date now = new Date();
		long nowTime = now.getTime();
		if (dueDate.getTime() < nowTime) {
			long diff = nowTime - latestUpdatedDate.getTime();
			return (diff/ONE_DAY) > 0;
		} else {
			latestUpdatedDate = now;
			return false;
		}
	}

	private void increaseInterest() {
		Date now = new Date();
		long nowTime = now.getTime();
		long diff = nowTime - latestUpdatedDate.getTime();
		int addedDay = (int)(diff/ONE_DAY);
		interest += interest * (addedDay * INTEREST_INCREASE_PERCENT) / 100;
	}
	
	private boolean shouldPayBackLoan() {
		return (principal + interest) != 0;
	}

	private boolean isAllPrincipalAndInterest(int balance) {
		return balance == 0;
	}
	
	private boolean isInterestOnly(int moneyToPayOff) {
		return interest >= moneyToPayOff;
	}
	
	private void initLoanInfo() {
		principal = 0;
		interest = 0;
		isAleadyDownGrade = false;
		createdDate = null;
		paymentDate = null;
		dueDate = null;
		latestUpdatedDate = null;
	}

	private Date createDueDateWithAddedDay(int day) {
		Calendar cal = Calendar.getInstance(); 
		cal.setTime(new Date()); 
		cal.add(Calendar.DATE, day);
		
		return cal.getTime();
	}

	public CustomerGrade updateLoanInfoAndGetGrade(CustomerGrade grade) {
    	if (isPassedDueDate()) {
    		increaseInterest();
    		if (isAleadyDownGrade == false) {
    			grade = grade.getPreviousGrade();
    		}
    	}
		return grade;
	}
	
	@Override
	public String toString() {
		if (createdDate == null) {
			return "대출받은적없음";
		} else {
			return "대출원금: " + principal + ", 이자: " + interest + ", 대출일: " + createdDate + ", 납기일: " + dueDate;
		}
	}
}