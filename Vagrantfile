# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
  
  config.vm.box = "precise64"
  config.vm.box_url = "https://vagrantcloud.com/hashicorp/precise64"

  config.vm.provision :shell, :path => "bootstrap.sh"
  config.ssh.forward_agent = true

  #config.vm.network :private_network, ip: "192.168.33.10"
  config.vm.network :forwarded_port, guest: 80, host: 8000

  #config.vm.synced_folder "./src", "/home/vagrant/projects"
  vm_golang_folder = "/go/src/github.com/garethstokes/streamco-application"
  config.vm.synced_folder ".", vm_golang_folder, :id => "vagrant-root", :owner => "vagrant", :group => "vagrant"

   config.vm.provider :virtualbox do |vb|
     # Don't boot with headless mode
     #vb.gui = true
  
     # Use VBoxManage to customize the VM. For example to change memory:
     vb.customize ["modifyvm", :id, "--memory", "1024"]
     #vb.customize ["modifyvm", :id, "--natdnshostresolver1", "on"]
   end
end
