# pull official base image
FROM node:12.12.0-alpine

# set working directory
WORKDIR /app

# add `/app/node_modules/.bin` to $PATH
ENV PATH /app/node_modules/.bin:$PATH

# install app dependencies
COPY yarn.lock ./

RUN yarn
RUN yarn global add react-scripts

# add app
COPY . ./

# start app
CMD ["yarn", "start"]

EXPOSE 3000