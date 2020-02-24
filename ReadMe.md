# sdialog

![License](http://img.shields.io/badge/license-Simplified_BSD-blue.svg?style=flat) [![Go Doc](http://img.shields.io/badge/godoc-sdialog-blue.svg?style=flat)](http://pkg.go.dev/github.com/nathanaelle/sdialog/v2) [![Build Status](https://travis-ci.org/nathanaelle/sdialog.svg?branch=master)](https://travis-ci.org/nathanaelle/sdialog) [![Go Report Card](https://goreportcard.com/badge/github.com/nathanaelle/sdialog)](https://goreportcard.com/report/github.com/nathanaelle/sdialog)


## Supported & tested Systemd API

 Supported | Tested | API
-----------|--------|-----
 ✓ | ✓ | [stderr log](https://www.freedesktop.org/software/systemd/man/sd-daemon.html)
 ✓ | ✓ | [sd_notify](https://www.freedesktop.org/software/systemd/man/systemd-notify.html)
 ✓ | ✓ | [watchdog](https://www.freedesktop.org/software/systemd/man/sd_watchdog_enabled.html)
 ✓ | ✓ | accept [socket activation](https://www.freedesktop.org/software/systemd/man/sd_listen_fds.html)
 ? | ✗ | notify [socket activation](https://www.freedesktop.org/software/systemd/man/sd_listen_fds.html)
 ✗ | ✗ | [sd-bus](https://www.freedesktop.org/software/systemd/man/sd-bus.html)

## License

2-Clause BSD

## Todo

  * Documentation
  * More tests
  * [sd-bus](https://www.freedesktop.org/software/systemd/man/sd-bus.html)
