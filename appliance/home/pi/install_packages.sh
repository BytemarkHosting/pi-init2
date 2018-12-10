#!/usr/bin/env bash

if [ $(mount | grep mmcblk0p2 | grep -o 'rw' || echo ro) = "ro" ]; then
        echo "Please make file system read-write and rerun"
        exit;
fi

if [ $(hostname) = "raspberrypi" ]; then
        read -p 'Current hostname: raspberrypi. Please change hostname to: ' new_hostname
        [ ! -z "$new_hostname" ] && echo "$new_hostname" | sudo tee -a /etc/hostname
        [ ! -z "$new_hostname" ] && sudo hostnamectl set-hostname "$new_hostname"
fi
read -p 'Enter the current date [YYYY-MM-DD HH:MM] or hit enter to skip: ' datetime
[ ! -z "$datetime" ] && sudo date -s "$datetime"

sudo apt update
sudo apt install -y git vim screen python3-pip python-pip

yes | sudo pip3 install ipython
yes | sudo pip install ipython

#### UDPCOMMS
git clone https://github.com/stanfordroboticsclub/UDPComms.git
sudo bash UDPComms/install.sh

#### ODRIVE
git clone https://github.com/stanfordroboticsclub/RoverODrive.git
sudo bash RoverODrive/install.sh

#### COMMAND
git clone https://github.com/stanfordroboticsclub/RoverCommand.git
sudo bash RoverCommand/install.sh

#### GPS
git clone https://github.com/stanfordroboticsclub/RoverGPS.git
sudo bash RoverGPS/install.sh
