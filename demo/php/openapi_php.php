<?php
/**
 * Created by PhpStorm.
 * User: zengzhihai
 * Date: 2019/11/14
 * Time: 14:44
 */

error_reporting(0);
$demoUrl = "https://openapi.toponad.com/v1/fullreport";
$body = "{}";
$publisherKey = "Your publisherKey";
$httpMethod = "POST";
$contentType = "application/json";
$publisherTimestamp = intval(microtime(true) * 1000);

$headerArrs = [
    'X-Up-Timestamp' => $publisherTimestamp,
    'X-Up-Key' => $publisherKey
];


$contentMd5 = strtoupper(md5($body));
var_dump($contentMd5);

$t = parse_url($demoUrl);
$resource = $t["path"];


$publisherSignature = signature($httpMethod, $contentMd5, $contentType, headerJoin($headerArrs), $resource);

$headerArrs['Content-Type'] = $contentType;
$headerArrs['X-Up-Signature'] = $publisherSignature;

$lastHeader = [];
foreach ($headerArrs as $k => $v) {
    $lastHeader[] = $k . ":" . $v;
}

var_dump(httpPostJson($demoUrl, $body, $lastHeader));


function httpPostJson($url, $jsonStr, $header = array())
{
    var_dump($header);
    $ch = curl_init();
    curl_setopt($ch, CURLOPT_POST, 1);
    curl_setopt($ch, CURLOPT_URL, $url);
    curl_setopt($ch, CURLOPT_POSTFIELDS, $jsonStr);
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
    curl_setopt($ch, CURLOPT_HTTPHEADER, $header);
    $response = curl_exec($ch);
    $httpCode = curl_getinfo($ch, CURLINFO_HTTP_CODE);
    curl_close($ch);
    return array($httpCode, $response);
}

function headerJoin($headers = [])
{
    $headerKeys = [
        "X-Up-Timestamp",
        "X-Up-Key"
    ];
    sort($headerKeys, SORT_STRING);
    $ret = [];
    foreach ($headerKeys as $v) {
        if ($headers[$v]) {
            $ret[] = $v . ":" . strval($headers[$v]);
        }
    }
    return implode($ret, "\n");
}

function signature($httpMethod, $contentMD5, $contentType, $headerString, $resource)
{
    $stringSection = [
        $httpMethod,
        $contentMD5,
        $contentType,
        $headerString,
        $resource
    ];
    $stringSection = implode($stringSection, "\n");
    return strtoupper(md5($stringSection));
}