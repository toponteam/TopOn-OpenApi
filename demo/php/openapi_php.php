<?php

$publisherKey = 'your publisher key';
$serverHost = 'https://openapi.toponad.com';

// POST example:
$contentArr   = [
    "limit"   => 1,
];
$reqUrl=$serverHost . "/v1/apps";
var_dump(doRequest("POST", $contentArr, $publisherKey, $reqUrl));

// GET example
$contentArr   = [
    "placement_id"   => "xxx",
];
$reqUrl=$serverHost . "/v1/waterfall/units";
var_dump(doRequest("GET", $contentArr, $publisherKey, $reqUrl));

/**
 * @param $httpMethod  GET/POST
 * @param array $contentArr
 * @param $publisherKey
 * @param $reqUrl
 * @return array
 */
function doRequest($httpMethod, array $contentArr, $publisherKey, $reqUrl)
{
    $time    = time() * 1000;
    $headerArr = [
        'X-Up-Key'       => $publisherKey,
        'X-Up-Timestamp' => $time,
    ];

    // create the content of common headers except the header 'X-Up-Signature'.
    $headerStr = '';
    foreach ($headerArr as $key => $value) {
        if (empty($headerStr)) {
            $headerStr = "$key:$value";
        } else {
            $headerStr = "$headerStr\n$key:$value";
        }
    }

    // create the MD5 value of the request body.
    // convert to the query param for the 'GET' method.
    if ($httpMethod == "GET") {
        $reqUrl=$reqUrl . "?" . http_build_query($contentArr);
        $contentStr  = "";
    } else {
        $contentStr = json_encode($contentArr);
    }
    $contentMD5  = strtoupper(md5($contentStr));
    // create the final signature.
    $contentType = 'application/json';
    $relativePath = parse_url($reqUrl)["path"];
    $signStr = $httpMethod . "\n"
        . $contentMD5 . "\n"
        . $contentType . "\n"
        . $headerStr . "\n"
        . $relativePath;
    $headerArr['X-Up-Signature'] = strtoupper(md5($signStr));;
    $headerArr['Content-Type'] = $contentType;

    return execCurl($httpMethod,$reqUrl,$contentStr,$headerArr);
}


function execCurl($httpMethod, $reqUrl, $contentStr, array $headerArr)
{
    $ch = curl_init();
    curl_setopt($ch, CURLOPT_URL, $reqUrl);
    if ($httpMethod == 'GET') {
        curl_setopt($ch, CURLOPT_POST, false);
        curl_setopt($ch, CURLOPT_HTTPGET, true);
    } elseif ($httpMethod == 'POST') {
        curl_setopt($ch, CURLOPT_POST, true);
        curl_setopt($ch, CURLOPT_POSTFIELDS, $contentStr);
    } else {
        curl_setopt($ch, CURLOPT_CUSTOMREQUEST, $httpMethod);
        curl_setopt($ch, CURLOPT_POSTFIELDS, $contentStr);
    }
    $finalHeaderArr = [];
    foreach ($headerArr as $key => $value) {
        $finalHeaderArr[] = $key . ":" . $value;
    }
    curl_setopt($ch, CURLOPT_HTTPHEADER, $finalHeaderArr);
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
    curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, FALSE);
    curl_setopt($ch, CURLOPT_SSL_VERIFYHOST, FALSE);

    $response = curl_exec($ch);
    $httpCode = curl_getinfo($ch, CURLINFO_HTTP_CODE);
    curl_close($ch);
    return array($httpCode, $response);
}