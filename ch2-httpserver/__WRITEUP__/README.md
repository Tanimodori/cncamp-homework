# Ch3. Dockerfile

## Build image

```bash
sudo docker build -t tanimodori/testserver .
```

## Run image

```bash
sudo docker run --publish 6060:8888 --name test --rm tanimodori/testserver
```

![Build and Run](CREATE_AND_RUN_IMAGE.png)

## Visit image

```bash
wget -qO- localhost:6060/healthz
```

![Visit image](SERVER_CALL.png)

# Show image network

```bash
sudo docker inspect --format {{.State.Pid}} test
sudo nsenter --target <PID> -n ip a
ip a
```

![Show image network](IMAGE_NET_INSPECT.png)

## Publish image

```bash
sudo docker push tanimodori/testserver
```

![Publish image](IMAGE_PUBLISH.png)