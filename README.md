PRI-Setup (WIP)
========

Purpose
-------
This repository allows you to set up a Raspberry Pi solely by writing to the /boot partition (i.e.  the one you can write from most computers!) in a repeatable manner. This allows you to distribute a small .zip file to set up a Raspberry Pi to do anything.  You tell the user to unzip it over the top of the Pi's boot partition - the system can set itself up perfectly on the first boot.

This is done using [pi-init2](src/projects.bytemark.co.uk/pi-init2/init.go). You can read more about how it works behind the scenes [here](https://blog.bytemark.co.uk/2016/01/04/setting-up-a-raspberry-pi-perfectly-on-the-first-boot)

Additionaly pi-init2 various system files are symlinked back to the /boot, allowing you to reliably edit those "user-serviceable" files from the computer in future. 

Module Specification
-------------
The goal of this is to section is to outline a specification modules installed on the Pi should have. This allows them to be easily installed/enabled in a predictable way, even by someone unfamilar with the module in question. Many of those modules will interact with each other over the rover network (with staticly assigned IP's in range 10.0.0.X) using the [UDPComms Library](https://github.com/stanfordroboticsclub/UDPComms).

Specifications:

- Each module shall be a single git repository
- Each module shall be downloadable with `git clone (address)`
- Each module shall contain a `install.sh` script which will prepare the module to be used including installing any requirements and symlinking the service files to the correct places 
- Each module shall be documeted using the `README.md` and the topics is publishes and subscribes to are listed on the [CS Comms System](https://docs.google.com/spreadsheets/d/1pqduUwYa1_sWiObJDrvCCz4Al3pl588ytE4u-Dwa6Pw/edit?usp=sharing) document
- Each executable to be used on the rover shall have a `name.service` script associated with it. Read the [intro to systemd](https://www.devdungeon.com/content/creating-systemd-service-files) and leanr more about writing [serive files](https://www.digitalocean.com/community/tutorials/understanding-systemd-units-and-unit-files#anatomy-of-a-unit-file) This allows them to be manipulated as a systemd service using the following commands:

| Command | Descripion |
|---------|------------|
| `sudo systemctl status name` | tell us what the service is doing right now|
|`sudo systemctl start name` | start the service right now|
|`sudo systemctl stop name` | stop the service right now|
|`sudo systemctl disable name` | stop the service from starting on boot|
|`sudo systemctl enable name` | make the service start on boot|
|`journalctl -u name` | display the output of the service |

For an example of module take a look at [RoverGPS](https://github.com/stanfordroboticsclub/RoverGPS)


Actually doing it
-------------
From your desktop / laptop:

* Download and write a standard [Raspbian Stretch Lite](https://www.raspberrypi.org/downloads/raspbian/) SD card. We recomend using [etcher](https://www.balena.io/etcher/) to flash the card
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
- Use `sudo date -s "11/26/2018 03:38"` to update the current datetime (replace with current time). This is to prevent the Pi from thinking the certificates used to download the packages are from the future.
- Run `sudo ./install_packages.sh` to install packages

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
