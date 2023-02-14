### Feature

- [x] Containerize Application Using Docker
- [x] Integerasi Unit Testing
- [x] And More

## Command

- ### Application Lifecycle

  - Install node modules

  ```sh
  $ go get . || go mod || make goinstall
  ```

  - Build application

  ```sh
  $ go build -o main || make goprod
  ```

  - Start application in development

  ```sh
  $ go run main.go | make godev
  ```

  - Test application

  ```sh
  $ go test main.go main_test.go || make gotest
  ```

* ### Docker Lifecycle

  - Build container

  ```sh
  $ docker-compose build | make dcb
  ```

  - Run container with flags

  ```sh
  $ docker-compose up -d --<flags name> | make dcu f=<flags name>
  ```

  - Run container build with flags

  ```sh
  $ docker-compose up -d --build --<flags name> | make dcubf f=<flags name>
  ```

  - Run container

  ```sh
  $ docker-compose up -d --build | make dcu
  ```

  - Stop container

  ```sh
  $ docker-compose down | make dcd
  ```

### Author

- [David Zhang](https://github.com/devalpha18)
