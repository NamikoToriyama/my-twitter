# My Twitter

2021.5記述
## About
Build@mercariのプログラムで作成。作成期間*5日*。

忘れているため、以下の文章は若干怪しい。

- クライアント `/cmd`
  - htmlとgoのテンプレート機能で作成
- クライアント
  - GCPにデプロイしたのと、GCPのDatastore(おそらく)を使ってDBを実現
  - Docker-compose upでローカルでGCPのデバッグができる
  - バックエンド重視のためフロントの見た目はちょっと微妙です。


以下2020.6記述
## How to run this server

```
$ docker-compose up
```

You can then access your server from `http://localhost:8080`

---------------------------------------
GCP 

```
$ gcloud app deploy
```

You can then access your server from `https://namiko-week6.an.r.appspot.com/`

## About
Week5 created a blog using React, but using REST API in React seemed difficult.

I didn't want to make a blog, so I created an application like Twitter.

Front end is quite simple because I focus on server side development.

## Appeal Point
- Dependency Injection
  - This was very very very hard😂
  - I had to rewrite it to write the test code using mock

- Unit test
  - I wrote all tests in `app` directory

- File structure
  - I made a file structure like go project

- Docker
  - Docker used for working GCP application in local environment.

- Database
  - I used Google Cloud Database(NoSQL).


## Weak point
- Front end
  - Especially, login page.
- repository test
  - `google.golang.org/appengine/aetest` do not work my environment.
  - So, I write pseudo code.


## Dep
```
dep ensure
```

## Docker

Docker stopped working when I added an authentication function
Maybe some library was missing.

```
$ cd docker
$ docker build -t namiko-docker .
$ docker-compose up
```

## Homework
Homework 6 is about extending Homework 5's Blog Assignment by creating your own server!

By 6th of July 16:00 JST, students will be expected to (in Go, JavaScript/TypeScript (Node.js), PHP, or Python):

### Create Endpoints that:

- [x] Respond with all blog pages (GET request)tweet
- [x] Respond with data for a blog page (GET request)
- [x] Handle creation of a blog page (POST request)
- [x] Handle editing of a blog page (POST or PUT request)
- [x] Handle deletion of a blog page (DELETE request)
- [x] Respond with the correct HTTP Status Codes for every above request

### Store data:
- ~~basic: store on file system (e.g. as json files, text files, xml files, csv files, etc.)~~
- [x] stretch: setup a database of your choice

Provide authentication for your blog system (user name + password login):
- ~~basic: sessions~~
- [x] stretch: token based -  `Firebase Auth` 

Ensure your code is properly tested:
- [x] basic: unit tests
- [ ] stretch: integration / e2e tests

### Stretch Goals:
- [x] Deploy your application on a cloud provider - `GCP` 
- [x] Containerize your app with Docker
- [ ]  Setup Kubernetes (service + deployment)
  - we recommend using Minikube here to minimize costs
