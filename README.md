<!-- GETTING STARTED -->
## Getting Started

First of all, clone the repository with git.
  ```sh
  git clone https://github.com/MaximeMcs/lbcTest.git
  ```

### Prerequisites


* Install Golang https://go.dev/doc/install
* Install Docker https://docs.docker.com/engine/install/

### Installation

1. Launch docker containers
   ```sh
   docker-compose up
   ```
2. Once everything up, get API Documentation on http://localhost:8000/swagger/index.html
3. Play with the API on Swagger or with your favorite API tool.

<!-- USAGE EXAMPLES -->
## Usage

1. fizzbuzz endpoint: send 5 parameters and get a fizzbuzzed string.
http://localhost:8000/fizzbuzz?int1=3&int2=5&limit=100&str1=toto&str2=test

2. queries endpoint: get the most popular query and its number of hits.
http://localhost:8000/queries

<p align="right">(<a href="#readme-top">back to top</a>)</p>