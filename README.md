# Install
- Using docker for running redis

```
docker run --name some-redis -d redis
```

# Jump into container
- Using container what have created by the above command. We can attach into container

```
docker exec -it c1b9d9fc4919 /bin/bash
```

## Interactor with redis server by redis-cli

- Run this command :
```
$ redis-cli
```

- Get all keys in redis
```
> keys *
```

- Set key to hold the string value
```
set name hogehoge
```

- 