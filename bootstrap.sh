#!/usr/bin/env bash

apt-get update

# vim
if [ ! -f "/usr/bin/vim" ]; then
  apt-get install -y vim
fi

# git
if [ ! -f "/usr/bin/git" ]; then
  apt-get install -y git
fi

# bzr
if [ ! -f "/usr/bin/bzr" ]; then
  apt-get install -y bzr
fi

# mecurial
if [ ! -f "/usr/bin/hg" ]; then
  apt-get install -y mercurial
fi

# curl
if [ ! -f "/usr/bin/curl" ]; then
  apt-get install -y curl
fi

# nginx
if [ ! -f "/usr/sbin/nginx" ]; then
  apt-get install -y nginx

  cat << EOF > /etc/nginx/sites-available/default
server {
  #listen   80; ## listen for ipv4; this line is default and implied
  #listen   [::]:80 default ipv6only=on; ## listen for ipv6

  root /go/src/github.com/garethstokes/streamco-application/public;
  index index.html;
  server_name localhost;

  location / {
    proxy_pass http://127.0.0.1:8080;
    proxy_redirect  off;
    proxy_set_header   Host             \$host;
    proxy_set_header   X-Real-IP        \$remote_addr;
    proxy_set_header   X-Forwarded-For  \$proxy_add_x_forwarded_for;
  }
}
EOF

  service nginx restart
fi

# mysqld
if [ ! -f "/usr/sbin/mysqld" ]; then
  export DEBIAN_FRONTEND=noninteractive
  apt-get install -y mysql-server
fi

# build golang from source
if [ ! -d "/home/vagrant/go" ]; then
  echo 'download the golang src code'

  wget https://storage.googleapis.com/golang/go1.2.2.src.tar.gz
  
  echo 'extract'
  tar zxvf go1.2.2.src.tar.gz

  echo 'and now compile'
  cd go/src/
  ./all.bash

  # install into /usr/local/bin
  cd ../bin
  cp go /usr/local/bin/go
  cp gofmt /usr/local/bin/gofmt

  # configure the gopath to ~/projects/
  echo "" >> /home/vagrant/.bashrc
  echo "export GOPATH=/go/" >> /home/vagrant/.bashrc
  
  echo "" >> /home/vagrant/.bashrc
  echo "export PATH=$PATH:/go/bin/" >> /home/vagrant/.bashrc
  
  echo "" >> /home/vagrant/.bashrc
  echo "alias sw='cd /go/src/github.com/garethstokes/streamco-application/'" >> /home/vagrant/.bashrc
fi

chown -R vagrant:vagrant /go/

echo "bootstrap done"
