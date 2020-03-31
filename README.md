RPI-Setup (WIP)
========

Purpose
-------
This repository allows you to set up a Raspberry Pi solely by writing to the /boot partition (i.e.  the one you can write from most computers!) in a repeatable manner. This allows you to distribute a small .zip file to set up a Raspberry Pi to do anything.  You tell the user to unzip it over the top of the Pi's boot partition - the system can set itself up perfectly on the first boot.

This is done using [pi-init2](src/projects.bytemark.co.uk/pi-init2/init.go). You can read more about how it works behind the scenes [here](https://blog.bytemark.co.uk/2016/01/04/setting-up-a-raspberry-pi-perfectly-on-the-first-boot)

Additionaly pi-init2 various system files are symlinked back to the /boot, allowing you to reliably edit those "user-serviceable" files from the computer in future. 

Actually doing it
-------------
From your desktop / laptop:

* Download and write a standard Raspbian Buster Lite SD card. Use [this version](https://slack-files.com/T0RAWRCGY-FQG7WTSBH-eb9549ed22) so everyone is using the same version. We recomend using [etcher](https://www.balena.io/etcher/) to flash the card
* Download the latest release of this repository into the /boot partition. Unzip and move all the files into the /boot folder (replace any files that conflict so the repository's version overwrites the original version). Delete the zip file and now empty folder.
* Remove the SD card and put it into your Pi.

The Raspberry Pi should now boot and set everything up for development. 

Getting internet access
-------------
This script will make so the RPi automatically wants to connect the Stanford network. Initially it won't be able to do that as it is not yet authenticated to do it. To set that up:

- Plug your Pi in to power (over the onboard micro USB port). Either plug a monitor and keyboard into the Pi or SSH into it using your laptop over Ethernet. Log in to the Pi. In the welcome message that comes after the login line, look for the Pi's **MAC address**, which will appear under the line that says "wireless Hardware MAC address". Note that address down.
- Use another computer to navigate to [iprequest.stanford.edu](http://iprequest.stanford.edu).
- Log in using your Stanford credentials.
- Follow the on-screen instructions to add another device:
   - **First page:** Device Type: Other, Operating System: Linux, Hardware Address: put Pi's MAC address
   - **Second page:** Make and model: Other PC, Hardware Addresses Wired: delete what's there, Hardware Addresses Wireless: put Pi's MAC address
- Confirm that the Pi is connected to the network:
   - Wait for an email (to your Stanford email) that the device has been accepted
   - `sudo reboot` on the Pi
   - After it's done rebooting, type `ping www.google.com` and make sure you are receiving packets over the network

Getting started with the Pi
-------------
- Configure your computer to access the Rover network:
	- Go to your network settings for the interface you wish to use (ethernet/wifi)
	- Change your Configure IPv4: Manually
	- Change your IP Address: 10.0.0.X (see the [CS Comms System](https://docs.google.com/spreadsheets/d/1pqduUwYa1_sWiObJDrvCCz4Al3pl588ytE4u-Dwa6Pw/edit?usp=sharing) document for what X to use)
	- Change your Subnet Mask: 255.255.255.0
	- Leave the Router blank
	- After disconnecting from the Rover network remeber to return those settings to what they orignially were, otherwise your internet on that interface won't work
- Ssh into the pi using `ssh pi@10.0.0.10` from your computer
- Type `rw` to enter read-write mode. Confirm that the terminal prompt ends with `(rw)` instead of `(ro)`
- Run `sudo ./install_packages.sh` to install packages
	- If the IP is still 10.0.0.10 you will be prompted to change it
	- If the hostname is still raspberry you will be prompted to change it
	- You will be asked to enter the current time and date. This is needed so that certificates don't get marked as expired
- Once you have done this you can ssh into the pi using the [rover command](https://github.com/stanfordroboticsclub/UDPComms). by typing `rover connect hostname`

What this repo does
-------------
- Enables ssh (only on ethernet)
- Sets the Pi to connect to the rover network (10.0.0.X) over ethernet)
- Expands the SD card file system
- Sets the file system up as read only
- Prepares to connect to Stanford WiFi (see above for details)
- Gives the script to install tools and repos needed for development


Building pi-init2
-----------------
This repo inculdes the pi-init2 binary and there shouldn't be any reason to recompile it. If you need to there is a included Makefile
