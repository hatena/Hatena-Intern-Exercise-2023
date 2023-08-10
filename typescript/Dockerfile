FROM node:16

WORKDIR /app

COPY package.json yarn.lock ./
RUN yarn install

COPY tsconfig.json babel.config.js webpack.config.js jest.config.js .eslintrc.json .prettierrc.json .prettierignore ./
COPY index.html ./
COPY src ./src
