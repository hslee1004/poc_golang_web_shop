#!/bin/bash
#
# JENKINS_HOME: /var/lib/jenkins
# WORKSPACE:    /var/lib/jenkins/jobs/0_poc_jupiter/workspace
# pwd:          /disks/sdb/jenkins-workspace/0_poc_jupiter/workspace
# sh:           /var/lib/jenkins/jobs/0_poc_jupiter/workspace/src/jupiter/build.sh
#

check=$1

echo JENKINS_HOME:$JENKINS_HOME
echo WORKSPACE:   ${WORKSPACE}
echo pwd:         $(pwd) 
cd ${WORKSPACE}/src/jupiter/
set -x
export GOROOT=/usr/local/go
export GOBIN="$GOROOT/bin"
export GOPATH="${WORKSPACE}"
export PATH=$PATH:$GOBIN

if [ "$check" = "check" ]; then
	make check
fi
make -B
make install env=prod
