package com.tuyy.WordChecker;

import static org.hamcrest.CoreMatchers.is;
import static org.hamcrest.object.HasToString.hasToString;
import static org.junit.Assert.assertThat;

import java.util.HashMap;
import java.util.Map;

import org.junit.Test;

import com.tuyy.WordChecker.WordCheckResult;

public class WordCheckerResultTest {

	@Test
	public void printCheckResult() throws Exception {
		Map<String, Integer> resultMap = new HashMap<String, Integer>();
		resultMap.put("love", 3);
		resultMap.put("like", 2);
		resultMap.put("money", 1);

		WordCheckResult checkResult = new WordCheckResult(resultMap);
		String expected = "1.[love] 3, 2.[like] 2, 3.[money] 1, ";
		assertThat(checkResult, is(hasToString(expected)));
		
		resultMap.put("lay", 4);
		expected = "1.[lay] 4, 2.[love] 3, 3.[like] 2, 4.[money] 1, ";
		assertThat(checkResult, is(hasToString(expected)));
	}
}
