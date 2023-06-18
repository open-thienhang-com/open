# Point a domain - sensitive server 10.0.0.0

## Step by step

1. Create a symbolic link inside the /etc/nginx/sites-enabled
   > cd /etc/nginx/sites-enabled
2. Create your own domain name
   > vi <i><<i>domain-name</i>></i>
3. Add content to this screen

```
    server {
            listen 80;
            server_name <domain-name> www.<domain-name>;
            location / {
                proxy_set_header Host $host;
                proxy_set_header X-Real-IP $remote_addr;
                proxy_pass <ip-source-app>;
                add_header Access-Control-Allow-Origin *;
            }
    }
```

- Note:
  > <i><<i>domain-name</i>></i> : tên domain mới, ví dụ: thienhang.com
  > <i><<i>ip-source-app</i>></i> : server build source app, ví dụ : http://10.1.4.69:80

4. Important step, test before reloading domain -> avoid generating errors that affect neighbors.
   > nginx -t
5. Success status then reload nginx
   > nginx -s reload
6. Finally, Point the domain in the hosts file[C:\Windows\System32\drivers\etc]
   > ip <i><<i>domain-name</i>></i>

# Attention:

1. Always have a test [nginx -t] step before reloading [nginx -s reload] to avoid crashing other domains
2. When editing nginx config, if a .save file is generated, <strong>delete .save file</strong>

## Linux editor

- Go to file, create new file

* vi <path_to_file>
* Edit -? Key: [Insert] or i
* Copy -> Key : yy
* Paste -> Key: P
* Delete -> Key: DD
* Undo -> Key: U
* save file -> :w
* save file and exit -> :wq
* exit but not save -> :q!

## Example with Subdomain

- step 1: create or copy a file in this path: /etc/nginx/sites-enabled. in order to create one subdomain quickly and reuse your config

```
    vd:
    cp thienhang.com mitu.thienhang.com // copy
    vi mitu.thienhang.com // create
```

- step 2: If you copy, you will only change server_name

```
    server {
            listen 80;
            server_name mitu.thienhang.com;
            location / {
                proxy_pass <ip-source-app>; //ex: http://loacalhost:9000
            }
    }
```

- step 3:

```
    sudo nginx -t
    sudo systemctl restart nginx
```

# GOOD LUCK!
