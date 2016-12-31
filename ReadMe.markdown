# sdialog

![License](http://img.shields.io/badge/license-Simplified_BSD-blue.svg?style=flat) [![Go Doc](http://img.shields.io/badge/godoc-sdialog-blue.svg?style=flat)](http://godoc.org/github.com/nathanaelle/sdialog) [![Build Status](https://travis-ci.org/nathanaelle/sdialog.svg?branch=master)](https://travis-ci.org/nathanaelle/sdialog)

## What is sdialog ?

Binding for :
  * Systemd
  * Shesha

## Supported & tested Systemd API

 Supported | Tested | API
-----------|--------|-----
 [x] | [x] | [stderr log](https://www.freedesktop.org/software/systemd/man/sd-daemon.html)
 [x] | [x] | [sd_notify](https://www.freedesktop.org/software/systemd/man/systemd-notify.html)
 [x] | [x] | [watchdog](https://www.freedesktop.org/software/systemd/man/sd_watchdog_enabled.html)
 [x] | [ ] | [socket activation](https://www.freedesktop.org/software/systemd/man/sd_listen_fds.html)
 [ ] | [ ] | [sd-bus](https://www.freedesktop.org/software/systemd/man/sd-bus.html)

## Shesha extensions

  * none

## License

2-Clause BSD

## Todo

  * Documentation
  * [sd-bus](https://www.freedesktop.org/software/systemd/man/sd-bus.html)
