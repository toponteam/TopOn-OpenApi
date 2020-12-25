package com.test;

import org.apache.commons.codec.digest.DigestUtils;
import org.apache.http.HttpEntity;
import org.apache.http.HttpResponse;
import org.apache.http.client.HttpClient;
import org.apache.http.client.methods.HttpPost;
import org.apache.http.entity.StringEntity;
import org.apache.http.message.BasicHeader;
import org.apache.http.util.EntityUtils;
import com.test.SSLClient;

import java.security.MessageDigest;

public class ApplicationMain {

    public static String signature(String httpMethod, String contentMD5,
                                   String contentType, String headerString, String resource) {

        String SignString = httpMethod + "\n" + contentMD5 + "\n" + contentType + "\n" + headerString + "\n" + resource;

        System.out.println(SignString);

        return ApplicationMain.md5Default(SignString).toUpperCase();

    }


    public static String doPost(String url, String cintent, String publisherKey, String publisherSignature,
                                String publisherTimestamp) {
        HttpClient httpClient = null;
        HttpPost httpPost = null;
        String result = null;
        try {
            httpClient = new SSLClient();
            httpPost = new HttpPost(url);
            httpPost.setHeader("Content-Type", "application/json");
            httpPost.setHeader("X-Up-Timestamp", publisherTimestamp);
            httpPost.setHeader("X-Up-Key", publisherKey);
            httpPost.setHeader("X-Up-Signature", publisherSignature);

            httpPost.setEntity(new StringEntity(cintent));
            HttpResponse response = httpClient.execute(httpPost);
            if (response != null) {
                HttpEntity resEntity = response.getEntity();
                if (resEntity != null) {
                    result = EntityUtils.toString(resEntity, "utf-8");
                }
            }
        } catch (Exception ex) {
            ex.printStackTrace();
        }
        return result;
    }

    public static void main(String[] args) {


        String cintent = "{}";

        String contentMD5 = ApplicationMain.md5Default(cintent).toUpperCase();
        System.out.println(cintent);
        String url = "https://openapi.toponad.com/v2/fullreport";
        String publisherTimestamp = "" + System.currentTimeMillis();
        String publisherKey = "Your publisherKey";
        String headerString = "X-Up-Key:" + publisherKey + "\n" + "X-Up-Timestamp:" + publisherTimestamp;
        String publisherSignature = ApplicationMain.signature("POST", contentMD5, "application/json", headerString, "/v1/fullreport");

        String s = ApplicationMain.doPost(url, cintent, publisherKey, publisherSignature, publisherTimestamp);
        System.out.println(s);

    }

    public static String getMD5(String str) {
        try {
            return DigestUtils.md5Hex(str);
        } catch (Exception e) {
            return null;
        }
    }

    public final static String md5Default(String str) {
        try {
            MessageDigest md5 = MessageDigest.getInstance("MD5");
            md5.update(str.getBytes());
            byte b[] = md5.digest();

            StringBuffer sb = new StringBuffer("");
            for (int n = 0; n < b.length; n++) {
                int i = b[n];
                if (i < 0) i += 256;
                if (i < 16) sb.append("0");
                sb.append(Integer.toHexString(i));
            }
            return sb.toString();
        } catch (Exception e) {
            // ignore
        }
        return null;
    }


}
