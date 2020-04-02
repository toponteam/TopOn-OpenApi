package com.test;

import com.opencsv.CSVParser;
import com.opencsv.CSVReader;

import java.io.DataInputStream;
import java.io.File;
import java.io.FileInputStream;
import java.io.InputStreamReader;

public class ParseCsv {
    public void main() {
        String filePath = "Z:/server/java/com.test/src/main/java/com/test/000000_0";
        try {
            DataInputStream in = new DataInputStream(new FileInputStream(new File(filePath)));
            CSVReader csvReader = new CSVReader(new InputStreamReader(in, "GBK"), CSVParser.DEFAULT_SEPARATOR,
                    CSVParser.DEFAULT_QUOTE_CHARACTER, CSVParser.DEFAULT_ESCAPE_CHARACTER, 1);
            String[] strs;
            while ((strs = csvReader.readNext()) != null) {
                for (String s : strs) {
                    System.out.print(s + "\n");
                }
            }
            csvReader.close();
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}

/**
 *
 *         这里是依赖库
 *         <dependency>
 *             <groupId>com.opencsv</groupId>
 *             <artifactId>opencsv</artifactId>
 *             <version>3.3</version>
 *         </dependency>
 *
 **/