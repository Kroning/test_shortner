#!/bin/sh
#
# PROVIDE: webinterface
# REQUIRE: networking
# KEYWORD:

. /etc/rc.subr

name="admin"
rcvar="admin_enable"
goprogram_command="/usr/local/www/go.kroning.ru/www/cmd/main -app admin"
pidfile="/var/run/${name}.pid"
command="/usr/sbin/daemon"
command_args="-P ${pidfile} -r -f ${goprogram_command}"

load_rc_config $name
: ${goprogram_enable:=no}

run_rc_command "$1"

