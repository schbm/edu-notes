| Command                                                               | Description                                                                                                                                                                                        |
|-----------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| enable                                                                | enter privileged mode. allows users to view systemconfiguration, restart the system, and enter configuration mode. allows commands available in user mode. privileged mode can be identified by #. |
| disable                                                               | exit privileged mode                                                                                                                                                                               |
| enable password                                                       | todo                                                                                                                                                                                               |
| enable secret                                                         | todo                                                                                                                                                                                               |
| #configure terminal                                                   | enter configuration mode. allows users to modify the running system configuration                                                                                                                  |
| Router(config)# end                                                   | exit configuration mode                                                                                                                                                                            |
| Router(config)# hostname name                                         | change hostname                                                                                                                                                                                    |
| R1(config)#line console 0 R1(config-line)#logging synchronous         | synchronize unsolicited message                                                                                                                                                                    |
| R1(config)#line console 0 R1(config-line)#exec-timeout 0 0            | set the interval that the EXEC command interpreter waits until user input is detected,                                                                                                             |
| R1#debug ip routing                                                   | shows when routes are added, modified, and deleted from therouting table.                                                                                                                          |
| Router(config)#no ip domain lookup                                    | The problem inherent to name resolution is that when you mistype a command, the Routerinterprets it as a hostnam                                                                                   |
| Router(config-if)#description **Link to XYZ**                         | interface description                                                                                                                                                                              |
| Router(config)# ip route <network-address> <subnet-mask> <ip-address> | add static route                                                                                                                                                                                   |
| R1(config)# interface <interface>                                     | add interface configuration e.g "fastethernet 0/0", "fastethernet 0/1-10", "GI 1/0/5-8". For subinterface e.g:"FastEthernet0/1.10"                                                                 |
| R1(config-if)# no shutdown                                            | enable interface                                                                                                                                                                                   |
| ?                                                                     | show possible commands                                                                                                                                                                             |
| show version                                                          | show current version information                                                                                                                                                                   |
| show running-config                                                   | view current configuration (DRAM)                                                                                                                                                                  |
| show startup-config                                                   | view current configuration (NVRAM)                                                                                                                                                                 |
| show flash                                                            | show IOS file and flash space                                                                                                                                                                      |
| show ip route                                                         | display ip routing table                                                                                                                                                                           |
| show cdp neighbor                                                     | summary of connected cdp devices                                                                                                                                                                   |
|                                                                       |                                                                                                                                                                                                    |