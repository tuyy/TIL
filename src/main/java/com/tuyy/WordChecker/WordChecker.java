package com.tuyy.WordChecker;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class WordChecker {

	private List<String> ignoreWordLists;

	public WordChecker(List<String> ignoreWordLists) {
		this.ignoreWordLists = ignoreWordLists;
	}

	public Map<String, Integer> check(String text) {
		if (isNullOrEmptyText(text)) {
			return new HashMap<String, Integer>();
		} else {
			return getResultMap(text);
		}
	}
	
	private boolean isNullOrEmptyText(String emptyText) {
		return emptyText == null || emptyText.trim().equals("");
	}

	private Map<String, Integer> getResultMap(String text) {
		HashMap<String, Integer> resultMap = new HashMap<String, Integer>();
		String[] tokens = text.toLowerCase().split("\n");
		for (String line : tokens) {
			for (String word : line.split(" ")) {
				word = word.trim();
				if (isCountableWord(word)) {
					addWordCountWithResultMap(resultMap, word);
				}
			}
		}
		
		return resultMap;
	}

	private boolean isCountableWord(String word) {
		if (isCharacter(word) || isIgnoreWord(word)) {
			return false;
		} else {
			return true;
		}
	}

	private boolean isCharacter(String word) {
		return word.length() == 1;
	}

	private boolean isIgnoreWord(String word) {
		for (String ignoreWord : ignoreWordLists) {
			if (ignoreWord.equals(word)) {
				return true;
			}
		}
		return false;
	}

	private void addWordCountWithResultMap(HashMap<String, Integer> resultMap, String word) {
		if (isNullKey(resultMap, word)) {
			resultMap.put(word, 1);
		} else {
			resultMap.put(word, resultMap.get(word) + 1);
		}
	}

	private boolean isNullKey(HashMap<String, Integer> resultMap, String word) {
		return resultMap.get(word) == null;
	}
}