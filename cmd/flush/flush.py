#!/usr/bin/env python
#-*-coding:utf-8-*-

import urllib
import hashlib
import datetime
import requests

#========config_begin======
BUCKETNAME = ''
USERNAME = ''
PASSWORD = ''
purge = ''
#========config_end========

def httpdate_rfc1123(dt):
    """Return a string representation of a date according to RFC 1123
    (HTTP/1.1).

    The supplied date must be in UTC.
    """
    weekday = ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"][dt.weekday()]
    month = ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep",
             "Oct", "Nov", "Dec"][dt.month - 1]
    return "%s, %02d %s %04d %02d:%02d:%02d GMT" % \
        (weekday, dt.day, month, dt.year, dt.hour, dt.minute, dt.second)


date = httpdate_rfc1123(datetime.datetime.utcnow())
sign = hashlib.md5(purge+"&"+BUCKETNAME+"&" +
                       date + "&" + hashlib.md5(PASSWORD).hexdigest()).hexdigest()
Header = {"Expect": "",
          "Authorization": 'UpYun '+BUCKETNAME+':'+USERNAME+':'+sign,
          "Date": date,
          "Content-Type": "application/x-www-form-urlencoded",
          }
post = urllib.urlencode({'purge': purge})
r = requests.post("http://purge.upyun.com/purge/", post, headers=Header)
print r.text

