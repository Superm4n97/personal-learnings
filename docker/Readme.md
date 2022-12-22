### Download Docker
Docker official site: https://docs.docker.com/engine/install/ubuntu/

### Learn Docker
* Official docs: https://docs.docker.com/get-started/02_our_app/
* Sysdevbd: https://github.com/sysdevbd/sysdevbd.github.io/tree/master/containers/docker
* Dockerfile reference: https://docs.docker.com/engine/reference/builder/

### Development Workflow
A dockerfile is added to the application in order to dockerized it. Dockerfile is a plain text instruction which is used to package up the application into an image.
* A cut down operating system
* A runtime environment
* Application files
* Third party libraries
* Environment variables

### Notes
* `docker build -t <dockerhub-account>/<applicationName>:<version-tag>` - your image from `Dockerfile`, `-t` is used for tag (image name). 
* `docker images` will show the full list of the images on your device.
* `docker run -it <imageName>` - run the imageName image
* `dokcer push <dokcerhub-account>/<applicationName>:<version-tag>` - will push the image into the dockerhub-account according to the applicationName and version.