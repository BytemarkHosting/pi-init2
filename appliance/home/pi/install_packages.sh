#!/bin/bash

sudo apt update
sudo apt install -y git vim screen python3-pip python-pip

yes | sudo pip3 install ipython
yes | sudo pip install ipython

#### ODRIVE
yes | sudo pip3 install odrive
git clone https://github.com/stanfordroboticsclub/UDP-ODrive.git


#### COMMAND
git clone https://github.com/stanfordroboticsclub/RoverCommand.git

#### GPS
yes | sudo pip3 install adafruit-circuitpython-gps
git clone https://github.com/stanfordroboticsclub/RoverGPS.git

#### UDPCOMMS
git clone https://github.com/stanfordroboticsclub/UDPComms.git
cd UDPComms
sudo python setup.py install
sudo python3 setup.py install
cd ..
