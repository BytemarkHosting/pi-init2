pi-init2
========

Purpose
-------
A program which lets you set up a Raspberry Pi solely by writing to the /boot partition (i.e.  the one you can write from most computers!)

This allows you to distribute a small .zip file to set up a Raspberry  Pi to do anything.  You tell the user to unzip it over the top of the Pi's boot partition - the system can set itself up perfectly on the first boot.

Additionally, once a Raspberry Pi has been set up using [pi-init2](src/projects.bytemark.co.uk/pi-init2/init.go), various sytem files are symlinked back to the /boot, allowing you to reliably edit those "user-serviceable" files from the computer in future.  So e.g. the list of wireless networks and passwords, or other files specific to the kind of appliance you're building.

I keep a small example appliance in this repository which sets the Pi up as a web "kiosk", showing a full screen web browser after it boots up.  It hopefully shows some of the more fiddly stuff.

Trying it out
-------------
From your desktop / laptop:

* Download and write a standard [Raspbian "jessie" SD card](https://www.raspberrypi.org/downloads/raspbian/)
* Unzip the latest release into the /boot partition
* Remove the SD card, and put it into your Pi.

The Raspberry Pi should now boot into a full screen web browser.  The first boot takes 2-5 minutes depending on your network, and which model of Raspberry Pi you use (I tested with models B+ and 2).

You can edit either of these files:

* [appliance/etc/wpa_supplicant/wpa_supplicant.conf](appliance/home/pi/graphical-startup.sh) - for wireless network credentials, if you're using a wireless adapter.
* [appliance/home/pi/graphical_startup.sh](appliance/home/pi/graphical_startup.sh) - to change the URL to load on startup.

Building pi-init2
-----------------
I've included a script called 'build-and-copy' which I use from an Ubuntu system to build the [pi-init2](src/projects.bytemark.co.uk/pi-init2/init.go) program, copy all the appliance files into place, and unmount the card.  Any contributions appreciated.
