Just trying Go (Echo) [Backend] + Next.Js (+ Tailwind) [Frontend] + MySQL [DB] + Docker Compose [Containerization]  

If you want to try, please make sure docker and docker compose are installed.  
1. git clone this repo
2. checkout to branch `dockerReadyV1`
2. run this command  
  Windows:  
  ``` docker compose -f docker-compose.yml up -d ```    
  Linux:  
  ``` docker-compose -f docker-compose.yml up -d ```  
  

Now available with kubernetes!  
thx to https://kompose.io/
Here's how you can do it:
1. build the image  
  ``` docker-compose build ```
2. run it  
  ``` kubectl apply -f kubemanifests.yaml --validate=false ```

Now with ingress!
1. build the image  
  ``` docker-compose build ```
2. Installing Ingress Controller (Nginx LoadBalancer)
  ``` kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.0.0/deploy/static/provider/cloud/deploy.yaml ```
3. run it  
  ``` kubectl apply -f kubemanifests.yaml --validate=false ```
  
Postman : https://documenter.getpostman.com/view/17343050/U16qKPRz
  
just fire it up with the command above and make sure port 80 and 443 are available   


If You want to try develop the apps, please refer to first read the file `docker-compose-dev.yml` and then run the docker compose on that file
