FROM circleci/cci-demo-docker-primary:0.0.2

RUN apt-get update && apt-get -y upgrade
#install python
RUN apt-get install -y python-pip
RUN chsh root -s /bin/bash
#install elastic beanstalk command line
RUN pip install awscli --upgrade --ignore-installed six
#install node
RUN curl -sL https://deb.nodesource.com/setup_8.x | bash
RUN apt-get install -y nodejs
#install go
RUN apt-get update && apt-get install -y --no-install-recommends \
		g++ \
		gcc \
		libc6-dev \
		make \
		pkg-config \
	&& apt-get install -y zlib1g-dev \
	&& apt-get install -y zip \
	&& apt-get install -y unzip \
	&& rm -rf /var/lib/apt/lists/*
ENV GOLANG_VERSION 1.9.4
ENV GOLANG_DOWNLOAD_URL https://golang.org/dl/go$GOLANG_VERSION.linux-amd64.tar.gz
ENV GOLANG_DOWNLOAD_SHA256 15b0937615809f87321a457bb1265f946f9f6e736c563d6c5e0bd2c22e44f779
RUN curl -fsSL "$GOLANG_DOWNLOAD_URL" -o golang.tar.gz \
	&& echo "$GOLANG_DOWNLOAD_SHA256  golang.tar.gz" | sha256sum -c - \
	&& tar -C /usr/local -xzf golang.tar.gz \
	&& rm golang.tar.gz
# Set environment
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
# Create project in Docker
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
RUN mkdir -p "$GOPATH/src"

#install glide
RUN curl https://glide.sh/get | sh