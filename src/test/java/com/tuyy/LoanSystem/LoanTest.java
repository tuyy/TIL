package com.tuyy.LoanSystem;

import static org.junit.Assert.*;

import org.junit.Before;

import static org.hamcrest.CoreMatchers.*;
import org.junit.Test;

public class LoanTest {

	private Loan loan;
	
	@Before
	public void setUp() {
		loan = new Loan();
	}
	
	@Test
	public void testLoan_default_value() throws Exception {
		checkEmptyLoan(loan);
	}
	
	@Test
	public void test_addLoan_happy() throws Exception {
		loan.addLoan(500);
		
		assertThat(loan.getPrincipal(), is(500));
		assertThat(loan.getInterest(), is(0));
		assertThat(loan.getCreatedDate(), is(not(nullValue())));
		assertThat(loan.getDueDate(), is(not(nullValue())));
		assertThat(loan.getPaymentDate(), is(not(nullValue())));
	}

	@Test
	public void test_removeLoan_then_원금모두갚다() throws Exception {
		Loan loan = new Loan(500);
		loan.removePrincipalAndInterest(500);
		
		checkEmptyLoan(loan);
	}
	
	@Test
	public void test_removeLoan_then_일부이자만갚다() throws Exception {
		Loan loan = new Loan(500);
		loan.addInterest(400);
		loan.removePrincipalAndInterest(300);
		
		assertThat(loan.getPrincipal(), is(500));
		assertThat(loan.getInterest(), is(100));
	}
	
	@Test
	public void test_removeLoan_then_일부원금까지갚다() throws Exception {
		Loan loan = new Loan(500);
		loan.addInterest(400);
		loan.removePrincipalAndInterest(500);
		
		assertThat(loan.getPrincipal(), is(400));
		assertThat(loan.getInterest(), is(0));
	}
	
	@Test
	public void test_removeLoan_then_원금과이자보다큰돈을주다() throws Exception {
		Loan loan = new Loan(500);
		loan.addInterest(100);
		loan.removePrincipalAndInterest(1000);
		
		assertThat(loan.getPrincipal(), is(500));
		assertThat(loan.getInterest(), is(100));
	}
	
	private void checkEmptyLoan(Loan loan) {
		assertThat(loan.getPrincipal(), is(0));
		assertThat(loan.getInterest(), is(0));
		assertThat(loan.getCreatedDate(), is(nullValue()));
		assertThat(loan.getDueDate(), is(nullValue()));
		assertThat(loan.getPaymentDate(), is(nullValue()));
	}
}