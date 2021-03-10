# RedisCacheAPI

## Prerequisites
* Go 1.14.3 [Download]
* RedisCache
      ```
      docker run -d --name redis-cache -p 6379:6379 redis:alpine
      ```
      
## How to run project
  1. Install package
      ```sh
      go get
      ```
  2. Run project
      ```sh
      gin -a 80
      ```
      
  [Download]: <https://golang.org/doc/install>
