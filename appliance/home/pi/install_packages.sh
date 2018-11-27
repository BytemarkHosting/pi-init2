#!/bin/bash

read -p 'Please enter the current datetime in YYYY-MM-DD HH:MM format: ' datetime
sudo date -s "$datetime"

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
