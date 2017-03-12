package com.tuyy.WordChecker;

import java.util.ArrayList;
import java.util.Collections;
import java.util.Comparator;
import java.util.Iterator;
import java.util.List;
import java.util.Map;

public class WordCheckResult {

	private final int MAX_RANK = 25;
	
	private Map<String, Integer> resultMap;

	public WordCheckResult(Map<String, Integer> resultMap) {
		this.resultMap = resultMap;
	}

	public void print() {
		System.out.println(getWordCheckResultString());
	}

	@Override
	public String toString() {
		return getWordCheckResultString();
	}

	private String getWordCheckResultString() {
		int rank = 1;
		StringBuffer stringBuffer = new StringBuffer();
		
		Iterator<String> iter = getSortedIteratorWithResultMap();
		while(iter.hasNext() && rank <= MAX_RANK) {
			String key = iter.next();
			stringBuffer.append((rank++) + ".[" + key + "] " + resultMap.get(key) + ", ");
		}
		
		return stringBuffer.toString();
	}

	private Iterator<String> getSortedIteratorWithResultMap() {
		List<String> checkResultList = new ArrayList<String>();
		checkResultList.addAll(resultMap.keySet());
		Collections.sort(checkResultList, getComparatorForDescending());
		
		return checkResultList.iterator();
	}

	private Comparator<String> getComparatorForDescending() {
		return new Comparator<String>() {
			public int compare(String key1, String key2) {
				int v1 = resultMap.get(key1);
				int v2 = resultMap.get(key2);
				return ((Comparable<Integer>) v2).compareTo(v1);
			}
		};
	}
}
