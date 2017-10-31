# add path: C:\Python27\Scripts
# pip install flask

#!/usr/bin/python
import time
from datetime import datetime
import thread
import threading
import json
import urllib
import urllib2
import logging
import webbrowser
from flask import Flask, request, redirect, url_for

logging.basicConfig(filename='log_test_token.log' \
    , datefmt='%m/%d/%Y %H:%M:%S' \
    , format='%(asctime)s.%(msecs)d %(message)s' \
    , level=logging.DEBUG)

#API = 'api.nexon.net'
API = 'api-stage.nexon.net'  #api-stage
product_id = "30400"
product_key = "4a5dd491947ce9d3c231ee3a1526d1d706c9382a"

user = {
    'user_id' : 'hslee1115',
    'user_pw' : '',
    'token' : '',
    'refresh_token': '',
    'expire_in' : 0,
    'dt_updated' : ''
}

thread_list = []

API_ENDPOINTS = {
    'login':'https://' + API + '/login?id=%s&pw=%s',
    'ticket':'https://' + API + '/auth/ticket',
    'token':'https://' + API + '/auth/token',
    'refresh_token':'https://' + API + '/auth/token',
    'games':'https://' + API + '/library?pp=1&%s',
    'updateSession':'https://' + API + '/users/me/token',
    'profile':'https://' + API + '/users/me?access_token=%s',
    'activate':'https://' + API + '/activate',
    'access':'https://' + API + '/library/access?product_id=%s&access_token=%s'
}

def update_last_access_time():
    user['dt_updated'] = datetime.today().strftime('%Y-%m-%d %H:%M:%S')

def print_log(str):
    print(str+"\n")
    logging.debug(str)

def print_token_status(title, bToken, bRToken):
    if bToken:
        print_log("%s, token:%s" % (title, user['token']))
    if bRToken:
        print_log("%s, refresh-token:%s" % (title, user['refresh_token']))

def createTicket(token, prod_id):
    api_name = 'ticket'
    try:
        url = API_ENDPOINTS[api_name]
        data = {'access_token':token, 'prod_id':prod_id}
        #data = {'access_token':token, 'product_id':prod_id}
        data = json.dumps(data)
        logging.debug("[API:%s]:%s" % (api_name, "Start..."))
        req = urllib2.Request(url, data, {'Content-Type': 'application/json'})
        f = urllib2.urlopen(req)
        data = f.read()
        f.close()
        print_log("[API:%s] resp_data:%s" % (api_name, data))
        js = json.loads(data)
        print_log("[API:%s] ticket:%s" % (api_name, js[api_name]))
        return js[api_name]
        
    except urllib2.HTTPError, e:
        print_log("[API:%s][ERR] http code: %d %s" % (api_name, e.code, e.read()))
    except urllib2.URLError, e:
        print_log("[API:%s][ERR] network error: %s" % (api_name, e.reason.args[1]))

        
def createToken(ticket, product_id, secret_key):
    api_name = 'token'
    try:
        url = API_ENDPOINTS[api_name]
        data = {'ticket':ticket, 'product_id':product_id, 'secret_key':secret_key}
        data = json.dumps(data)
        logging.debug("[API:%s]:%s" % (api_name, "Start..."))
        req = urllib2.Request(url, data, {'Content-Type': 'application/json'})
        f = urllib2.urlopen(req)
        data = f.read()
        f.close()
        print_log("[API:%s] resp_data:%s" % (api_name, data))
        js = json.loads(data)
        return (js[api_name], js["refresh_token"])

    except urllib2.HTTPError, e:
        print_log("[API:%s][ERR] http code: %d %s" % (api_name, e.code, e.read()))
    except urllib2.URLError, e:
        print_log("[API:%s][ERR] network error: %s" % (api_name, e.reason.args[1]))
        

def refreshtoken(refresh_token, product_id, secret_key):
    api_name = 'refresh_token'
    try:
        url = API_ENDPOINTS[api_name]
        data = {'grant_type':'refresh_token', 'refresh_token':refresh_token, 'product_id':product_id, 'secret_key':secret_key}
        data = json.dumps(data)
        logging.debug("[API:%s]:%s" % (api_name, "Start..."))
        req = urllib2.Request(url, data, {'Content-Type': 'application/json'})
        f = urllib2.urlopen(req)
        data = f.read()
        f.close()
        print_log("[API:%s] resp_data:%s" % (api_name, data))
        js = json.loads(data)
        return (js["token"], js[api_name])

    except urllib2.HTTPError, e:
        print_log("[API:%s][ERR] http code: %d %s" % (api_name, e.code, e.read()))
    except urllib2.URLError, e:
        print_log("[API:%s][ERR] network error: %s" % (api_name, e.reason.args[1]))


def checkToken(token):
    api_name = 'token'
    try:
        url = '%s?%s' % (API_ENDPOINTS[api_name], urllib.urlencode({'token':token}))
        req = urllib2.Request(url, None, {'Content-Type': 'application/json'})
        f = urllib2.urlopen(req)
        data = f.read()
        f.close()
        print_log("[API:%s] resp_data:%s" % (api_name, data))
        js = json.loads(data)
        if js.has_key('success') and js['success'].has_key('data'):
            return  js['success']['data'].get('expires_in', 60)
        return 60
    except urllib2.HTTPError, e:
        print_log("[API:%s][ERR] http code: %d %s" % (api_name, e.code, e.read()))
    except urllib2.URLError, e:
        print_log("[API:%s][ERR] network error: %s" % (api_name, e.reason.args[1]))


def login(idx, uid, pw):
    logging.debug("")
    api_name = 'login'
    try:
        url = API_ENDPOINTS[api_name]%(uid,pw)
        data = {'id':uid,'pw':pw}
        data = json.dumps(data)
        logging.debug("[API:%s, %d]:%s" % (api_name, idx, "Start..."))
        req = urllib2.Request(url, data, {'Content-Type': 'application/json'})
        f = urllib2.urlopen(req)
        data = f.read()
        f.close()
        print_log("[API:%s, %d] resp_data:%s" % (api_name, idx, data))
        js = json.loads(data)
        #print_log("[API:%s] access_token:%s" % (api_name, js["access_token"]))
        return js["access_token"]

    except urllib2.HTTPError, e:
        print_log("[API:%s, %d][ERR] http code: %d %s" % (api_name, idx, e.code, e.read()))
    except urllib2.URLError, e:
        print_log("[API:%s][ERR] network error: %s" % (api_name, e.reason.args[1]))

def do_refreshtoken(token, refresh_token):
    while(1):
        print_log('%s'%user['expire_in'])
        time.sleep(user['expire_in'] + (60))  # expire_in + 1 min
        # to see it is valid:
        checkToken(user['token'])
        print_token_status('[%s].[before]'%do_refreshtoken.__name__, 0, 1)
        user['token'], user['refresh_token'] = refreshtoken(user['refresh_token'], product_id, product_key)
        user['expire_in'] = checkToken(user['token'])
        update_last_access_time()
        print_token_status('[%s].[end]'%do_refreshtoken.__name__, 0, 1)
        print_token_status('[%s].[end]'%do_refreshtoken.__name__, 1, 0)

bRunning = False
app = Flask(__name__)
@app.route('/ticket')
def index():
    global bRunning
    if bRunning:
        return redirect(url_for('status'))
    try:
        bRunning = True
        ticket = request.args.get('ticket')
        print_log("[flash:index] ticket:%s" % (ticket))
        user['token'], user['refresh_token'] = createToken(ticket, product_id, product_key)
        update_last_access_time()
        print_token_status(index.__name__, 1, 1)
        user['expire_in'] = checkToken(user['token'])
        thread.start_new_thread(do_refreshtoken, (user['token'], user['refresh_token']))
        return redirect(url_for('status'))
    except:
        print_log('error')

@app.route('/status')
def status():
    try:
        rs = "[%s] expire_in:%s, token:%s, refresh_token:%s" % (user['dt_updated'], user['expire_in'], user['token'], user['refresh_token'])
        return rs
    except:
        print_log('error')


def start_flask(port):
    app.run(debug=True, port=port)

def start_login(val):
    qry = { 'prod_id' : product_id, 'redirect_uri' : ('http://localhost:%d/ticket?' % port), 'int':val }
    qry = urllib.urlencode(qry)
    url = 'http://localhost:8080/api/test/async?%s' % qry
    webbrowser.open(url)

def test_invoice(user_id):
    url = 'http://localhost:8080/api/store/invoice'
    data = { 'product_id' : product_id,'user_id':user_id
			, 'user_ip':'127.0.0.1'
			, 'items': '12'
			, 'total_price': '1000'
			, 'data':'111' }
    data = json.dumps(data)
    req = urllib2.Request(url, data, {'Content-Type': 'application/json'})
    f = urllib2.urlopen(req)
    data = f.read()
    f.close()
    print_log("[API:%s] resp_data:%s" % ('test_invoice', data))
    js = json.loads(data)
    #return js["access_token"]

if __name__ == '__main__':
    try:
		# test invoice
		test_invoice('mantistest1')

        # port = 8080
        # thread.start_new_thread(start_login, (100,))
        # thread.start_new_thread(start_login, (18446744073709551615,))
        # time.sleep(1)
        # thread.start_new_thread(start_login, (100,))
        # time.sleep(1)
        # thread.start_new_thread(start_login, (18446744073709551615,))
        # time.sleep(1)
        # thread.start_new_thread(start_login, (100,))
        # # start_flask(port)
    except:
        print 'error starting flask'

    while 1:
       pass
