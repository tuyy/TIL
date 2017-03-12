package com.tuyy.WordChecker;

import java.io.BufferedReader;
import java.io.FileNotFoundException;
import java.io.FileReader;
import java.io.IOException;
import java.util.Arrays;
import java.util.List;
import java.util.Map;

public class WordCheckController {

	private String text;
	private List<String> ignoreWordList;

	public WordCheckController(String text, String ignoreText) {
		initController(text, ignoreText);
	}

	private void initController(String text, String ignoreText) {
		readFileAndSetText(text);
		readFileAndSetIgnoreWordList(ignoreText);
	}
	
	private void readFileAndSetText(String fileName) {
		try {
			this.text = readFile(fileName);
		} catch (IOException e) {}
	}

	
	private void readFileAndSetIgnoreWordList(String fileName) {
		try {
			String buffer = readFile(fileName);
			this.ignoreWordList = Arrays.asList(buffer.split("\n"));
		} catch (IOException e) {}
	}
	
	public String readFile(String fileName) throws FileNotFoundException, IOException {
		BufferedReader in = new BufferedReader(new FileReader(fileName));
		StringBuffer buffer = new StringBuffer();
		
		String str;
		while ((str = in.readLine()) != null) {
			buffer.append(str + "\n");
		}
		in.close();
		
		return buffer.toString().trim();
	}

	public void run() {
		WordChecker checker = new WordChecker(ignoreWordList);
		Map<String, Integer> resultMap = checker.check(text);
		WordCheckResult checkResult = new WordCheckResult(resultMap);
		checkResult.print();
	}
}