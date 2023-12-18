import os
import requests
import sys

# user = os.getlogin()
# password = ""
g5k_auth = None

site_id = "lyon"

nodes = sys.argv[1]

api_job_url = f"https://api.grid5000.fr/stable/sites/{site_id}/jobs"

payload = {
    "resources": f"nodes={nodes}",
    "command": "cd makefiles;make premier;make matrix;cd ..;export PATH=$PATH:$(cat ~/.path);cd dismake/server;go run main.go",
    "name": "servers_start"
}
job = requests.post(api_job_url, data=payload, auth=g5k_auth).json()
job_id = job["uid"]

print(f"Server job submitted ({job_id})")

hosts = []
while (True):
    job = requests.get(api_job_url+f"/{job_id}", auth=g5k_auth).json()
    if job["state"] == "running":
        hosts = job["assigned_nodes"]
        break

print(hosts)

for i in range(len(hosts)):
    hosts[i] += ":50051"
hosts_string = " ".join(hosts)


payload = {
    "resources": "nodes=1",
    "command": f"export PATH=$PATH:~/go/bin;make -C ~/dismake/client;~/dismake/client/client -server \"{hosts_string}\" ~/makefiles/premier/Makefile;~/dismake/client/client -server \"{hosts_string}\" ~/makefiles/premier/Makefile;~/dismake/client/client -server \"{hosts_string}\" ~/makefiles/premier/Makefile;~/dismake/client/client -server \"{hosts_string}\" ~/makefiles/matrix/Makefile;~/dismake/client/client -server \"{hosts_string}\" ~/makefiles/matrix/Makefile;~/dismake/client/client -server \"{hosts_string}\" ~/makefiles/matrix/Makefile",
    "name": "client_start"
}
client_job = requests.post(api_job_url, data=payload, auth=g5k_auth).json()
client_job_id = client_job["uid"]

print(f"Client job submitted ({client_job_id})")

client_hosts = []
while (True):
    client_job = requests.get(api_job_url+f"/{client_job_id}", auth=g5k_auth).json()
    if client_job["state"] == "running":
        client_hosts = client_job["assigned_nodes"]
        break

print(client_hosts)

while True:
    state = requests.get(api_job_url+f"/{client_job_id}", auth=g5k_auth).json()["state"]
    if state == "terminated":
        requests.delete(api_job_url+f"/{client_job_id}", auth=g5k_auth)
        print("client job deleted")
        requests.delete(api_job_url+f"/{job_id}", auth=g5k_auth)
        print("servers job deleted")
        break
