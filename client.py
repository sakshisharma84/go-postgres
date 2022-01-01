import requests
import json

print("################################################################")
print("Executing READ request: GET: http://localhost:8080/api/vehicle/1")
url = "http://localhost:8080/api/vehicle/1"

#payload = "{\n    \"key1\": 1,\n    \"key2\": \"value2\"\n}"
payload = "{}"
headers = {
    'Content-Type': "application/json",
    'User-Agent': "PostmanRuntime/7.15.0",
    'Accept': "*/*",
    'Cache-Control': "no-cache",
    'Postman-Token': "e908a437-88ea-4b00-af53-7a9a49033830,ba90e008-0f7f-4576-beb8-b7739c8961f1",
    'Host': "httpbin.org",
    'accept-encoding': "gzip, deflate",
    'content-length': "42",
    'Connection': "keep-alive",
    'cache-control': "no-cache"
    }

response = requests.request("GET", url, data=payload, headers=headers)

print(response.text)


print("################################################################")
print("Executing GETALL request: http://localhost:8080/api/vehicle")

url = "http://localhost:8080/api/vehicle"

payload = "{}"
headers = {
    'Content-Type': "application/json",
    'User-Agent': "PostmanRuntime/7.15.0",
    'Accept': "*/*",
    'Cache-Control': "no-cache",
    'Postman-Token': "e908a437-88ea-4b00-af53-7a9a49033830,ba90e008-0f7f-4576-beb8-b7739c8961f1",
    'Host': "httpbin.org",
    'accept-encoding': "gzip, deflate",
    'content-length': "42",
    'Connection': "keep-alive",
    'cache-control': "no-cache"
    }

response = requests.request("GET", url, data=payload, headers=headers)

print(response.text)


print("################################################################")
print("Executing Create request: http://localhost:8080/api/newvehicle")

url = 'http://localhost:8080/api/newvehicle'
    
# Additional headers.
headers = {'Content-Type': 'application/json' } 

# Body
payload = {'vin': "addff", 'make': 'audi', 'model': 'q8', 'color': 'red', 'type': 'car', 'condition': 'new'}
    
# convert dict to json string by json.dumps() for body data. 
resp = requests.post(url, headers=headers, data=json.dumps(payload,indent=4))       
    
# Validate response headers and body contents, e.g. status code.
assert resp.status_code == 200
resp_body = resp.json()
    
# print response full body as text
print(resp.text)


print("################################################################")
print("Executing Update request: PUT http://localhost:8080/api/vehicle/1")

url = 'http://localhost:8080/api/vehicle/1'

# Additional headers.
headers = {'Content-Type': 'application/json' }

# Body
payload = {'vin': "NZXD", 'make': 'audi', 'model': 'q5', 'color': 'red', 'type': 'car', 'condition': 'used'}

# convert dict to json string by json.dumps() for body data.
resp = requests.put(url, headers=headers, data=json.dumps(payload,indent=4))

# Validate response headers and body contents, e.g. status code.
assert resp.status_code == 200
resp_body = resp.json()

# print response full body as text
print(resp.text)


print("################################################################")
print("Executing Update request for invalid ID: http://localhost:8080/api/vehicle/1000")

url = 'http://localhost:8080/api/vehicle/1000'

# Additional headers.
headers = {'Content-Type': 'application/json' }

# Body
payload = {'vin': "NZXD", 'make': 'audi', 'model': 'q5', 'color': 'red', 'type': 'car', 'condition': 'used'}

# convert dict to json string by json.dumps() for body data.
resp = requests.put(url, headers=headers, data=json.dumps(payload,indent=4))

# Validate response headers and body contents, e.g. status code.
assert resp.status_code == 500
resp_body = resp.json()

# print response full body as text
print(resp.text)

print("################################################################")
print("Executing Delete request: http://localhost:8080/api/vehicle/2")

url = 'http://localhost:8080/api/deletevehicle/2'

# Additional headers.
headers = {'Content-Type': 'application/json' }

# Body
payload = {}

# convert dict to json string by json.dumps() for body data.
resp = requests.delete(url, headers=headers, data=json.dumps(payload,indent=4))

# Validate response headers and body contents, e.g. status code.
assert resp.status_code == 200
resp_body = resp.json()

# print response full body as text
print(resp.text)


print("################################################################")
print("Executing Delete request for Invalid id: http://localhost:8080/api/vehicle/2000")

url = 'http://localhost:8080/api/deletevehicle/20000'

# Additional headers.
headers = {'Content-Type': 'application/json' }

# Body
payload = {}

# convert dict to json string by json.dumps() for body data.
resp = requests.delete(url, headers=headers, data=json.dumps(payload,indent=4))

# Validate response headers and body contents, e.g. status code.
assert resp.status_code == 500
resp_body = resp.json()

# print response full body as text
print(resp.text)


print("################################################################")
print("Executing Search request: http://localhost:8080/api/search/vehicle?color=red")

url = 'http://localhost:8080/api/search/vehicle?color=red'

# Additional headers.
headers = {'Content-Type': 'application/json' }

# Body
payload = {}

# convert dict to json string by json.dumps() for body data.
resp = requests.get(url, headers=headers, data=json.dumps(payload,indent=4))

# Validate response headers and body contents, e.g. status code.
assert resp.status_code == 200
resp_body = resp.json()

# print response full body as text
print(resp.text)

print("################################################################")
print("Executing Search request for invalid criteria: http://localhost:8080/api/search/vehicle?color=maroon")

url = 'http://localhost:8080/api/search/vehicle?color=maroon'

# Additional headers.
headers = {'Content-Type': 'application/json' }

# Body
payload = {}

# convert dict to json string by json.dumps() for body data.
resp = requests.get(url, headers=headers, data=json.dumps(payload,indent=4))

# Validate response headers and body contents, e.g. status code.
assert resp.status_code == 500
resp_body = resp.json()

# print response full body as text
print(resp.text)

print("     #####   XML Type   #####")
print("################################################################")
print("Executing GET request for XML: http://localhost:8080/api/vehicle/1")

xml = ""
headers = {'Content-Type': 'application/xml'} # set what your server accepts
print requests.get('http://localhost:8080/api/vehicle/1', data=xml, headers=headers).text


print("################################################################")
print("Executing GETALL request for XML: http://localhost:8080/api/vehicle")

xml = ""
headers = {'Content-Type': 'application/xml'} # set what your server accepts
print requests.get('http://localhost:8080/api/vehicle', data=xml, headers=headers).text

