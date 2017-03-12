package com.tuyy.WordChecker;

import static org.hamcrest.CoreMatchers.is;
import static org.hamcrest.CoreMatchers.nullValue;
import static org.junit.Assert.assertThat;

import java.util.ArrayList;
import java.util.List;
import java.util.Map;

import org.junit.Before;
import org.junit.Test;

public class WordCheckerTest {

	private WordChecker wordChecker;

	@Before
	public void setUp() {
		List<String> ignoreWordList = new ArrayList<String>();
		ignoreWordList.add("am");
		ignoreWordList.add("are");
		ignoreWordList.add("were");
		ignoreWordList.add("you");
		wordChecker = new WordChecker(ignoreWordList);
	}
	
	@Test
	public void testCheck_null_or_emptyText() throws Exception {
		String emptyText = "";
		Map<String, Integer> result = wordChecker.check(emptyText);
		assertThat(result.size(), is(0));

		emptyText = null;
		result = wordChecker.check(emptyText);
		assertThat(result.size(), is(0));
	}
	
	@Test
	public void testCheck_happy() throws Exception {
		String text = "love happy happy bye";

		Map<String, Integer> result = wordChecker.check(text);
		assertThat(result.get("love"), is(1));
		assertThat(result.get("happy"), is(2));
		assertThat(result.get("bye"), is(1));
	}
	
	@Test
	public void testCheck_to_ignore_case() throws Exception {
		String text = "KING king KinG";
		
		Map<String, Integer> result = wordChecker.check(text);
		assertThat(result.get("king"), is(3));
	}
	
	@Test
	public void testCheck_if_exists_ignore_word() throws Exception {
		String text = "hello am are were";
		
		Map<String, Integer> result = wordChecker.check(text);
		assertThat(result.get("hello"), is(1));
		assertThat(result.get("am"), is(nullValue()));
		assertThat(result.get("are"), is(nullValue()));
		assertThat(result.get("were"), is(nullValue()));
	}
	
	@Test
	public void testCheck_if_exists_character_and_ignore_word() throws Exception {
		String text = "I love you";
		
		Map<String, Integer> result = wordChecker.check(text);
		assertThat(result.get("i"), is(nullValue()));
		assertThat(result.get("love"), is(1));
		assertThat(result.get("you"), is(nullValue()));
	}
}