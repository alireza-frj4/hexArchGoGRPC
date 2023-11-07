# hexArchGoGRPC

The project is a simple microservice using gRPC in Golang that performs mathematical operations. It is structured using the Hexagonal architecture and connects to a MySQL database. Additionally, unit tests have been written to ensure the functionality of the microservice.

## Technologies Used

- Golang
- gRPC
- MySQL
- Docker
- Docker Compose

## Project Structure

The project follows the Hexagonal architecture, separating the core logic from external concerns such as databases and network communication. The structure ensures modularity and flexibility in adding new features or changing existing ones.

## How to Run

To run this microservice, follow these steps:

1. Clone the repository from GitHub:

 - git clone <https://github.com/alireza-frj4/hexArchGoGRPC.git>

2. install dependency 

 -go mod tidy

3. for run on docker-compose

 - docker compose up --build