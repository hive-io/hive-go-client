#!/bin/bash

echo "127.0.0.1 cma.hiveio.internal" >> /etc/hosts

cp -r /home/admin1/.ssh /root/
chown -R root:root /root/.ssh
