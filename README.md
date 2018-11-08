PRI-Setup (WIP)
========

Purpose
-------
This repository allows you to set up a Raspberry Pi solely by writing to the /boot partition (i.e.  the one you can write from most computers!) in a repeatable manner.

This allows you to distribute a small .zip file to set up a Raspberry Pi to do anything.  You tell the user to unzip it over the top of the Pi's boot partition - the system can set itself up perfectly on the first boot.

This is done using [pi-init2](src/projects.bytemark.co.uk/pi-init2/init.go). You can read more about how it works behind the scenes [here](https://blog.bytemark.co.uk/2016/01/04/setting-up-a-raspberry-pi-perfectly-on-the-first-boot)


Additionaly pi-init2 various system files are symlinked back to the /boot, allowing you to reliably edit those "user-serviceable" files from the computer in future. 



Actually doing it
-------------
From your desktop / laptop:

* Download and write a standard [Raspbian "jessie" SD card](https://www.raspberrypi.org/downloads/raspbian/). We recomend using [etcher](https://www.balena.io/etcher/) to flash the card
* Unzip the latest release of this repository into the /boot partition
* Remove the SD card and put it into your Pi.

The Raspberry Pi should now boot and set everything up for development. 

If you want the chagnes you make to be under version control instead of unzipping you can do the following (EDIT: DON'T DO THIS, it messes with file permissions and is a mess to clean up. Use the inculded install_packages.sh script)

```
$cd /Volumes/boot
$git init
$git remote add origin https://github.com/stanfordroboticsclub/RPI-Setup.git
$git fetch --all
$git reset --hard origin/master
```
 
Getting internet access
-------------
This script will make so the RPi automatically wants to connect the Stanford network. Initially it won't be able to do that as it is not yet authenticated to do it. To set that up:

- Use another device to navigate to [iprequest.stanford.edu](http://iprequest.stanford.edu) 
- Log in using your Stanford credentials
- Follow the on-screen instructions to add another device (The Pi's MAC address appears as part of the welcome screen)


What this repo does
-------------
- Enables ssh (only on ethernet)
- Sets the Pi to connect to the rover network (10.0.0.X) over ethernet)
- Expands the SD card file system
- Sets the file system up as read only
- Prepares to connect to Stanford WiFi (see above for details)
- Gives the script to instal tools and repos needed for development


Building pi-init2
-----------------
This repo inculdes the pi-init2 binary and there shouldn't be any reason to recompile it. If you need to there is a included Makefile
