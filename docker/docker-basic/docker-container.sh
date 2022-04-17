docker container ls -a

docker container ls

docker container create --name contohredis redis:latest

# If create db container we must set env PASSWORD ROOT, to set env we can use key -e
# ex: -e MARIADB_ROOT_PASSWORD=root

# -p or --publish => forward port from public to internal container
# ex: 3306:3306 => port_yang_bisa_diakses_dari_luar:port_yang_ada_di_serve_container
# ex: 8080:80

# -d untuk melakukan proses di background agar tidak terlihat, using key -d

docker container create --name mariadb-server -p 3306:3306 -e MARIADB_ROOT_PASSWORD=root -d mariadb:10.7
