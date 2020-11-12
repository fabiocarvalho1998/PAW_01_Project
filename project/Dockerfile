FROM node:10.15.0

# install simple http server for serving static content
RUN npm install -g http-server

# make the 'app' folder the current working directory
RUN mkdir /usr/src/app
WORKDIR /usr/src/app

# add `/usr/src/app/node_modules/.bin` to $PATH
ENV PATH /usr/src/app/node_modules/.bin:$PATH

# install and cache app dependencies
COPY package.json /usr/src/app/package.json
RUN npm install
RUN npm install -g @vue/cli
RUN npm install -g vue-browser-geolocation
RUN npm install -g vue2-google-maps

# start app
CMD ["npm", "run", "serve"]