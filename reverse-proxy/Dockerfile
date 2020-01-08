FROM nginx:1.12

#  default conf for proxy service
COPY ./example.conf /etc/nginx/conf.d/example.conf

# TLS and HTTPS certificate
COPY ./certs/ /etc/nginx/certs

# Not found 
COPY ./backend-not-found.html /var/www/html/backend-not-found.html


#RUN apt-get update
#RUN apt-get install vim -y
