package com.tuyy.WordChecker;

import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;

import static org.junit.Assert.*;
import static org.hamcrest.CoreMatchers.*;
import org.junit.Test;

import com.tuyy.WordChecker.WordCheckController;

public class WordCheckControllerTest {

	@Test
	public void createWordCheckController_with_two_filename() throws Exception {
		String textFileName = "/home/tuyy/text.txt";
		String ignoreWordListFileName = "/home/tuyy/ignoreWordList.txt";
		WordCheckController controller = new WordCheckController(textFileName, ignoreWordListFileName);
	}

	@Test
	public void testFileReader() throws Exception {
		try {
			String filename = "./text.txt";
			BufferedReader in = new BufferedReader(new FileReader(filename));
			StringBuffer buffer = new StringBuffer();
			String str;

			while ((str = in.readLine()) != null) {
				buffer.append(str + "\n");
			}
			in.close();
		} catch (IOException e) {
			fail(e.getMessage());
		}
	}

	@Test
	public void testRun() throws Exception {
		WordCheckController controller = new WordCheckController("text.txt", "ignoretext.txt");
		controller.run();
	}
}
