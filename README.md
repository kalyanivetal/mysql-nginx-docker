# go-microservices-with-mysql-nginx-as-reverse-proxy

How to get to your host machines?

Dependency need: git
To install Command: sudo apt-get install git

git clone https://github.com/shubham1010/mysql-nginx-docker.git


Things to handle?

If you are computer has nginx and mysql service on host machine then follow the following command to stop them respectively otherwise ignore below two lines:

sudo systemctl stop nginx.service
sudo systemctl stop mysql.service

#### To start microservices ####

Give your shell as root previledge then execute following commands:
Dependency need: see the docker.README file

Step-[I]:
	
	1) change directory to mysql-nginx-docker (make your present working directory as 'mysql-nginx-docker')

	2) docker-compose build 
	(it may take time if the images are not on your local machine)

	3) docker-compose up -d 
	(it will start your go-web and mysql containers)

	4) See the which network is formed by command 'docker network ls'
	
	5) Copy that network name.

Step-[II]:


	1) change directory to reverse-proxy (make your present working directory as 'reverse-proxy')

	2) Above copied network name paste it into docker-compose.yml file of current directory at 
	'external:
	   name: copied network name'
	   (Dont use quotes)

	3) docker-compose build

	4) docker-compose up -d


Make sure that all three containers are up(running) by using command: docker ps

### open /etc/hosts file of your host machine using sudo previledge and add example.com along to 127.0.0.1 IP address ###

#### Insert records ####

##### (POST METHOD) #####
Dependency:
curl => command to install
sudo apt-get install curl (ubuntu/debian)

curl --insecure -H "Content-Type: application/json" -X POST -d '{"id":"1","date":"2020/01/10","name":"shubham1010"}' https://example.com/endpoint


#### RestAPI's on browser ####

##### (GET METHODS) #####
visit https://example.com/endpoint => (it will display all records in the table)

visit https://example.com/endpoint/id => (id is an integer, it will display the records which has id in table if exists)

#### To stop and remove ALL microservices ####

docker-compose down => stop and then removes all up(running) containers of respective docker-compose.yml file.

(Note: use docker-compose down command within the respective directory which contains docker-compose.yml file)

docker rm -f $(docker ps -aq) => to remove all container forcefully



###### IGNORE THE mysqlCRUD.go FILE ######
