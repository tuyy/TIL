package com.tuyy.LoanSystem;

import static org.junit.Assert.*;

import org.junit.Before;

import static org.hamcrest.CoreMatchers.*;

import org.junit.Test;

public class CustomerTest {
	
	private CustomerGrade bronzeCustomerGrade;
	private CustomerGrade silverCustomerGrade;
	private CustomerGrade goldCustomerGrade;

	@Before
	public void setup() {
		bronzeCustomerGrade = new BronzeCustomerGrade();
		silverCustomerGrade = new SilverCustomerGrade();
		goldCustomerGrade = new GoldCustomerGrade();
	}
	
	@Test
	public void test_applyForLoan_브론즈고객대출가능여부() throws Exception {
		assertThat(bronzeCustomerGrade.isPossibleLoan(-1), is(false));
		assertThat(bronzeCustomerGrade.isPossibleLoan(0), is(false));
		assertThat(bronzeCustomerGrade.isPossibleLoan(100), is(false));
		assertThat(bronzeCustomerGrade.isPossibleLoan(500), is(false));
		assertThat(bronzeCustomerGrade.isPossibleLoan(1000), is(false));
		assertThat(bronzeCustomerGrade.isPossibleLoan(5000), is(false));
	}
	
	@Test
	public void test_applyForLoan_실버고객대출가능여부() throws Exception {
		assertThat(silverCustomerGrade.isPossibleLoan(-1), is(false));
		assertThat(silverCustomerGrade.isPossibleLoan(0), is(false));
		assertThat(silverCustomerGrade.isPossibleLoan(100), is(true));
		assertThat(silverCustomerGrade.isPossibleLoan(500), is(true));
		assertThat(silverCustomerGrade.isPossibleLoan(1000), is(false));
		assertThat(silverCustomerGrade.isPossibleLoan(5000), is(false));
	}
	
	@Test
	public void test_applyForLoan_골드고객대출가능여부() throws Exception {
		assertThat(goldCustomerGrade.isPossibleLoan(-1), is(false));
		assertThat(goldCustomerGrade.isPossibleLoan(0), is(false));
		assertThat(goldCustomerGrade.isPossibleLoan(100), is(true));
		assertThat(goldCustomerGrade.isPossibleLoan(500), is(true));
		assertThat(goldCustomerGrade.isPossibleLoan(1000), is(true));
		assertThat(goldCustomerGrade.isPossibleLoan(5000), is(false));
	}
}
