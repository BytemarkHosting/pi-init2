PRI-Setup (WIP)
========

Purpose
-------
This repository allows you to set up a Raspberry Pi solely by writing to the /boot partition (i.e.  the one you can write from most computers!) in a repeatable manner.

This allows you to distribute a small .zip file to set up a Raspberry Pi to do anything.  You tell the user to unzip it over the top of the Pi's boot partition - the system can set itself up perfectly on the first boot.

This is done using [pi-init2](src/projects.bytemark.co.uk/pi-init2/init.go). You can read more about how it works behind the scenes [here](https://blog.bytemark.co.uk/2016/01/04/setting-up-a-raspberry-pi-perfectly-on-the-first-boot)


Additionaly pi-init2 various system files are symlinked back to the /boot, allowing you to reliably edit those "user-serviceable" files from the computer in future. 


When you will be developing various subsystems that we would want to setup reliably you can fork this repository and make it automatically setup that subsystem's Pi


Actually doing it
-------------
From your desktop / laptop:

* Download and write a standard [Raspbian "jessie" SD card](https://www.raspberrypi.org/downloads/raspbian/)
* Unzip the latest release of this repository into the /boot partition
* Remove the SD card and put it into your Pi.

The Raspberry Pi should now boot and set everything up for development


Getting internet access
-------------
This script will make so the RPi automatically wants to connect the Stanford network. Initially it won't be able to do that as it is not yet authenticated to do it. To set that up:

- Use another device to navigate to [iprequest.stanford.edu](http://iprequest.stanford.edu) 
- Log in using your Stanford credentials
- Follow the on-screen instructions to add another device
- As the wireless MAC address give the string which appear in ~/MAC.txt on the RPi


What this repo does (WIP)
-------------
- Enables ssh (only on ethernet)
- Prepares to connect to Stanford WiFi (see above for details)
- Instals tools needed for development
- Removes some unnecessary packages
- Sets the file system up as read only


Building pi-init2
-----------------
This repo inculdes the pi-init2 binary and there shouldn't be any reason to recompile it. If you need to there is a included Makefile
