Welcome to "simple_api" for Cisco

1. Please write a small API, in the language of your choice, 
 - api written in golang

2. that allows users to log in using
a valid key. You should decide what is considered a valid vs. invalid key. Your service
should handle valid, invalid, and malformed keys.
 - using BasicAuth we are able validate the authentication header format, username and password

3. Instructions on how to run your service locally in a container
- Change directory: `cd cisco_test`
- Build docker image: `docker build -t "test_api" . `
- Run Docker image: `docker run -p 8080:9000 -d test_api`

4. Instructions on how to run tests, first create your base64 encoded authentication string then using curl initiate a web request with Basic Auth header
- Testing strategy: test username, test password, test basic auth header, this can be achived by forming base64 encoded strings:
`printf 'admin:password123!' | base64`
- Then test by using curl:
`curl -v http://localhost:8080/cisco -H "Authorization: Basic YWRtaW46cGFzc3dvcmQxMjMh"`

Further notes:

Tradeoffs: 
- basic authentication, no encryption
- this code supports 1 user, and could be improved by using a data base and store user data, something like JWT
How would you improve the service if you had more time?:
- scale up the number of replicas
- expose behind a load balancer or k8s service or something like NGINX
How would you deploy this service to the cloud?:
- create a helm chart, deploy to a kubernetes cluster
How would you scale the service if you started seeing a lot of traffic?:
- increase number of replicas, increase resources if required
Ideas for automation and/or CI/CD:
- with helm override arguments automation can be created to deploy your application to different regions/data centers
Observability and monitoring strategy:
- latency: time it takes to respond
- traffic: how much traffic is coming into application , requests per second
- errors: number of failed requests
- saturation: how busy is the service, any resource contention? 
