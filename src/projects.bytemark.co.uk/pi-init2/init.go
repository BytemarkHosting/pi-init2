/* pi-init2
 *
 * A shim to drop onto a Raspberry Pi to write some files to its root 
 * filesystem before giving way to the real /sbin/init.  Its goal is simply 
 * to allow you to customise a RPi by dropping files into that FAT32 /boot 
 * partition, as opposed to either 1) booting it and manually setting it up, or
 * 2) having to mount the root partition, which Windows & Mac users can't easily 
 * do.
 *
 * Cross-compile on Mac/Linux:
 *   GOOS=linux GOARCH=arm go build packages projects.bytemark.co.uk/pi-init2
 *
 * Cross-compile on Windows:
 *   set GOOS=linux
 *   set GOARCH=arm 
 *   go build packages projects.bytemark.co.uk/pi-init2
 */

package main

import "os"
import "fmt"
import "golang.org/x/sys/unix"
import "syscall" // for Exec only

func checkFatalAllowed(desc string, err error, allowedErrnos []syscall.Errno) {
	if (err != nil) {
		errno, ok := err.(syscall.Errno)
		if ok {
			for _, b := range allowedErrnos {
				if b == errno {
					return
				}
			}
		}
		fmt.Println("error " + desc + ":" +err.Error())
		unix.Exit(1);
	}
}

func checkFatal(desc string, err error) {
	checkFatalAllowed(desc, err, []syscall.Errno{})
}

func main() {

	exists := []syscall.Errno{syscall.EEXIST};
	noent  := []syscall.Errno{syscall.ENOENT};

	checkFatal("changing directory", 
		unix.Chdir("/setup"))
	checkFatal("remount rw", 
		unix.Mount("/","/","vfat", syscall.MS_REMOUNT, ""), )
	checkFatalAllowed(
		"making tmp", 
		unix.Mkdir("tmp", 0770),
		exists)
	checkFatalAllowed(
		"making new_root", unix.Mkdir("new_root", 0770), exists)
	checkFatal("mounting tmp", 
		unix.Mount("", "tmp", "tmpfs", 0, ""))
	checkFatal("create device node", 
		unix.Mknod("tmp/mmcblk0p2", 0660 | syscall.S_IFBLK, 179<<8 | 2))
	checkFatal("mounting real root", 
		unix.Mount("tmp/mmcblk0p2", "new_root", "ext4", 0, ""))
	checkFatal("pivoting", 
		unix.PivotRoot("new_root", "new_root/boot"))
	checkFatal("changing to /boot", 
		unix.Chdir("/boot"))
	checkFatal("unmounting tmp", 
		unix.Unmount("setup/tmp", 0))
	checkFatalAllowed(
		"remove wpa_supplicant.conf",
		unix.Unlink("/etc/wpa_supplicant/wpa_supplicant.conf"), noent)
	checkFatalAllowed(
		"remove rc.local",
		unix.Unlink("/etc/rc.local"), noent)
	checkFatal("symlink wpa_supplicant.conf", 
		unix.Symlink(
			"/boot/setup/wpa_supplicant.conf", 
			"/etc/wpa_supplicant/wpa_supplicant.conf"))
	checkFatal("symlink rc.local", 
		unix.Symlink("/boot/setup/rc.local", "/etc/rc.local"))

	// use deprecated API because Exec has been removed from rebuild syscall
	// stuff :-O  Hopefully we will get a hook in Raspbian before this becomes
	// useless.
	checkFatal("exec real init", 
		syscall.Exec("/sbin/init", os.Args, nil))
}
