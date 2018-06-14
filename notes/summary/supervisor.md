# Supervisord 配置
- install
    ```bash
    prod@ubuntu:~$ sudo pip install supervisord
    ```

- configure
    ```bash
    # 创建配置文件路径
    prod@ubuntu:~$ mkdir -p /etc/supervisor/conf.d/

    # 添加 supervisord 的配置文件
    prod@ubuntu:~$ sudo cat /etc/supervisor/supervisord.conf 
    ; supervisor config file

    [unix_http_server]
    file=/var/run/supervisor.sock   ; (the path to the socket file)
    chmod=0700                       ; sockef file mode (default 0700)

    [inet_http_server]
    port = 0.0.0.0:9009
    username = admin
    password = 123456

    [supervisord]
    logfile=/var/log/supervisor/supervisord.log ; (main log file;default $CWD/supervisord.log)
    pidfile=/var/run/supervisord.pid ; (supervisord pidfile;default supervisord.pid)
    childlogdir=/var/log/supervisor  ; ('AUTO' child log dir, default $TEMP)

    ; the below section must remain in the config file for RPC
    ; (supervisorctl/web interface) to work, additional interfaces may be
    ; added by defining them in separate rpcinterface: sections
    [rpcinterface:supervisor]
    supervisor.rpcinterface_factory = supervisor.rpcinterface:make_main_rpcinterface

    [supervisorctl]
    serverurl=unix:///var/run/supervisor.sock ; use a unix:// URL  for a unix socket

    ; The [include] section can just contain the "files" setting.  This
    ; setting can list multiple files (separated by whitespace or
    ; newlines).  It can also contain wildcards.  The filenames are
    ; interpreted as relative to this file.  Included files *cannot*
    ; include files themselves.

    [include]
    files = /etc/supervisor/conf.d/*.ini
    ```

- running

    ```bash
    # 验证Running 可以通过登录配置的网页 127.0.0.1:9009
    prod@ubuntu:~$ sudo supervisord -c /etc/supervisor/supervisord.conf 
    ```

# Supervisorctl

- config 
    创建 touch /etc/supervisor/conf.d/bootnode.ini
    ```ini
    [program:bootnode]
    command= /home/prod/ethereum/geth-alltools/bootnode -nodekey /home/prod/ethereum/poa/nodes/bootnode/boot.key -addr :30310
    autorestart=true
    autostart=true
    stdout_logfile=/home/prod/ethereum/poa/log/bootnode.log
    redirect_stderr=true
    stopsignal=QUIT
    ```

- update
    ```bash
    prod@ubuntu:~/ethereum/poa$ sudo supervisorctl update bootnode
    bootnode: added process group

    prod@ubuntu:~/ethereum/poa$ sudo supervisorctl status
    bootnode                         RUNNING   pid 23898, uptime 0:00:09
    ```

- add nodes
    ```ini
    [program:poanode1]
    command= /home/prod/ethereum/geth-alltools/geth --syncmode 'full' --rpc --rpcport "8541" --rpcaddr "0.0.0.0" --datadir /home/prod/ethereum/poa/nodes/node1 --port "30301" --rpccorsdomain "*" --rpcapi "personal,db,eth,net,web3,admin,txpool,miner,clique" --networkid 1034 --unlock 0x5dab1b9d9da91d77c3b130785c1117507e252412 --password /home/prod/ethereum/poa/nodes/node1/password --bootnodes enode://8ea516886bfaa84dae6fec94c6eb4d9babbe4a986a5d21e3e089b1ec6be1ad704281a9c6f51fd1bfaaa611de9228c37f9e945ffe3a49e6ffbc055e55d16373d6@127.0.0.1:30310 --mine
    autorestart=true
    autostart=true
    stdout_logfile=/home/prod/ethereum/poa/log/node1.log
    redirect_stderr=true
    stopsignal=QUIT

    [program:poanode2]
    command= /home/prod/ethereum/geth-alltools/geth --syncmode 'full' --rpc --rpcport "8542" --rpcaddr "0.0.0.0" --datadir /home/prod/ethereum/poa/nodes/node2 --port "30302" --rpccorsdomain "*" --rpcapi "personal,db,eth,net,web3,admin,txpool,miner,clique" --networkid 1034 --unlock 0xe1ace7e1a98ab69abdd45f58aa2ed0899f7fe236 --password /home/prod/ethereum/poa/nodes/node2/password --bootnodes enode://8ea516886bfaa84dae6fec94c6eb4d9babbe4a986a5d21e3e089b1ec6be1ad704281a9c6f51fd1bfaaa611de9228c37f9e945ffe3a49e6ffbc055e55d16373d6@127.0.0.1:30310 --mine
    autorestart=true
    autostart=true
    stdout_logfile=/home/prod/ethereum/poa/log/node2.log
    redirect_stderr=true
    stopsignal=QUIT

    [program:poanode3]
    command= /home/prod/ethereum/geth-alltools/geth --syncmode 'full' --rpc --rpcport "8543" --rpcaddr "0.0.0.0" --datadir /home/prod/ethereum/poa/nodes/node3 --port "30303" --rpccorsdomain "*" --rpcapi "personal,db,eth,net,web3,admin,txpool,miner,clique" --networkid 1034 --unlock 0x2befa8094fd7117c1a518dfd795a4e44accfa245 --password /home/prod/ethereum/poa/nodes/node3/password --bootnodes enode://8ea516886bfaa84dae6fec94c6eb4d9babbe4a986a5d21e3e089b1ec6be1ad704281a9c6f51fd1bfaaa611de9228c37f9e945ffe3a49e6ffbc055e55d16373d6@127.0.0.1:30310 --mine
    autorestart=true
    autostart=true
    stdout_logfile=/home/prod/ethereum/poa/log/node3.log
    redirect_stderr=true
    stopsignal=QUIT
    ```

    ```bash
    # 重新加载, 对已经存在的 reread
    prod@ubuntu:~/ethereum/poa$ sudo supervisorctl reload
    ```