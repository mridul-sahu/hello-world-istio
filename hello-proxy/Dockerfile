FROM nginx:1.15.5-alpine
RUN apk add --update curl
COPY nginx.conf /etc/nginx/nginx.conf
RUN rm /etc/nginx/conf.d/default.conf
