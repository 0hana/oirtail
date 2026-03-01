# oirtaiL - esicrexe weivretni

Building the image
--------------------------------------------------------------------------------
```sh
docker build -t oirtail ./
```


Running the container
--------------------------------------------------------------------------------

### Interactive mode (for dev/debug)
```sh
docker run --rm -itp 80:80 oirtail
```
- `--rm`     : remove the container when stopped
- `-i`       : interactive mode
- `-t`       : allocate a pseudo-terminal
- `-p 80:80` : publish/port map host port 80 to container port 80


### Detached mode (for deployment)
```sh
docker run -d -p 80:80 --name evil-oirtail oirtail
```
- `-d`                  : detached mode
- `-p 80:80`            : publish/port map host port 80 to container port 80
- `--name evil-oirtail` : name container instance "evil-oirtail" for easy ref
