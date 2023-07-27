# Getting Started
This document guides you through the installation of required dependencies, project setup, and using the basic functionalities.

## Prerequisites
To work with Humble API Gateway, you need to install Go and Etcd. Follow the instructions below to install these prerequisites.

## Go Installation
### Mac Users:
1. Download the Go binary from the official Go [download page](https://go.dev/doc/install).
2. Open the package file you downloaded and follow the prompts to install Go.
   The package installs the Go distribution to /usr/local/go. The package should put the /usr/local/go/bin directory in your PATH environment variable. You may need to restart any open Terminal sessions for the change to take effect.
   Verify that you've installed Go by opening a command prompt and typing the following command:
   shell
   $ go version

Confirm that the command prints the installed version of Go.

### Windows users:
1. Download the Go binary from the official Go [download page](https://go.dev/doc/install).
2. Open the MSI file you downloaded and follow the prompts to install Go.
   By default, the installer will install Go to Program Files or Program Files (x86). You can change the location as needed. After installing, you will need to close and reopen any open command prompts so that changes to the environment made by the installer are reflected at the command prompt.
3. Verify that you've installed Go.
   In Windows, click the Start menu.
   In the menu's search box, type cmd, then press the Enter key.
   In the Command Prompt window that appears, type the following command:
   $ go version


## Etcd Installation
### Windows Users:
The easiest way to install etcd on Windows is through Docker. Docker allows you to run applications in containers, which are like lightweight virtual machines. Here's a step-by-step guide to get etcd running on your Windows machine:
Install Docker Desktop for Windows: You can download it from the official Docker website. Once you've downloaded the installer, run it and follow the instructions to install Docker Desktop. After installation, you should be able to run Docker commands in the Windows Command Prompt or PowerShell.
Pull the etcd Docker image: Open Command Prompt or PowerShell and pull the official etcd image from Docker Hub by running:
bash
docker pull gcr.io/etcd-development/etcd:v3.5.0
You can replace v3.5.0 with the version number of etcd you wish to use.
Run etcd in a Docker container: You can now start an etcd container with the following command:
perl
docker run -p 2379:2379 -p 2380:2380 gcr.io/etcd-development/etcd:v3.5.0 /usr/local/bin/etcd --advertise-client-urls http://0.0.0.0:2379 --listen-client-urls http://0.0.0.0:2379
This command starts a Docker container running etcd, and maps the container's ports 2379 and 2380 to the same ports on your host machine, so you can connect to etcd at localhost:2379.


### Mac Users:
Run the following command in the terminal:
brew install etcd

## Setting up Humble

### Mac Users:
1. Clone the project from GitHub:
   shell
   git clone https://github.com/redtailedfox/humble.git
2. Navigate to the project directory:
   shell
   cd humble
3. Make the bash script executable:
   chmod +x scripts/start-all.sh
4. Please ensure that the localhost ports 8080, 8888, 8889, 8890 and 8891 are available.
5. Run the bash script:
   ./start-all.sh
6. The gateway is now running and you may now use our following services.

### Windows Users:
1. Clone the project from GitHub:
   shell
   git clone https://github.com/redtailedfox/humble.git
2. Navigate to the project directory:
   shell
   cd humble
3. Please ensure that the localhost ports 8080, 8888, 8889, 8890 and 8891 are available.
4. Run the batch scripts:
   If this is your first time running it, go to the humble directory, right click on update-netpoll.bat, and click open.
   Go to the humble directory, right click on start-all.bat, and click open.
5. The gateway is now running and you may now use our following services.
6. To stop the gateway close all the cmd windows that popped up.



## Using the services
You can test the different methods of Humble using curl commands as follows:

### Mac Users:
1. The call method: This method echoes the message sent into the gateway and adds a Hello before it.
   shell
   curl --location --request POST 'http://127.0.0.1:8080/post' \
   --header 'Content-Type: application/json' \
   --data-raw '{"message": "Arbitrary Name or Value"}'

2. The concatenate method: This method concatenates the two messages sent into the gateway and returns it.

shell
curl --location --request POST 'http://127.0.0.1:8080/concat' \
--header 'Content-Type: application/json' \
--data-raw '{"message1": "Arbitrary Name or Value", "message2": "Arbitrary Name or Value"}'

3. The encrypt method: This method encrypts the message using our encryption method.
   shell
   curl --location --request POST 'http://127.0.0.1:8080/encrypt' \
   --header 'Content-Type: application/json' \
   --data-raw '{"message": "Arbitrary Name or Value"}'

4. The decrypt method: This method will decrypt any message encrypted using our encryption method.

shell
curl --location --request POST 'http://127.0.0.1:8080/decrypt' \
--header 'Content-Type: application/json' \
--data-raw '{"message": "Arbitrary Name or Value"}'
Please replace "Arbitrary Name or Value" with your own content before running these commands.


### Windows Users:
1. The call method: This method echoes the message sent into the gateway and adds a Hello before it.
   shell
   Invoke-RestMethod -Uri 'http://127.0.0.1:8080/concat' -Method POST -Headers @{"Content-Type" = "application/json"} -Body '{"message": "Arbitrary Name or Value"}'

2. The concatenate method: This method concatenates the two messages sent into the gateway and returns it.

shell
Invoke-RestMethod -Uri 'http://127.0.0.1:8080/concat' -Method POST -Headers @{"Content-Type" = "application/json"} -Body '{"message1": "Arbitrary Name or Value", "message2": "Arbitrary Name or Value"}'

3. The encrypt method: This method encrypts the message using our encryption method.
   shell
   Invoke-RestMethod -Uri 'http://127.0.0.1:8080/encrypt' -Method POST -Headers @{"Content-Type" = "application/json"} -Body '{"message": "Arbitrary Name or Value"}'

4. The decrypt method: This method will decrypt any message encrypted using our encryption method.

shell
Invoke-RestMethod -Uri 'http://127.0.0.1:8080/decrypt' -Method POST -Headers @{"Content-Type" = "application/json"} -Body '{"message": "Arbitrary Name or Value"}'

Please replace "Arbitrary Name or Value" with your own content before running these commands.


Software Design Document

Table of Contents

Introduction
Technology Stack
System Architecture
Detailed System Design
Performance Considerations and Comparisons
Testing Strategy
Conclusion


1. Introduction

This Software Design Document presents a detailed exploration of the architectural blueprint and design methodology used in creating an API Gateway. The Gateway is built leveraging the high-performance capabilities of CloudWeGo Hertz and Kitex. It processes HTTP requests with JSON-encoded bodies, converts these requests to Thrift binary format, and efficiently routes them to backend RPC servers. The inclusion of a service registry and discovery mechanism ensures system robustness and adaptability.

2. Technology Stack

Programming Language: Golang
API Gateway: Hertz (with the Generic-Call feature)
HTTP Server Framework: CloudWeGo Hertz: An HTTP library for efficient management and processing of HTTP requests.
RPC Server Framework: CloudWeGo Kitex: A high-performance RPC framework that offers swift and efficient message serialization and deserialization.
Data Interchange Formats: JSON for incoming requests, Thrift for internal communication
Load Balancer: Kitex-provided Load Balancer: Weighted Round Robin (from kitex-contrib)
Service Register and Discovery: Kitex-provided Service Register: Etcd
Version Control System: Git



3. System Architecture

Clients: The clients send HTTP requests encoded in JSON to the API Gateway.
API Gateway: This is the entry point to the system, developed using Golang and the Kitex framework. It accepts HTTP requests in JSON format, then utilizes the Generic-Call feature of Kitex to translate these requests into Thrift binary format.
Service Register and Discovery: Services register themselves using the Service Register provided by Kitex when they start up. The Service Discovery mechanism is used by the API Gateway to find suitable backend services to handle the incoming requests.
Load Balancer: The Load Balancer, provided by Kitex, distributes the incoming Thrift requests to the available backend RPC servers.
RPC Servers: These are developed using Kitex and are responsible for handling the Thrift requests. After processing, they return the responses back to the API Gateway which, in turn, sends the results back to the client.
HTTP Server: An HTTP server is built using Hertz to accept and manage the incoming HTTP requests.


			Given diagrams of the overall framework



4. Detailed System Design



			Diagram for the structure of our project

4.1 API Gateway
The API Gateway, built on Cloudwego Hertz, processes HTTP requests, establishing a reliable interface for client communication. On receiving an HTTP request, Hertz extracts the JSON payload from the request body. It then utilizes Cloudwego Kitex's generic call function to serialize these payloads into Thrift binary format, optimizing network transmission speed.

4.2 Load Balancer
The Load Balancer ensures consistent system performance by efficiently distributing workloads across backend RPC servers. It leverages Kitex's inbuilt load balancing strategies, ensuring optimal workload distribution based on server capacity and current system load.

4.3 RPC Server
Powered by Cloudwego Kitex, the backend RPC servers handle Thrift requests forwarded by the Load Balancer. These servers process requests and return responses that are relayed back to the API Gateway via the Load Balancer.

4.4 Our services
Our RPC Servers provide the following services, call, concatenate, encrypt and decrypt.
Our call function greets the user with a hello when the user’s name is inputted. Our concatenate function joins two strings together and returns the result. Our encrypt function is based on the AES encryption algorithm, which is a symmetric key encryption algorithm that provides high levels of security. Our key is a  32-byte array for the AES-256 encryption standard. The method returns the user the encrypted key to the user, which can then be decrypted using our decrypt function.

4.5 Service Registry and Discovery Mechanism
Utilizing Etcd, we've implemented a dynamic configuration, service discovery, and service governance solution. On startup, each instance of our services (API Gateway, Load Balancer, and RPC servers) registers themselves with Etcd. This registration includes the service name, IP address, port, and metadata such as the service's current version.


5. Performance Considerations and Comparisons
   Our system design outperforms traditional solutions by leveraging CloudWeGo Hertz and Kitex's high-performance features.

Hertz and Kitex are high-performance HTTP and RPC frameworks, respectively, developed in Go. Go itself is known for its efficiency, performance, and suitability for building high-performance network servers. Hertz and Kitex are designed to take full advantage of these attributes, providing the necessary infrastructure to build efficient and scalable servers with lower resource consumption.

Compared to traditional data serialization formats like JSON and XML, Thrift provides a more efficient way of encoding data. It uses a binary format that results in smaller message sizes, which means less network bandwidth usage and faster data transmission. Also, because it's a binary protocol, it uses less CPU and memory when encoding and decoding messages, thereby improving the overall performance of the system.

Since the API Gateway, built using Kitex, translates HTTP requests into Thrift binary format before forwarding them to backend RPC servers, it reduces the data size and processing requirements at the server end. This can significantly reduce latency and speed up response times.

The efficient load balancing provided by Kitex ensures that no single server becomes a bottleneck, which is a common issue in traditional systems. By distributing incoming requests evenly across servers, it ensures that all servers share the load, thereby enhancing system performance and reducing the risk of server failure.

Our system promises enhanced scalability, resilience, and performance due to Kitex's efficient load balancing and customizable features.


6. Testing Strategy
   Our comprehensive testing strategy includes unit testing, integration testing , functionality testing and load testing.

Unit tests ensure that individual components or units of our software work as expected. We can use a library like Testify, which provides a set of packages that offer additional functionalities to test our code together with the simple assertion commands in Go.


Integration tests verify that different components of our application work correctly together. For our project, that might include testing the interactions between the API Gateway and the backend RPC servers, or between the services and the service registry.

Functional tests validate that our application works correctly from the user's perspective. For our project, functional testing might involve sending HTTP requests to your API Gateway and verifying that the response is correct.

Load Testing ensures our system behaves and performs as expected under heavy loads. We can use tools like POSTMAN for load testing in Go. We will be running many inputs and comparing the runtime to see how the load balancer performs. This would be particularly important for validating the performance of the load balancing feature of Kitex.


7. Conclusion
   This Software Design Document offers a comprehensive perspective of our API Gateway system that utilizes Cloudwego Hertz and Kitex. The design maintains a focus on performance, scalability, and reliability, which will be confirmed to bring you to the next unclaimed user message.









