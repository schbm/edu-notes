# OSI

## Aufbau
Layer | Schicht | Desc | Example | Units | Classification
------|---------|------|---------|-------|---------------
Application | Anwendungen | Provides functions for apps | SMTP, FTP, TELNET, SSH, HTTP | Data | EtE
Presentation | Darstellung | Presentation of the information abd adaptation of tasks to the operating system | ASCI, JPEG, GIF, MP4 | Data | EtE
Session | Sitzung | Organization and synchronization of the dialog and handling of the data exchange | smb, iscsi, interprocess communication | Data | EtE
Transport | Transport | Logical end-to-end connections between applications on different computers | TCP (reliable), UDP (unreliable) | segment (TCP PDU), datagram (UDP PDU) | EtE
Network | Vermittlung | Worldwide addressing, path determination in the network: routing Data flow control | IPv4, IPv6, Routingprotocols (RIP, OSPF, BGP etc..), icmp | packet, datagram (IP PDU) | EtE
Data Link | Sicherung | Logical connections with data packets, Elementary error detection mechanisms | Ethernet, WLAN, PPP, LTE | frames | PtP
Physical | Bitübertragung | Transmission of bit information via the physical medium | 1000BASE-T, Token Ring, (Encoding, modulation, cable, connector) | bits, symbols | PtP

## Quellen und Unterlagen
- [ISO Standards for OSI](https://www.iso.org/ics/35.100/x/)
