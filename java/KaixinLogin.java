package com.yue.app.web.common.util;

import java.io.InputStream;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

import org.apache.commons.httpclient.Header;
import org.apache.commons.httpclient.HttpClient;
import org.apache.commons.httpclient.methods.GetMethod;
import org.apache.commons.httpclient.methods.PostMethod;
import org.apache.commons.io.IOUtils;
import org.apache.commons.lang.StringUtils;

public class KaixinLogin {

    private static Map<String, String> requestHeadersMap;
    private static final String        KX_username    = "yourname";
    private static final String        KX_password    = "passwd";
    private static final String        kaixinLoginUrl = "https://security.kaixin001.com/login/login_auth.php";

    public static void main(String[] argv) throws Exception {
        System.out.println(getKaixinCookie());
        GetMethod getMethod = getGetMethodWithHeader("http://www.kaixin001.com/home/");
        getMethod.setFollowRedirects(false);
        getMethod.addRequestHeader("Cookie", getKaixinCookie());
        HttpClient httpClient = new HttpClient();
        httpClient.executeMethod(getMethod);
        String homeHtmlContent = getMethod.getResponseBodyAsString();
        System.out.println(homeHtmlContent.contains("张小凡"));
    }

    private static String getKaixinCookie() {
        String setCookie = getLoginHeader(kaixinLoginUrl, "Set-Cookie");
        String cookie = getCookieFromSetCookie(setCookie);
        return cookie;
    }

    private static String getCookieFromSetCookie(String setCookie) {
        List<String> cookieStrList = new ArrayList<String>();
        for (String item : setCookie.split(";")) {
            if (item.indexOf('=') != -1) {
                String[] splits = item.split("=");
                String key = splits[0];
                String value = splits[1];
                if (!key.matches("^\\s*Domain\\s*$") && !key.matches("^\\s*Path\\s*$")) {
                    cookieStrList.add(String.format("%s=%s", key, value));
                }
            }
        }
        return StringUtils.join(cookieStrList, "; ");
    }

    private static String getLoginHeader(String kaixinLoginUrl, String header) {
        PostMethod postMethod = getPostMethodWithHeader(kaixinLoginUrl);
        postMethod.setParameter("email", KX_username);
        postMethod.setParameter("password", KX_password);
        postMethod.setParameter("rcode", "");
        postMethod.setFollowRedirects(false);
        InputStream inputStream = null;
        try {
            HttpClient httpClient = new HttpClient();
            httpClient.executeMethod(postMethod);
            postMethod.getResponseBodyAsString();
            Header[] headers = postMethod.getResponseHeaders(header);
            StringBuilder sb = new StringBuilder();
            for (int i = 0; i < headers.length; i++) {
                sb.append(headers[i].getValue());
                if (i != headers.length - 1) {
                    sb.append(";");
                }
            }
            return sb.toString();
        } catch (Exception e) {
            e.printStackTrace();
        } finally {
            IOUtils.closeQuietly(inputStream);
            postMethod.releaseConnection();
        }
        return null;
    }

    private static PostMethod getPostMethodWithHeader(String kaixinLoginUrl) {
        PostMethod postMethod = new PostMethod(kaixinLoginUrl);
        for (Map.Entry<String, String> entry : getRequestHeadersMap().entrySet()) {
            postMethod.addRequestHeader(entry.getKey(), entry.getValue());
        }
        return postMethod;
    }

    private static GetMethod getGetMethodWithHeader(String kaixinLoginUrl) {
        GetMethod getMethod = new GetMethod(kaixinLoginUrl);
        for (Map.Entry<String, String> entry : getRequestHeadersMap().entrySet()) {
            getMethod.addRequestHeader(entry.getKey(), entry.getValue());
        }
        return getMethod;
    }

    private static synchronized Map<String, String> getRequestHeadersMap() {
        if (requestHeadersMap != null) {
            return requestHeadersMap;
        }
        requestHeadersMap = new HashMap<String, String>();
        requestHeadersMap.put("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8");
        requestHeadersMap.put("Accept-Language", "en-US,en;q=0.5");
        requestHeadersMap.put("Host", "security.kaixin001.com");
        requestHeadersMap.put("User-Agent",
                              "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:18.0) Gecko/20100101 Firefox/18.0");
        requestHeadersMap.put("Content-Type", "application/x-www-form-urlencoded");
        requestHeadersMap.put("Referer", "https://security.kaixin001.com/");
        return requestHeadersMap;
    }
}

