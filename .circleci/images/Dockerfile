FROM golang:1.9.4

RUN apt-get update && apt-get -y upgrade
#install python
RUN apt-get install -y python-pip
RUN chsh root -s /bin/bash
#install elastic beanstalk command line
#RUN pip install awsebcli --upgrade --ignore-installed six
RUN pip install awsebcli --upgrade --user
RUN echo PATH=~/.local/bin:$PATH >> ~/.bashrc 
#install node
RUN curl -sL https://deb.nodesource.com/setup_8.x | bash
RUN apt-get install -y nodejs
RUN npm install -g gulp
# #install glide
RUN curl https://glide.sh/get | sh
RUN go get -u github.com/pressly/goose/cmd/goose    