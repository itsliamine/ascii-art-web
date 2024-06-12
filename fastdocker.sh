sudo docker build -f Dockerfile -t asciiartwebimage .
sudo docker container run -p 8080 --detach --name asciiweb asciiartwebimage
sudo docker exec -it asciiweb /bin/bash