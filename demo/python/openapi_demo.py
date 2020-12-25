# ~ /usr/bin/python

import hashlib
import json
import requests
import time

publisher_key = 'publisher key'
host = 'https://openapi.toponad.com/v1/apps'
path = '/v1/apps'

request_content = {}
timestamp = int(time.time() * 1e3)
headers = {
    'X-Up-Key': publisher_key,
    'X-Up-Timestamp': str(timestamp),  # milliseconds
}

http_method = 'POST'
content_md5 = hashlib.md5(json.dumps(request_content, sort_keys=False).encode()).hexdigest().upper()
content_type = 'application/json'
header_str = "X-Up-Key:" + publisher_key + "\n" + "X-Up-Timestamp:" + str(timestamp)

signature_str = '{}\n{}\n{}\n{}\n{}'.format(http_method, content_md5, content_type, header_str, path)
signature = hashlib.md5(signature_str.encode()).hexdigest().upper()
headers['X-Up-Signature'] = signature
headers['Content-Type'] = content_type

r = requests.post(host, headers=headers, data=json.dumps(request_content))
print(r.text)
