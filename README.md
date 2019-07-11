
#App
export REST_API_DB=""
export REST_API_HOST=""
export REST_API_PORT=""
export REST_API_USER=""
export REST_API_PASSWORD=""

source ~/.profile

#Nginx

server {
	listen 80;
	listen [::]:80;
	 
	server_name meusite.com;

	location / {
		proxy_set_header X-Real-IP $remote_addr;
	    proxy_set_header X-Forwarded-For $remote_addr;
	    proxy_set_header Host $host;
		proxy_pass http://127.0.0.1:12345;		
	} 
	 
}