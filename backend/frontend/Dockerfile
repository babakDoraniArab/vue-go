# syntax=docker/dockerfile:1.4
FROM --platform=$BUILDPLATFORM node:18 AS development

RUN mkdir /project
WORKDIR /project

COPY . .

# RUN yarn global add @vue/cli
# RUN yarn install
RUN npm install @vue/cli -g
RUN npm install 
ENV HOST=0.0.0.0
# CMD [ "vue", "ui", "--headless" , "--port", "5173",  "--host", "0.0.0.0" ]
CMD ["npm", "run", "dev", "--host", "0.0.0.0"]

FROM development as dev-envs
RUN <<EOF
apk update
apk add git
EOF

RUN <<EOF
addgroup -S docker
adduser -S --shell /bin/bash --ingroup docker vscode
EOF
# install Docker tools (cli, buildx, compose)
COPY --from=gloursdocker/docker / /
CMD ["yarn", "run", "serve"]