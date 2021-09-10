import hashlib
import json
import urllib.parse

import requests
import time

PUBLISHER_KEY = 'your publisher key'
CONTENT_TYPE = "application/json"


def do_request(http_method, req_url, req_body):
    now_millis = int(time.time() * 1e3)

    # Create the final signature.
    content_md5 = hashlib.md5(req_body.encode()).hexdigest().upper()
    header_str = "X-Up-Key:{}\nX-Up-Timestamp:{}".format(PUBLISHER_KEY, str(now_millis))
    relative_path = urllib.parse.urlsplit(req_url).path
    sign_str = "{}\n{}\n{}\n{}\n{}".format(http_method, content_md5, CONTENT_TYPE, header_str, relative_path)
    final_sign = hashlib.md5(sign_str.encode()).hexdigest().upper()

    # do the request
    headers = {
        "X-Up-Key": PUBLISHER_KEY,
        "X-Up-Timestamp": str(now_millis),
        "X-Up-Signature": final_sign,
        "Content-Type": CONTENT_TYPE,
    }
    if http_method == "POST":
        response = requests.post(req_url, headers=headers, data=req_body)
    elif http_method == "GET":
        response = requests.get(req_url, headers=headers)
    else:
        # TODO
        response = "todo"
    return response.text


# POST example:
req_url = 'https://openapi.toponad.com/v1/apps'
req_body = {"limit": 1}
req_body_str = json.dumps(req_body, sort_keys=False)
response = do_request("POST", req_url, req_body=req_body_str)
print("post response: ", response)

# GET example:
req_url = "https://openapi.toponad.com/v1/waterfall/units?placement_id=xxx&is_abtest=0"
response = do_request("GET", req_url, req_body="")
print("get response: ", response)
