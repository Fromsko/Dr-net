@echo off

pip config set global.index-url http://pypi.douban.com/simple/

pip config set install.trusted-host pypi.douban.com

python -m pip install -r requirement.txt -i http://pypi.douban.com/simple/

exit