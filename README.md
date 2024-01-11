https://qiita.com/takehanKosuke/items/b809d8dfe8195b746278

docker containerを起動した後は以下のURLで戻り値がくる
```
http://localhost/ping
http://localhost/products
```

# 注意点

1. go.modのmoduleはproject folder名で書くこと

$ docker commit {対象のコンテナID(コンテナ名)} {作成したいイメージ名}
docker commit 4fdc72e6e43d 075c3b8587ff

docker pull 075c3b8587ff

docker up

docke docker-compose build --no-cache mysql

docker container rm go_docker_ec2-db 

