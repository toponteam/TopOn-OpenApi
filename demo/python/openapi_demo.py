#~ /usr/bin/python

import time
import hashlib
import json
from urlparse import urlparse
import requests


host = 'https://openapi.toponad.com/v1/fullreport'

payload = {
    'startdate': 20191103,
    'enddate': 20191105,
    'limit':120,
    'group_by':['date'],
    'metric':['all'],
    'start':0,
    'app_id':'Your app_id',
    'placement_id':'',
}
publisherTimestamp = int(time.time() * 1e3)
    
publisherKey = '6017bf891f5492422559c87437f8f62b95dc15cd'
    
headers = {
	'X-Up-Key': publisherKey,
    'X-Up-Timestamp': str(publisherTimestamp),  # milliseconds
}

http_method = 'POST'
content_md5 = hashlib.md5(json.dumps(payload, sort_keys=False).encode()).hexdigest().upper()
content_type = 'application/json'
header_str = "X-Up-Key:" + publisherKey + "\n" + "X-Up-Timestamp:" + str(publisherTimestamp);
resource = urlparse(host).path
signature_str = '{}\n{}\n{}\n{}\n{}'.format(http_method, content_md5, content_type, header_str, resource)
signature = hashlib.md5(signature_str.encode()).hexdigest().upper()
headers['X-Up-Signature'] = signature
headers['Content-Type'] = content_type
r = requests.post(host,headers=headers,data=json.dumps(payload))
print r.text
print r.status_code
