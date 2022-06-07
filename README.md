Welcome to "simple_api" for Cisco

1. Please write a small API, in the language of your choice, 
 - api written in golang

2. that allows users to log in using
a valid key. You should decide what is considered a valid vs. invalid key. Your service
should handle valid, invalid, and malformed keys.
 - using BasicAuth we are able validate the authentication header format, username and password

3. Instructions on how to run your service locally in a container
`cd cisco_test`
`docker build -t "test_api" .`
`docker run -p 8080:9000 -d test_api`

4. Instructions on how to run tests, first create your base64 encoded authentication string then using curl initiate a web request with Basic Auth header
`printf 'admin:password123!' | base64`
`curl -v http://localhost:8080/cisco -H "Authorization: Basic YWRtaW46cGFzc3dvcmQxMjMh"`
