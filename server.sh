#!/bin/sh
set -e
# Start Xvfb, Chrome, and Selenium in the background
export DISPLAY=:21
cd /vagrant

echo "Starting Xvfb ..."
Xvfb :21 -screen 0 1366x768x24 -ac -extension RANDR &

echo "Starting Google Chrome ..."
google-chrome --remote-debugging-port=9222 &

echo "Starting Selenium ..."
nohup java -jar ./selenium-server-standalone-2.50.1.jar &

