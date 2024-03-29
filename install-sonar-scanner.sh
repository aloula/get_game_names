#!/bin/bash
# Source: https://gist.github.com/Chetan07j/85e80d2a17c4d2d75172bf378efe93b9

cd /tmp || exit
echo "Downloading sonar-scanner....."
if [ -d "/tmp/sonar-scanner-cli-4.7.0.2747-linux.zip" ];then
    sudo rm /tmp/sonar-scanner-cli-4.7.0.2747-linux.zip
fi
wget -q https://binaries.sonarsource.com/Distribution/sonar-scanner-cli/sonar-scanner-cli-4.7.0.2747-linux.zip
echo "Download completed."

echo "Unziping downloaded file..."
unzip sonar-scanner-cli-4.7.0.2747-linux.zip
echo "Unzip completed."
rm sonar-scanner-cli-4.7.0.2747-linux.zip

echo "Installing to opt..."
if [ -d "/var/opt/sonar-scanner-4.7.0.2747-linux" ];then
    sudo rm -rf /var/opt/sonar-scanner-4.7.0.2747-linux
fi
sudo mv sonar-scanner-4.7.0.2747-linux /var/opt

echo "Installation completed successfully."

echo "Please add the following to your PATH"

echo "export SONAR_SCANNER_PATH='/var/opt/sonar-scanner-4.7.0.2747-linux/bin'"

echo "export PATH=$PATH:$SONAR_SCANNER_PATH"


