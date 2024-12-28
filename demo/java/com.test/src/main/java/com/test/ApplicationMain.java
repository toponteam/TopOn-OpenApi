package com.test;

import org.apache.http.HttpEntity;
import org.apache.http.HttpResponse;
import org.apache.http.client.methods.HttpGet;
import org.apache.http.client.methods.HttpPost;
import org.apache.http.client.methods.HttpRequestBase;
import org.apache.http.entity.StringEntity;
import org.apache.http.impl.client.CloseableHttpClient;
import org.apache.http.impl.client.HttpClients;
import org.apache.http.util.EntityUtils;
import org.apache.commons.codec.digest.DigestUtils;
import java.net.URL;

public class ApplicationMain {

    public final static String CONTENT_TYPE = "application/json";
    public final static String PUBLISHER_KEY = "your publisher key";

    public static void main(String[] args) {

        // POST example:
        String reqBody = "{\"limit\":1}";
        String url = "https://openapi.toponad.com/v1/apps";
        String response = doRequest(HttpPost.METHOD_NAME, url, reqBody);
        System.out.println("post response: " + response);

        // GET example:
        url="https://openapi.toponad.com/v1/waterfall/units?placement_id=xxx";
        response = doRequest(HttpGet.METHOD_NAME, url, "");
        System.out.println("get response: " + response);
    }

    public static String doRequest(String httpMethod, String reqUrl, String reqBody) {
        String result = null;
        try {
            CloseableHttpClient httpClient = HttpClients.createDefault();
            HttpRequestBase httpRequest = null;
            if (httpMethod.equals(HttpPost.METHOD_NAME)) {
                HttpPost httpPost = new HttpPost(reqUrl);
                httpPost.setEntity(new StringEntity(reqBody,"utf-8")); // 兼容中文
//                httpPost.setEntity(new StringEntity(reqBody));
                httpRequest = httpPost;
            } else if (httpMethod.equals(HttpGet.METHOD_NAME)) {
                httpRequest = new HttpGet(reqUrl);
            } else {
                // TODO
            }
            // create the final signature
            String contentMD5 = DigestUtils.md5Hex(reqBody).toUpperCase();
            String nowMillis = System.currentTimeMillis() + "";
            String headerStr = "X-Up-Key:" + PUBLISHER_KEY + "\n" + "X-Up-Timestamp:" + nowMillis;
            String relativePath = new URL(reqUrl).getPath();
            String finalSign = genSignature(httpMethod, contentMD5, CONTENT_TYPE, headerStr, relativePath);
            // set the headers
            httpRequest.setHeader("Content-Type", CONTENT_TYPE);
            httpRequest.setHeader("X-Up-Timestamp", nowMillis);
            httpRequest.setHeader("X-Up-Key", PUBLISHER_KEY);
            httpRequest.setHeader("X-Up-Signature", finalSign);

            HttpResponse response = httpClient.execute(httpRequest);
            if (response != null) {
                HttpEntity resEntity = response.getEntity();
                result = resEntity != null ? EntityUtils.toString(resEntity, "utf-8") : "";
            }
        } catch (Exception e) {
            e.printStackTrace();
        }
        return result;
    }

    /**
     * create the final signature
     *
     * @param httpMethod   GET/POST
     * @param contentMD5
     * @param contentType
     * @param headerStr
     * @param relativePath the relative path of url, such as "/v1/apps"
     * @return
     */
    public static String genSignature(String httpMethod, String contentMD5,
                                      String contentType, String headerStr, String relativePath) {
        StringBuffer buf = new StringBuffer();
        buf.append(httpMethod).append('\n').
                append(contentMD5).append('\n').
                append(contentType).append('\n').
                append(headerStr).append('\n').
                append(relativePath);
        return DigestUtils.md5Hex(buf.toString()).toUpperCase();

    }
}
