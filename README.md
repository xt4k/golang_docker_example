

# Build image (golangapp with version 7)
docker build -t golangapp:v7 .

# Run container with app
docker run -p 8009:8080 --name goapp-7 golangapp:v7

# List of containers
 docker ps -a

# Enter into container by path (/bin/sh)
 docker exec -ti goapp-7 /bin/sh

# Remove image
docker image rm golangapp:v7 

# List of all images 
docker images -a  

# Stop container (goapp-7)
docker stop goapp-7

# Remove container (goapp-7)
docker rm goapp-7

# Remove image (golangapp with ver 7)
docker image rm golangapp:v7

# Run container
docker run -p 8009:8080 -v "$(pwd):/app/src" --name golangapp-v8 -d golangapp:v8   
# or 
docker run -p 8009:8080 --name goapp-7 golangapp:v7
