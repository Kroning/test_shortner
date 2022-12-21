This release uses docker to deploy "all-in-one". Still there are many impovement could be done (like secrets).
It is assumed you have installed everything needed software (like docker).

1. Downloading 1.1.0 release:
```
 git clone --branch v1.1.0 https://github.com/Kroning/test_shortner.git .
```
2. Copy expample configs (and change if you like):
```
cp configs/admin_example.yml configs/admin.yml
cp configs/redirect_example.yml configs/redirect.yml
cp configs/shared_example.yml configs/shared.yml
```
3. Add to /etc/hosts (you can choose another hostnames, but don't forget to change in "docker/\*-nginx.conf" files too)
```
127.0.0.1	go.kroning.ru
127.0.0.1	redirect.kroning.ru
```
4. Optional: change DB passwords at docker-compose.yml ("POSTGRES_DB" and "DB_PASSWORD"x2). ("config/" pass will be replaced for env from docker-compose. It was done just to easy change one file instead of two)
5. Build 2 images (I prefer to build images and use it at docker compose.)
```
sudo docker build -t admin1.1.3:multy -f docker/Dockerfile_admin_multy .
sudo docker build -t redirect1.1.3:multy -f docker/Dockerfile_redirect_multy .
```
6. Run with
```
sudo docker compose up
```

After that you can use browser to see services at go.kroning.ru and redirect.kroning.ru
